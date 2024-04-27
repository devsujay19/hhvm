// @generated by Thrift for thrift/compiler/test/fixtures/exceptions/src/module.thrift
// This file is probably not the place you want to edit!

//! Thrift error definitions for `module`.

/// Error definitions for `Raiser`.
pub mod raiser {

    pub trait AsBanal {
        fn as_banal(&self) -> Option<&crate::types::Banal>;
    }

    impl AsBanal for ::anyhow::Error {
        fn as_banal(&self) -> Option<&crate::types::Banal> {
            for cause in self.chain() {
                if let Some(DoRaiseError::b(e)) = cause.downcast_ref::<DoRaiseError>() {
                    return Some(e);
                }
                if let Some(Get500Error::b(e)) = cause.downcast_ref::<Get500Error>() {
                    return Some(e);
                }
            }
            None
        }
    }

    pub trait AsFiery {
        fn as_fiery(&self) -> Option<&crate::types::Fiery>;
    }

    impl AsFiery for ::anyhow::Error {
        fn as_fiery(&self) -> Option<&crate::types::Fiery> {
            for cause in self.chain() {
                if let Some(DoRaiseError::f(e)) = cause.downcast_ref::<DoRaiseError>() {
                    return Some(e);
                }
                if let Some(Get500Error::f(e)) = cause.downcast_ref::<Get500Error>() {
                    return Some(e);
                }
            }
            None
        }
    }

    pub trait AsSerious {
        fn as_serious(&self) -> Option<&crate::types::Serious>;
    }

    impl AsSerious for ::anyhow::Error {
        fn as_serious(&self) -> Option<&crate::types::Serious> {
            for cause in self.chain() {
                if let Some(DoRaiseError::s(e)) = cause.downcast_ref::<DoRaiseError>() {
                    return Some(e);
                }
                if let Some(Get500Error::s(e)) = cause.downcast_ref::<Get500Error>() {
                    return Some(e);
                }
            }
            None
        }
    }

    pub type DoBlandError = ::fbthrift::NonthrowingFunctionError;


    pub(crate) enum DoBlandReader {}

    impl ::fbthrift::help::DeserializeExn for DoBlandReader {
        type Success = ();
        type Error = DoBlandError;

        fn read_result<P>(p: &mut P) -> ::anyhow::Result<::std::result::Result<Self::Success, Self::Error>>
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
                                "DoBlandError",
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

    /// Errors for doRaise (client side).
    #[derive(Debug)]
    pub enum DoRaiseError {
        b(crate::types::Banal),
        f(crate::types::Fiery),
        s(crate::types::Serious),
        ApplicationException(::fbthrift::ApplicationException),
        ThriftError(::anyhow::Error),
    }

    /// Human-readable string representation of the Thrift client error.
    ///
    /// By default, this will not print the full cause chain. If you would like to print the underlying error
    /// cause, either use `format!("{:?}", anyhow::Error::from(client_err))` or print this using the
    /// alternate formatter `{:#}` instead of just `{}`.
    impl ::std::fmt::Display for DoRaiseError {
        fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::result::Result<(), ::std::fmt::Error> {
            match self {
                Self::b(inner) => {
                    if f.alternate() {
                        write!(f, "Raiser::doRaise failed with variant `b`: {:#}", inner)?;
                    } else {
                        write!(f, "Raiser::doRaise failed with b(Banal)")?;
                    }
                }
                Self::f(inner) => {
                    if f.alternate() {
                        write!(f, "Raiser::doRaise failed with variant `f`: {:#}", inner)?;
                    } else {
                        write!(f, "Raiser::doRaise failed with f(Fiery)")?;
                    }
                }
                Self::s(inner) => {
                    if f.alternate() {
                        write!(f, "Raiser::doRaise failed with variant `s`: {:#}", inner)?;
                    } else {
                        write!(f, "Raiser::doRaise failed with s(Serious)")?;
                    }
                }
                Self::ApplicationException(inner) => {
                    write!(f, "Raiser::doRaise failed with ApplicationException")?;

                    if f.alternate() {
                      write!(f, ": {:#}", inner)?;
                    }
                }
                Self::ThriftError(inner) => {
                    write!(f, "Raiser::doRaise failed with ThriftError")?;

                    if f.alternate() {
                      write!(f, ": {:#}", inner)?;
                    }
                }
            }

            Ok(())
        }
    }

