// Copyright (c) Meta Platforms, Inc. and affiliates.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the "hack" directory of this source tree.

mod display;

use derive_more::Display;
use hash::IndexMap;

pub struct File {
    pub defs: IndexMap<TypeName, Def>,
}

pub enum Def {
    Alias {
        doc: Vec<String>,
        tparams: Vec<String>,
        ty: Type,
    },
    Record {
        doc: Vec<String>,
        tparams: Vec<String>,
        fields: Vec<Field>,
    },
    Variant {
        doc: Vec<String>,
        tparams: Vec<String>,
        variants: Vec<Variant>,
    },
}

pub struct Variant {
    pub name: VariantName,
    pub fields: Option<VariantFields>,
    pub doc: Vec<String>,
}

pub enum VariantFields {
    Unnamed(Vec<Type>),
    Named(Vec<Field>),
}

pub struct Field {
    pub name: FieldName,
    pub ty: Type,
    pub doc: Vec<String>,
}

pub enum Type {
    Path(TypePath),
    Tuple(TypeTuple),
}

pub struct TypePath {
    pub targs: Vec<Type>,
    pub modules: Vec<ModuleName>,
    pub ty: TypeName,
}

impl TypePath {
    pub fn simple(id: impl Into<String>) -> Self {
        Self {
            modules: vec![],
            ty: TypeName(id.into()),
            targs: vec![],
        }
    }
}

pub struct TypeTuple {
    pub elems: Vec<Type>,
}

#[derive(Clone, Hash, PartialEq, Eq)]
pub struct ModuleName(pub String);

impl ModuleName {
    pub fn as_str(&self) -> &str {
        self.0.as_str()
    }
}

#[derive(Clone, Hash, PartialEq, Eq)]
pub struct TypeName(pub String);

impl TypeName {
    pub fn as_str(&self) -> &str {
        self.0.as_str()
    }
}

#[derive(Clone, Hash, PartialEq, Eq, Display)]
pub struct FieldName(pub String);

#[derive(Clone, Hash, PartialEq, Eq, Display)]
pub struct VariantName(pub String);
