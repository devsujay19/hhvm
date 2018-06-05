(**
 * Copyright (c) 2015, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the "hack" directory of this source tree.
 *
 *)

let get_level_list check =
  let type_acc : (Pos.t, Typing_reason.t * Typing_defs.locl Typing_defs.ty_)
    Hashtbl.t = Hashtbl.create 0 in
  let fn = Typing.with_expr_hook
    (fun (p, _) ty -> Hashtbl.replace type_acc p ty) check in
  let level_of_type x = snd (Coverage_level.level_of_type_mapper fn x) in
  let result = Hashtbl.fold (fun p ty xs ->
    (Pos.to_absolute p, level_of_type (p, ty)) :: xs) type_acc [] in
  result

let go env f_in =
  get_level_list begin fun () ->
    fst (ServerIdeUtils.check_file_input
      env.ServerEnv.tcopt env.ServerEnv.files_info f_in)
  end
