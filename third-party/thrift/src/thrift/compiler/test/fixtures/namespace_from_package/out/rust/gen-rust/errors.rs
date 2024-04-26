// @generated by Thrift for thrift/compiler/test/fixtures/namespace_from_package/src/module.thrift
// This file is probably not the place you want to edit!

//! Thrift error definitions for `module`.

/// Error definitions for `TestService`.
pub mod test_service {

    pub type InitError = ::fbthrift::NonthrowingFunctionError;

    impl ::fbthrift::help::StreamExn for crate::services::test_service::InitExn {
        type Success = ::std::primitive::i64;
        type Return = ::std::primitive::i64;
        type Error = InitError;

        fn map_stream(res: ::std::result::Result<Self::Success, Self>) -> ::std::result::Result<Self::Return, Self::Error> {
            match res {
                ::std::result::Result::Ok(success) => ::std::result::Result::Ok(success),
                ::std::result::Result::Err(exn) => ::std::result::Result::Err(::std::convert::From::from(exn)),
            }
        }
    }

    impl ::std::convert::From<crate::services::test_service::InitExn> for InitError {
        fn from(e: crate::services::test_service::InitExn) -> Self {
            match e {
                crate::services::test_service::InitExn::ApplicationException(aexn) =>
                    InitError::ApplicationException(aexn),
            }
        }
    }

}