    impl ::std::error::Error for DoRaiseError {
        fn source(&self) -> ::std::option::Option<&(dyn ::std::error::Error + 'static)> {
            match self {
                Self::b(ref inner) => {
                    Some(inner)
                }
                Self::f(ref inner) => {
                    Some(inner)
                }
                Self::s(ref inner) => {
                    Some(inner)
                }
                Self::ApplicationException(ref inner) => {
                    Some(inner)
                }
                Self::ThriftError(ref inner) => {
                    Some(inner.as_ref())
                }
            }
        }
    }

    impl ::std::convert::From<crate::types::Banal> for DoRaiseError {
        fn from(e: crate::types::Banal) -> Self {
            Self::b(e)
        }
    }

    impl AsBanal for DoRaiseError {
        fn as_banal(&self) -> Option<&crate::types::Banal> {
            match self {
                Self::b(inner) => Some(inner),
                _ => None,
            }
        }
    }

    impl ::std::convert::From<crate::types::Fiery> for DoRaiseError {
        fn from(e: crate::types::Fiery) -> Self {
            Self::f(e)
        }
    }

    impl AsFiery for DoRaiseError {
        fn as_fiery(&self) -> Option<&crate::types::Fiery> {
            match self {
                Self::f(inner) => Some(inner),
                _ => None,
            }
        }
    }

    impl ::std::convert::From<crate::types::Serious> for DoRaiseError {
        fn from(e: crate::types::Serious) -> Self {
            Self::s(e)
        }
    }

    impl AsSerious for DoRaiseError {
        fn as_serious(&self) -> Option<&crate::types::Serious> {
            match self {
                Self::s(inner) => Some(inner),
                _ => None,
            }
        }
    }

    impl ::std::convert::From<::anyhow::Error> for DoRaiseError {
        fn from(err: ::anyhow::Error) -> Self {
            Self::ThriftError(err)
        }
    }

    impl ::std::convert::From<::fbthrift::ApplicationException> for DoRaiseError {
        fn from(ae: ::fbthrift::ApplicationException) -> Self {
            Self::ApplicationException(ae)
        }
    }

    impl ::std::convert::From<crate::services::raiser::DoRaiseExn> for DoRaiseError {
        fn from(e: crate::services::raiser::DoRaiseExn) -> Self {
            match e {
                crate::services::raiser::DoRaiseExn::ApplicationException(aexn) =>
                    DoRaiseError::ApplicationException(aexn),
                crate::services::raiser::DoRaiseExn::b(exn) =>
                    DoRaiseError::b(exn),
                crate::services::raiser::DoRaiseExn::f(exn) =>
                    DoRaiseError::f(exn),
                crate::services::raiser::DoRaiseExn::s(exn) =>
                    DoRaiseError::s(exn),
            }
        }
    }

    impl ::std::convert::From<DoRaiseError> for crate::services::raiser::DoRaiseExn {
        fn from(err: DoRaiseError) -> Self {
            match err {
                DoRaiseError::b(err) => crate::services::raiser::DoRaiseExn::b(err),
                DoRaiseError::f(err) => crate::services::raiser::DoRaiseExn::f(err),
                DoRaiseError::s(err) => crate::services::raiser::DoRaiseExn::s(err),
                DoRaiseError::ApplicationException(aexn) => crate::services::raiser::DoRaiseExn::ApplicationException(aexn),
                DoRaiseError::ThriftError(err) => crate::services::raiser::DoRaiseExn::ApplicationException(::fbthrift::ApplicationException {
                    message: err.to_string(),
                    type_: ::fbthrift::ApplicationExceptionErrorCode::InternalError,
                }),
            }
        }
    }

    pub(crate) enum DoRaiseReader {}

