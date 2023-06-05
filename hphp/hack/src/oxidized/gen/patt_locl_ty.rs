// Copyright (c) Facebook, Inc. and its affiliates.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the "hack" directory of this source tree.
//
// @generated SignedSource<<38d77fbb2e6a2bed126ffe357c51974b>>
//
// To regenerate this file, run:
//   hphp/hack/src/oxidized_regen.sh

pub use core::*;

use arena_trait::TrivialDrop;
use eq_modulo_pos::EqModuloPos;
use no_pos_hash::NoPosHash;
use ocamlrep::FromOcamlRep;
use ocamlrep::FromOcamlRepIn;
use ocamlrep::ToOcamlRep;
use serde::Deserialize;
use serde::Serialize;

pub use crate::validated as v;
#[allow(unused_imports)]
use crate::*;

#[derive(
    Clone,
    Debug,
    Deserialize,
    Eq,
    EqModuloPos,
    FromOcamlRep,
    Hash,
    NoPosHash,
    Ord,
    PartialEq,
    PartialOrd,
    Serialize,
    ToOcamlRep
)]
#[repr(C, u8)]
pub enum PattLoclTy {
    #[rust_to_ocaml(prefix = "patt_")]
    Apply {
        name: patt_name::PattName,
        params: Params,
    },
    Prim(Prim),
    Shape(ShapeFields),
    Option(Box<PattLoclTy>),
    Tuple(Vec<PattLoclTy>),
    Dynamic,
    Nonnull,
    Any,
    #[rust_to_ocaml(prefix = "patt_")]
    Or {
        fst: Box<PattLoclTy>,
        snd: Box<PattLoclTy>,
    },
    As {
        lbl: patt_var::PattVar,
        patt: Box<PattLoclTy>,
    },
    Invalid(Vec<validation_err::ValidationErr>, Box<PattLoclTy>),
}

#[derive(
    Clone,
    Debug,
    Deserialize,
    Eq,
    EqModuloPos,
    FromOcamlRep,
    Hash,
    NoPosHash,
    Ord,
    PartialEq,
    PartialOrd,
    Serialize,
    ToOcamlRep
)]
#[rust_to_ocaml(and)]
#[repr(C, u8)]
pub enum Params {
    Nil,
    Wildcard,
    #[rust_to_ocaml(prefix = "patt_")]
    Cons {
        hd: Box<Params>,
        tl: Box<Params>,
    },
    Exists(Box<Params>),
}

#[derive(
    Clone,
    Copy,
    Debug,
    Deserialize,
    Eq,
    EqModuloPos,
    FromOcamlRep,
    FromOcamlRepIn,
    Hash,
    NoPosHash,
    Ord,
    PartialEq,
    PartialOrd,
    Serialize,
    ToOcamlRep
)]
#[rust_to_ocaml(and)]
#[repr(u8)]
pub enum Prim {
    Null,
    Void,
    Int,
    Bool,
    Float,
    String,
    Resource,
    Num,
    Arraykey,
    Noreturn,
}
impl TrivialDrop for Prim {}
arena_deserializer::impl_deserialize_in_arena!(Prim);

#[derive(
    Clone,
    Debug,
    Deserialize,
    Eq,
    EqModuloPos,
    FromOcamlRep,
    Hash,
    NoPosHash,
    Ord,
    PartialEq,
    PartialOrd,
    Serialize,
    ToOcamlRep
)]
#[rust_to_ocaml(and)]
#[repr(C, u8)]
pub enum ShapeFields {
    #[rust_to_ocaml(prefix = "patt_")]
    Fld {
        fld: ShapeField,
        rest: Box<ShapeFields>,
    },
    Open,
    Closed,
}

#[derive(
    Clone,
    Debug,
    Deserialize,
    Eq,
    EqModuloPos,
    FromOcamlRep,
    Hash,
    NoPosHash,
    Ord,
    PartialEq,
    PartialOrd,
    Serialize,
    ToOcamlRep
)]
#[rust_to_ocaml(and)]
#[repr(C)]
pub struct ShapeField {
    pub lbl: ShapeLabel,
    pub optional: bool,
    pub patt: Box<ShapeField>,
}

#[derive(
    Clone,
    Debug,
    Deserialize,
    Eq,
    EqModuloPos,
    FromOcamlRep,
    Hash,
    NoPosHash,
    Ord,
    PartialEq,
    PartialOrd,
    Serialize,
    ToOcamlRep
)]
#[rust_to_ocaml(and)]
#[rust_to_ocaml(attr = "deriving (compare, eq, sexp, show, yojson)")]
#[repr(C, u8)]
pub enum ShapeLabel {
    StrLbl(String),
    IntLbl(String),
    CConstLbl { cls_nm: String, cnst_nm: String },
}
