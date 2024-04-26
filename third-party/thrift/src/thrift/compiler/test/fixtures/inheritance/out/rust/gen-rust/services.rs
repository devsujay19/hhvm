// @generated by Thrift for thrift/compiler/test/fixtures/inheritance/src/module.thrift
// This file is probably not the place you want to edit!

//! Thrift service definitions for `module`.


/// Service definitions for `MyRoot`.
pub mod my_root {
    #[derive(Clone, Debug)]
    pub enum DoRootExn {

        ApplicationException(::fbthrift::ApplicationException),
    }

    impl ::std::convert::From<crate::errors::my_root::DoRootError> for DoRootExn {
        fn from(err: crate::errors::my_root::DoRootError) -> Self {
            match err {
                crate::errors::my_root::DoRootError::ApplicationException(aexn) => DoRootExn::ApplicationException(aexn),
                crate::errors::my_root::DoRootError::ThriftError(err) => DoRootExn::ApplicationException(::fbthrift::ApplicationException {
                    message: err.to_string(),
                    type_: ::fbthrift::ApplicationExceptionErrorCode::InternalError,
                }),
            }
        }
    }

    impl ::std::convert::From<::fbthrift::ApplicationException> for DoRootExn {
        fn from(exn: ::fbthrift::ApplicationException) -> Self {
            Self::ApplicationException(exn)
        }
    }

    impl ::fbthrift::ExceptionInfo for DoRootExn {
        fn exn_name(&self) -> &'static str {
            match self {
                Self::ApplicationException(aexn) => aexn.exn_name(),
            }
        }

        fn exn_value(&self) -> String {
            match self {
                Self::ApplicationException(aexn) => aexn.exn_value(),
            }
        }

        fn exn_is_declared(&self) -> bool {
            match self {
                Self::ApplicationException(aexn) => aexn.exn_is_declared(),
            }
        }
    }

    impl ::fbthrift::ResultInfo for DoRootExn {
        fn result_type(&self) -> ::fbthrift::ResultType {
            match self {
                Self::ApplicationException(_aexn) => ::fbthrift::ResultType::Exception,
            }
        }
    }

    impl ::fbthrift::GetTType for DoRootExn {
        const TTYPE: ::fbthrift::TType = ::fbthrift::TType::Struct;
    }

    impl<P> ::fbthrift::Serialize<P> for DoRootExn
    where
        P: ::fbthrift::ProtocolWriter,
    {
        fn write(&self, p: &mut P) {
            ::fbthrift::help::SerializeExn::write_result(Err(self), p);
        }
    }

    impl ::fbthrift::help::SerializeExn for DoRootExn {
        type Success = ();

        fn write_result<P>(res: ::std::result::Result<&Self::Success, &Self>, p: &mut P)
        where
            P: ::fbthrift::ProtocolWriter,
        {
            if let ::std::result::Result::Err(Self::ApplicationException(aexn)) = res {
                ::fbthrift::Serialize::write(aexn, p);
                return;
            }
            p.write_struct_begin("DoRoot");
            match res {
                ::std::result::Result::Ok(success) => {
                    p.write_field_begin(
                        "Success",
                        ::fbthrift::TType::Void,
                        0i16,
                    );
                    ::fbthrift::Serialize::write(success, p);
                    p.write_field_end();
                }

                ::std::result::Result::Err(Self::ApplicationException(_aexn)) => unreachable!(),
            }
            p.write_field_stop();
            p.write_struct_end();
        }
    }

    impl ::fbthrift::help::DeserializeExn for DoRootExn {
        type Success = ();

        fn read_result<P>(p: &mut P) -> ::anyhow::Result<::std::result::Result<Self::Success, Self>>
        where
            P: ::fbthrift::ProtocolReader,
        {
            static RETURNS: &[::fbthrift::Field] = &[
                ::fbthrift::Field::new("Success", ::fbthrift::TType::Void, 0),
            ];
            let _ = p.read_struct_begin(|_| ())?;
            let mut once = false;
            let mut alt = ::std::result::Result::Ok(());
            loop {
                let (_, fty, fid) = p.read_field_begin(|_| (), RETURNS)?;
                match ((fty, fid as ::std::primitive::i32), once) {
                    ((::fbthrift::TType::Stop, _), _) => {
                        p.read_field_end()?;
                        break;
                    }
                    ((::fbthrift::TType::Void, 0i32), false) => {
                        once = true;
                        alt = ::std::result::Result::Ok(::fbthrift::Deserialize::read(p)?);
                    }
                    ((ty, _id), false) => p.skip(ty)?,
                    ((badty, badid), true) => return ::std::result::Result::Err(::std::convert::From::from(
                        ::fbthrift::ApplicationException::new(
                            ::fbthrift::ApplicationExceptionErrorCode::ProtocolError,
                            format!(
                                "unwanted extra union {} field ty {:?} id {}",
                                "DoRootExn",
                                badty,
                                badid,
                            ),
                        )
                    )),
                }
                p.read_field_end()?;
            }
            p.read_struct_end()?;
            ::std::result::Result::Ok(alt)
        }
    }
}