    impl ::fbthrift::help::DeserializeExn for DoRaiseReader {
        type Success = ();
        type Error = DoRaiseError;

        fn read_result<P>(p: &mut P) -> ::anyhow::Result<::std::result::Result<Self::Success, Self::Error>>
        where
            P: ::fbthrift::ProtocolReader,
        {
            static RETURNS: &[::fbthrift::Field] = &[
                ::fbthrift::Field::new("Success", ::fbthrift::TType::Void, 0),
                ::fbthrift::Field::new("b", ::fbthrift::TType::Struct, 1),
                ::fbthrift::Field::new("f", ::fbthrift::TType::Struct, 2),
                ::fbthrift::Field::new("s", ::fbthrift::TType::Struct, 3),
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
                    ((::fbthrift::TType::Struct, 1), false) => {
                        once = true;
                        alt = ::std::result::Result::Err(Self::Error::b(::fbthrift::Deserialize::read(p)?));
                    }
                    ((::fbthrift::TType::Struct, 2), false) => {
                        once = true;
                        alt = ::std::result::Result::Err(Self::Error::f(::fbthrift::Deserialize::read(p)?));
                    }
                    ((::fbthrift::TType::Struct, 3), false) => {
                        once = true;
                        alt = ::std::result::Result::Err(Self::Error::s(::fbthrift::Deserialize::read(p)?));
                    }
                    ((ty, _id), false) => p.skip(ty)?,
                    ((badty, badid), true) => return ::std::result::Result::Err(::std::convert::From::from(
                        ::fbthrift::ApplicationException::new(
                            ::fbthrift::ApplicationExceptionErrorCode::ProtocolError,
                            format!(
                                "unwanted extra union {} field ty {:?} id {}",
                                "DoRaiseError",
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

    pub type Get200Error = ::fbthrift::NonthrowingFunctionError;


    pub(crate) enum Get200Reader {}

    impl ::fbthrift::help::DeserializeExn for Get200Reader {
        type Success = ::std::string::String;
        type Error = Get200Error;

        fn read_result<P>(p: &mut P) -> ::anyhow::Result<::std::result::Result<Self::Success, Self::Error>>
        where
            P: ::fbthrift::ProtocolReader,
        {
            static RETURNS: &[::fbthrift::Field] = &[
                ::fbthrift::Field::new("Success", ::fbthrift::TType::String, 0),
            ];
            let _ = p.read_struct_begin(|_| ())?;
            let mut once = false;
            let mut alt = ::std::option::Option::None;
            loop {
                let (_, fty, fid) = p.read_field_begin(|_| (), RETURNS)?;
                match ((fty, fid as ::std::primitive::i32), once) {
                    ((::fbthrift::TType::Stop, _), _) => {
                        p.read_field_end()?;
                        break;
                    }
                    ((::fbthrift::TType::String, 0i32), false) => {
                        once = true;
                        alt = ::std::option::Option::Some(::std::result::Result::Ok(::fbthrift::Deserialize::read(p)?));
                    }
                    ((ty, _id), false) => p.skip(ty)?,
                    ((badty, badid), true) => return ::std::result::Result::Err(::std::convert::From::from(
                        ::fbthrift::ApplicationException::new(
                            ::fbthrift::ApplicationExceptionErrorCode::ProtocolError,
                            format!(
                                "unwanted extra union {} field ty {:?} id {}",
                                "Get200Error",
                                badty,
                                badid,
                            ),
                        )
                    )),
                }
                p.read_field_end()?;
            }
            p.read_struct_end()?;
            alt.ok_or_else(||
                ::fbthrift::ApplicationException::new(
                    ::fbthrift::ApplicationExceptionErrorCode::MissingResult,
                    format!("Empty union {}", "Get200Error"),
                )
                .into(),
            )
        }
    }

    /// Errors for get500 (client side).
    #[derive(Debug)]
    pub enum Get500Error {
        f(crate::types::Fiery),
        b(crate::types::Banal),
        s(crate::types::Serious),
        ApplicationException(::fbthrift::ApplicationException),
        ThriftError(::anyhow::Error),
    }

    /// Human-readable string representation of the Thrift client error.
    ///
    /// By default, this will not print the full cause chain. If you would like to print the underlying error
    /// cause, either use `format!("{:?}", anyhow::Error::from(client_err))` or print this using the
    /// alternate formatter `{:#}` instead of just `{}`.
    impl ::std::fmt::Display for Get500Error {
        fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::result::Result<(), ::std::fmt::Error> {
            match self {
                Self::f(inner) => {
                    if f.alternate() {
                        write!(f, "Raiser::get500 failed with variant `f`: {:#}", inner)?;
                    } else {
                        write!(f, "Raiser::get500 failed with f(Fiery)")?;
                    }
                }
                Self::b(inner) => {
                    if f.alternate() {
                        write!(f, "Raiser::get500 failed with variant `b`: {:#}", inner)?;
                    } else {
                        write!(f, "Raiser::get500 failed with b(Banal)")?;
                    }
                }
                Self::s(inner) => {
                    if f.alternate() {
                        write!(f, "Raiser::get500 failed with variant `s`: {:#}", inner)?;
                    } else {
                        write!(f, "Raiser::get500 failed with s(Serious)")?;
                    }
                }
                Self::ApplicationException(inner) => {
                    write!(f, "Raiser::get500 failed with ApplicationException")?;

                    if f.alternate() {
                      write!(f, ": {:#}", inner)?;
                    }
                }
                Self::ThriftError(inner) => {
                    write!(f, "Raiser::get500 failed with ThriftError")?;

                    if f.alternate() {
                      write!(f, ": {:#}", inner)?;
                    }
                }
            }

            Ok(())
        }
    }

    impl ::std::error::Error for Get500Error {
        fn source(&self) -> ::std::option::Option<&(dyn ::std::error::Error + 'static)> {
            match self {
                Self::f(ref inner) => {
                    Some(inner)
                }
                Self::b(ref inner) => {
                    Some(inner)
                }
                Self::s(ref inner) => {
                    Some(inner)
                }
                Self::ApplicationException(ref inner) => {
                    Some(inner)
                }
                Self::ThriftError(ref inner) => {
                    Some(inner.as_ref())
                }
            }
        }
    }

    impl ::std::convert::From<crate::types::Fiery> for Get500Error {
        fn from(e: crate::types::Fiery) -> Self {
            Self::f(e)
        }
    }

    impl AsFiery for Get500Error {
        fn as_fiery(&self) -> Option<&crate::types::Fiery> {
            match self {
                Self::f(inner) => Some(inner),
                _ => None,
            }
        }
    }

    impl ::std::convert::From<crate::types::Banal> for Get500Error {
        fn from(e: crate::types::Banal) -> Self {
            Self::b(e)
        }
    }

    impl AsBanal for Get500Error {
        fn as_banal(&self) -> Option<&crate::types::Banal> {
            match self {
                Self::b(inner) => Some(inner),
                _ => None,
            }
        }
    }

    impl ::std::convert::From<crate::types::Serious> for Get500Error {
        fn from(e: crate::types::Serious) -> Self {
            Self::s(e)
        }
    }

    impl AsSerious for Get500Error {
        fn as_serious(&self) -> Option<&crate::types::Serious> {
            match self {
                Self::s(inner) => Some(inner),
                _ => None,
            }
        }
    }

    impl ::std::convert::From<::anyhow::Error> for Get500Error {
        fn from(err: ::anyhow::Error) -> Self {
            Self::ThriftError(err)
        }
    }

    impl ::std::convert::From<::fbthrift::ApplicationException> for Get500Error {
        fn from(ae: ::fbthrift::ApplicationException) -> Self {
            Self::ApplicationException(ae)
        }
    }

    impl ::std::convert::From<crate::services::raiser::Get500Exn> for Get500Error {
        fn from(e: crate::services::raiser::Get500Exn) -> Self {
            match e {
                crate::services::raiser::Get500Exn::ApplicationException(aexn) =>
                    Get500Error::ApplicationException(aexn),
                crate::services::raiser::Get500Exn::f(exn) =>
                    Get500Error::f(exn),
                crate::services::raiser::Get500Exn::b(exn) =>
                    Get500Error::b(exn),
                crate::services::raiser::Get500Exn::s(exn) =>
                    Get500Error::s(exn),
            }
        }
    }

    impl ::std::convert::From<Get500Error> for crate::services::raiser::Get500Exn {
        fn from(err: Get500Error) -> Self {
            match err {
                Get500Error::f(err) => crate::services::raiser::Get500Exn::f(err),
                Get500Error::b(err) => crate::services::raiser::Get500Exn::b(err),
                Get500Error::s(err) => crate::services::raiser::Get500Exn::s(err),
                Get500Error::ApplicationException(aexn) => crate::services::raiser::Get500Exn::ApplicationException(aexn),
                Get500Error::ThriftError(err) => crate::services::raiser::Get500Exn::ApplicationException(::fbthrift::ApplicationException {
                    message: err.to_string(),
                    type_: ::fbthrift::ApplicationExceptionErrorCode::InternalError,
                }),
            }
        }
    }

    pub(crate) enum Get500Reader {}

    impl ::fbthrift::help::DeserializeExn for Get500Reader {
        type Success = ::std::string::String;
        type Error = Get500Error;

        fn read_result<P>(p: &mut P) -> ::anyhow::Result<::std::result::Result<Self::Success, Self::Error>>
        where
            P: ::fbthrift::ProtocolReader,
        {
            static RETURNS: &[::fbthrift::Field] = &[
                ::fbthrift::Field::new("Success", ::fbthrift::TType::String, 0),
                ::fbthrift::Field::new("b", ::fbthrift::TType::Struct, 2),
                ::fbthrift::Field::new("f", ::fbthrift::TType::Struct, 1),
                ::fbthrift::Field::new("s", ::fbthrift::TType::Struct, 3),
            ];
            let _ = p.read_struct_begin(|_| ())?;
            let mut once = false;
            let mut alt = ::std::option::Option::None;
            loop {
                let (_, fty, fid) = p.read_field_begin(|_| (), RETURNS)?;
                match ((fty, fid as ::std::primitive::i32), once) {
                    ((::fbthrift::TType::Stop, _), _) => {
                        p.read_field_end()?;
                        break;
                    }
                    ((::fbthrift::TType::String, 0i32), false) => {
                        once = true;
                        alt = ::std::option::Option::Some(::std::result::Result::Ok(::fbthrift::Deserialize::read(p)?));
                    }
                    ((::fbthrift::TType::Struct, 1), false) => {
                        once = true;
                        alt = ::std::option::Option::Some(::std::result::Result::Err(Self::Error::f(::fbthrift::Deserialize::read(p)?)));
                    }
                    ((::fbthrift::TType::Struct, 2), false) => {
                        once = true;
                        alt = ::std::option::Option::Some(::std::result::Result::Err(Self::Error::b(::fbthrift::Deserialize::read(p)?)));
                    }
                    ((::fbthrift::TType::Struct, 3), false) => {
                        once = true;
                        alt = ::std::option::Option::Some(::std::result::Result::Err(Self::Error::s(::fbthrift::Deserialize::read(p)?)));
                    }
                    ((ty, _id), false) => p.skip(ty)?,
                    ((badty, badid), true) => return ::std::result::Result::Err(::std::convert::From::from(
                        ::fbthrift::ApplicationException::new(
                            ::fbthrift::ApplicationExceptionErrorCode::ProtocolError,
                            format!(
                                "unwanted extra union {} field ty {:?} id {}",
                                "Get500Error",
                                badty,
                                badid,
                            ),
                        )
                    )),
                }
                p.read_field_end()?;
            }
            p.read_struct_end()?;
            alt.ok_or_else(||
                ::fbthrift::ApplicationException::new(
                    ::fbthrift::ApplicationExceptionErrorCode::MissingResult,
                    format!("Empty union {}", "Get500Error"),
                )
                .into(),
            )
        }
    }

}

#[doc(inline)]
#[allow(ambiguous_glob_reexports)]
pub use self::raiser::*;

