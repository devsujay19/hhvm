// @generated by Thrift for thrift/compiler/test/fixtures/basic/src/module.thrift
// This file is probably not the place you want to edit!

#![recursion_limit = "100000000"]
#![allow(non_camel_case_types, non_snake_case, non_upper_case_globals, unused_crate_dependencies, clippy::all)]

#[allow(unused_imports)]
#[doc(no_inline)]
pub use self::errors::*;
#[doc(inline)]
pub use self::types::*;

#[doc(inline)]
pub use :: as types;

#[doc(hidden)]
pub mod dependencies;
#[doc(inline)]
pub use ::::services as services;
#[doc(inline)]
pub use ::::{self as client, mock};
#[doc(inline)]
pub use :: as server;
#[doc(inline)]
pub use ::::errors as errors;