/// Service definitions for `MyNode`.
pub mod my_node {
    #[derive(Clone, Debug)]
    pub enum DoMidExn {

        ApplicationException(::fbthrift::ApplicationException),
    }

    impl ::std::convert::From<crate::errors::my_node::DoMidError> for DoMidExn {
        fn from(err: crate::errors::my_node::DoMidError) -> Self {
            match err {
                crate::errors::my_node::DoMidError::ApplicationException(aexn) => DoMidExn::ApplicationException(aexn),
                crate::errors::my_node::DoMidError::ThriftError(err) => DoMidExn::ApplicationException(::fbthrift::ApplicationException {
                    message: err.to_string(),
                    type_: ::fbthrift::ApplicationExceptionErrorCode::InternalError,
                }),
            }
        }
    }

    impl ::std::convert::From<::fbthrift::ApplicationException> for DoMidExn {
        fn from(exn: ::fbthrift::ApplicationException) -> Self {
            Self::ApplicationException(exn)
        }
    }

    impl ::fbthrift::ExceptionInfo for DoMidExn {
        fn exn_name(&self) -> &'static str {
            match self {
                Self::ApplicationException(aexn) => aexn.exn_name(),
            }
        }

        fn exn_value(&self) -> String {
            match self {
                Self::ApplicationException(aexn) => aexn.exn_value(),
            }
        }

        fn exn_is_declared(&self) -> bool {
            match self {
                Self::ApplicationException(aexn) => aexn.exn_is_declared(),
            }
        }
    }

    impl ::fbthrift::ResultInfo for DoMidExn {
        fn result_type(&self) -> ::fbthrift::ResultType {
            match self {
                Self::ApplicationException(_aexn) => ::fbthrift::ResultType::Exception,
            }
        }
    }

    impl ::fbthrift::GetTType for DoMidExn {
        const TTYPE: ::fbthrift::TType = ::fbthrift::TType::Struct;
    }

    impl<P> ::fbthrift::Serialize<P> for DoMidExn
    where
        P: ::fbthrift::ProtocolWriter,
    {
        fn write(&self, p: &mut P) {
            ::fbthrift::help::SerializeExn::write_result(Err(self), p);
        }
    }

    impl ::fbthrift::help::SerializeExn for DoMidExn {
        type Success = ();

        fn write_result<P>(res: ::std::result::Result<&Self::Success, &Self>, p: &mut P)
        where
            P: ::fbthrift::ProtocolWriter,
        {
            if let ::std::result::Result::Err(Self::ApplicationException(aexn)) = res {
                ::fbthrift::Serialize::write(aexn, p);
                return;
            }
            p.write_struct_begin("DoMid");
            match res {
                ::std::result::Result::Ok(success) => {
                    p.write_field_begin(
                        "Success",
                        ::fbthrift::TType::Void,
                        0i16,
                    );
                    ::fbthrift::Serialize::write(success, p);
                    p.write_field_end();
                }

                ::std::result::Result::Err(Self::ApplicationException(_aexn)) => unreachable!(),
            }
            p.write_field_stop();
            p.write_struct_end();
        }
    }

    impl ::fbthrift::help::DeserializeExn for DoMidExn {
        type Success = ();

        fn read_result<P>(p: &mut P) -> ::anyhow::Result<::std::result::Result<Self::Success, Self>>
        where
            P: ::fbthrift::ProtocolReader,
        {
            static RETURNS: &[::fbthrift::Field] = &[
                ::fbthrift::Field::new("Success", ::fbthrift::TType::Void, 0),
            ];
            let _ = p.read_struct_begin(|_| ())?;
            let mut once = false;
            let mut alt = ::std::result::Result::Ok(());
            loop {
                let (_, fty, fid) = p.read_field_begin(|_| (), RETURNS)?;
                match ((fty, fid as ::std::primitive::i32), once) {
                    ((::fbthrift::TType::Stop, _), _) => {
                        p.read_field_end()?;
                        break;
                    }
                    ((::fbthrift::TType::Void, 0i32), false) => {
                        once = true;
                        alt = ::std::result::Result::Ok(::fbthrift::Deserialize::read(p)?);
                    }
                    ((ty, _id), false) => p.skip(ty)?,
                    ((badty, badid), true) => return ::std::result::Result::Err(::std::convert::From::from(
                        ::fbthrift::ApplicationException::new(
                            ::fbthrift::ApplicationExceptionErrorCode::ProtocolError,
                            format!(
                                "unwanted extra union {} field ty {:?} id {}",
                                "DoMidExn",
                                badty,
                                badid,
                            ),
                        )
                    )),
                }
                p.read_field_end()?;
            }
            p.read_struct_end()?;
            ::std::result::Result::Ok(alt)
        }
    }
}

