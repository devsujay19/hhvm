// @generated by Thrift for thrift/compiler/test/fixtures/types/src/module.thrift
// This file is probably not the place you want to edit!

//! Thrift error definitions for `module`.

/// Error definitions for `SomeService`.
pub mod some_service {

    pub type BounceMapError = ::fbthrift::NonthrowingFunctionError;

    impl ::fbthrift::help::StreamExn for crate::services::some_service::BounceMapExn {
        type Success = included__types::SomeMap;
        type Return = included__types::SomeMap;
        type Error = BounceMapError;

        fn map_stream(res: ::std::result::Result<Self::Success, Self>) -> ::std::result::Result<Self::Return, Self::Error> {
            match res {
                ::std::result::Result::Ok(success) => ::std::result::Result::Ok(success),
                ::std::result::Result::Err(exn) => ::std::result::Result::Err(::std::convert::From::from(exn)),
            }
        }
    }

    impl ::std::convert::From<crate::services::some_service::BounceMapExn> for BounceMapError {
        fn from(e: crate::services::some_service::BounceMapExn) -> Self {
            match e {
                crate::services::some_service::BounceMapExn::ApplicationException(aexn) =>
                    BounceMapError::ApplicationException(aexn),
            }
        }
    }

    pub type BinaryKeyedMapError = ::fbthrift::NonthrowingFunctionError;

    impl ::fbthrift::help::StreamExn for crate::services::some_service::BinaryKeyedMapExn {
        type Success = ::std::collections::BTreeMap<crate::types::TBinary, ::std::primitive::i64>;
        type Return = ::std::collections::BTreeMap<crate::types::TBinary, ::std::primitive::i64>;
        type Error = BinaryKeyedMapError;

        fn map_stream(res: ::std::result::Result<Self::Success, Self>) -> ::std::result::Result<Self::Return, Self::Error> {
            match res {
                ::std::result::Result::Ok(success) => ::std::result::Result::Ok(success),
                ::std::result::Result::Err(exn) => ::std::result::Result::Err(::std::convert::From::from(exn)),
            }
        }
    }

    impl ::std::convert::From<crate::services::some_service::BinaryKeyedMapExn> for BinaryKeyedMapError {
        fn from(e: crate::services::some_service::BinaryKeyedMapExn) -> Self {
            match e {
                crate::services::some_service::BinaryKeyedMapExn::ApplicationException(aexn) =>
                    BinaryKeyedMapError::ApplicationException(aexn),
            }
        }
    }

}

