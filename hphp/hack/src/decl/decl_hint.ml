(*
 * Copyright (c) 2015, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the "hack" directory of this source tree.
 *
 *)

(*****************************************************************************)
(* Converts a type hint into a type  *)
(*****************************************************************************)
open Hh_prelude
open Aast
open Typing_defs

(* Unpacking a hint for typing *)
let rec hint env (p, h) =
  let h = hint_ p env h in
  mk (Typing_reason.Rhint p, h)

and shape_field_info_to_shape_field_type env { sfi_optional; sfi_hint; _ } =
  { sft_optional = sfi_optional; sft_ty = hint env sfi_hint }

and aast_user_attribute_to_decl_user_attribute { ua_name; ua_params } =
  {
    Typing_defs.ua_name;
    ua_classname_params =
      List.filter_map ua_params ~f:(function
          | (_, Class_const ((_, CI (_, cls)), (_, name)))
            when String.equal name SN.Members.mClass ->
            Some cls
          | _ -> None);
  }

and aast_contexts_to_decl_capability env ctxs default_pos =
  match ctxs with
  | Some (pos, hl) ->
    let hl = List.map ~f:(hint env) hl in
    CapTy (Typing_make_type.intersection (Reason.Rhint pos) hl)
  | None -> CapDefaults default_pos

and aast_tparam_to_decl_tparam env t =
  {
    tp_variance = t.Aast.tp_variance;
    tp_name = t.Aast.tp_name;
    tp_tparams =
      List.map ~f:(aast_tparam_to_decl_tparam env) t.Aast.tp_parameters;
    tp_constraints =
      List.map ~f:(Tuple.T2.map_snd ~f:(hint env)) t.Aast.tp_constraints;
    tp_reified = t.Aast.tp_reified;
    tp_user_attributes =
      List.map
        ~f:aast_user_attribute_to_decl_user_attribute
        t.Aast.tp_user_attributes;
  }

and hint_ p env = function
  | Hany -> Typing_defs.make_tany ()
  | Herr -> Terr
  | Hmixed -> Tmixed
  | Hnonnull -> Tnonnull
  | Hthis -> Tthis
  | Hdynamic -> Tdynamic
  | Hnothing -> Tunion []
  | Hdarray (h1, h2) -> Tdarray (hint env h1, hint env h2)
  | Hvarray h -> Tvarray (hint env h)
  | Hvarray_or_darray (h1, h2) ->
    let t1 =
      match h1 with
      | Some h -> hint env h
      | None -> mk (Typing_reason.Rvarray_or_darray_key p, Tprim Aast.Tarraykey)
    in
    Tvarray_or_darray (t1, hint env h2)
  | Hvec_or_dict (h1, h2) ->
    let t1 =
      match h1 with
      | Some h -> hint env h
      | None -> mk (Typing_reason.Rvec_or_dict_key p, Tprim Aast.Tarraykey)
    in
    Tvec_or_dict (t1, hint env h2)
  | Hprim p -> Tprim p
  | Habstr (x, argl) ->
    let argl = List.map argl (hint env) in
    Tgeneric (x, argl)
  | Hoption h ->
    let h = hint env h in
    Toption h
  | Hlike h -> Tlike (hint env h)
  | Hfun
      {
        hf_reactive_kind = reactivity;
        hf_param_tys = hl;
        hf_param_kinds = kl;
        hf_param_mutability = muts;
        hf_variadic_ty = vh;
        hf_ctxs = ctxs;
        hf_return_ty = h;
        hf_is_mutable_return = mut_ret;
      } ->
    let make_param ((p, _) as x) k mut =
      let mutability =
        match mut with
        | Some PMutable -> Some Param_borrowed_mutable
        | Some POwnedMutable -> Some Param_owned_mutable
        | Some PMaybeMutable -> Some Param_maybe_mutable
        | _ -> None
      in
      {
        fp_pos = p;
        fp_name = None;
        fp_type = possibly_enforced_hint env x;
        fp_flags =
          make_fp_flags
            ~mode:(get_param_mode k)
            ~accept_disposable:false
            ~mutability
            ~has_default:
              false
              (* Currently do not support external and cancall on parameters of function parameters *)
            ~ifc_external:false
            ~ifc_can_call:false
            ~is_atom:
              false
              (* TODO: support readonly on parameters of function parameters *)
            ~readonly:false;
        fp_rx_annotation = None;
      }
    in
    let paraml = List.map3_exn hl kl muts ~f:make_param in
    let implicit_params =
      let capability = aast_contexts_to_decl_capability env ctxs p in
      { capability }
    in
    let ret = possibly_enforced_hint env h in
    let arity =
      match vh with
      | Some t -> Fvariadic (make_param t None None)
      | None -> Fstandard
    in
    let reactivity =
      match reactivity with
      | FPure -> Pure None
      | FNonreactive -> Nonreactive
    in
    Tfun
      {
        ft_arity = arity;
        ft_tparams = [];
        ft_where_constraints = [];
        ft_params = paraml;
        ft_implicit_params = implicit_params;
        ft_ret = ret;
        ft_flags =
          make_ft_flags
            Ast_defs.FSync
            None
            ~return_disposable:false
            ~returns_void_to_rx:false
            ~returns_mutable:mut_ret
            ~returns_readonly:
              false (* TODO: support readonly on function type hints *)
            ~readonly_this:false;
        ft_reactive = reactivity;
        (* TODO: handle function parameters with <<CanCall>> *)
        ft_ifc_decl = default_ifc_fun_decl;
      }
  | Happly (id, argl) ->
    let argl = List.map argl (hint env) in
    Tapply (id, argl)
  | Haccess ((_, Hvar n), [(_, id)]) -> Tgeneric ("T" ^ n ^ "@" ^ id, [])
  | Haccess (root_ty, ids) ->
    let root_ty = hint_ p env (snd root_ty) in
    let rec translate res ids =
      match ids with
      | [] -> res
      | id :: ids ->
        translate (Taccess (mk (Typing_reason.Rhint p, res), id)) ids
    in
    translate root_ty ids
  | Htuple hl ->
    let tyl = List.map hl (hint env) in
    Ttuple tyl
  | Hunion hl ->
    let tyl = List.map hl (hint env) in
    Tunion tyl
  | Hintersection hl ->
    let tyl = List.map hl (hint env) in
    Tintersection tyl
  | Hshape { nsi_allows_unknown_fields; nsi_field_map } ->
    let shape_kind =
      if nsi_allows_unknown_fields then
        Open_shape
      else
        Closed_shape
    in
    let fdm =
      List.fold_left
        ~f:(fun acc i ->
          TShapeMap.add
            (TShapeField.of_ast (fun p -> p) i.sfi_name)
            (shape_field_info_to_shape_field_type env i)
            acc)
        ~init:TShapeMap.empty
        nsi_field_map
    in
    Tshape (shape_kind, fdm)
  | Hsoft (p, h_) -> hint_ p env h_
  | Hfun_context n -> Tgeneric ("Tctx" ^ n, [])
  | Hvar n -> Tgeneric ("T" ^ n, [])

and possibly_enforced_hint env h =
  (* Initially we assume that a type is not enforced at runtime.
   * We refine this during localization
   *)
  { et_enforced = false; et_type = hint env h }