/// Service definitions for `MyLeaf`.
pub mod my_leaf {
    #[derive(Clone, Debug)]
    pub enum DoLeafExn {

        ApplicationException(::fbthrift::ApplicationException),
    }

    impl ::std::convert::From<crate::errors::my_leaf::DoLeafError> for DoLeafExn {
        fn from(err: crate::errors::my_leaf::DoLeafError) -> Self {
            match err {
                crate::errors::my_leaf::DoLeafError::ApplicationException(aexn) => DoLeafExn::ApplicationException(aexn),
                crate::errors::my_leaf::DoLeafError::ThriftError(err) => DoLeafExn::ApplicationException(::fbthrift::ApplicationException {
                    message: err.to_string(),
                    type_: ::fbthrift::ApplicationExceptionErrorCode::InternalError,
                }),
            }
        }
    }

    impl ::std::convert::From<::fbthrift::ApplicationException> for DoLeafExn {
        fn from(exn: ::fbthrift::ApplicationException) -> Self {
            Self::ApplicationException(exn)
        }
    }

    impl ::fbthrift::ExceptionInfo for DoLeafExn {
        fn exn_name(&self) -> &'static str {
            match self {
                Self::ApplicationException(aexn) => aexn.exn_name(),
            }
        }

        fn exn_value(&self) -> String {
            match self {
                Self::ApplicationException(aexn) => aexn.exn_value(),
            }
        }

        fn exn_is_declared(&self) -> bool {
            match self {
                Self::ApplicationException(aexn) => aexn.exn_is_declared(),
            }
        }
    }

    impl ::fbthrift::ResultInfo for DoLeafExn {
        fn result_type(&self) -> ::fbthrift::ResultType {
            match self {
                Self::ApplicationException(_aexn) => ::fbthrift::ResultType::Exception,
            }
        }
    }

    impl ::fbthrift::GetTType for DoLeafExn {
        const TTYPE: ::fbthrift::TType = ::fbthrift::TType::Struct;
    }

    impl<P> ::fbthrift::Serialize<P> for DoLeafExn
    where
        P: ::fbthrift::ProtocolWriter,
    {
        fn write(&self, p: &mut P) {
            ::fbthrift::help::SerializeExn::write_result(Err(self), p);
        }
    }

    impl ::fbthrift::help::SerializeExn for DoLeafExn {
        type Success = ();

        fn write_result<P>(res: ::std::result::Result<&Self::Success, &Self>, p: &mut P)
        where
            P: ::fbthrift::ProtocolWriter,
        {
            if let ::std::result::Result::Err(Self::ApplicationException(aexn)) = res {
                ::fbthrift::Serialize::write(aexn, p);
                return;
            }
            p.write_struct_begin("DoLeaf");
            match res {
                ::std::result::Result::Ok(success) => {
                    p.write_field_begin(
                        "Success",
                        ::fbthrift::TType::Void,
                        0i16,
                    );
                    ::fbthrift::Serialize::write(success, p);
                    p.write_field_end();
                }

                ::std::result::Result::Err(Self::ApplicationException(_aexn)) => unreachable!(),
            }
            p.write_field_stop();
            p.write_struct_end();
        }
    }

    impl ::fbthrift::help::DeserializeExn for DoLeafExn {
        type Success = ();

        fn read_result<P>(p: &mut P) -> ::anyhow::Result<::std::result::Result<Self::Success, Self>>
        where
            P: ::fbthrift::ProtocolReader,
        {
            static RETURNS: &[::fbthrift::Field] = &[
                ::fbthrift::Field::new("Success", ::fbthrift::TType::Void, 0),
            ];
            let _ = p.read_struct_begin(|_| ())?;
            let mut once = false;
            let mut alt = ::std::result::Result::Ok(());
            loop {
                let (_, fty, fid) = p.read_field_begin(|_| (), RETURNS)?;
                match ((fty, fid as ::std::primitive::i32), once) {
                    ((::fbthrift::TType::Stop, _), _) => {
                        p.read_field_end()?;
                        break;
                    }
                    ((::fbthrift::TType::Void, 0i32), false) => {
                        once = true;
                        alt = ::std::result::Result::Ok(::fbthrift::Deserialize::read(p)?);
                    }
                    ((ty, _id), false) => p.skip(ty)?,
                    ((badty, badid), true) => return ::std::result::Result::Err(::std::convert::From::from(
                        ::fbthrift::ApplicationException::new(
                            ::fbthrift::ApplicationExceptionErrorCode::ProtocolError,
                            format!(
                                "unwanted extra union {} field ty {:?} id {}",
                                "DoLeafExn",
                                badty,
                                badid,
                            ),
                        )
                    )),
                }
                p.read_field_end()?;
            }
            p.read_struct_end()?;
            ::std::result::Result::Ok(alt)
        }
    }
}
