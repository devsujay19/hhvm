
/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT
 *  @generated
 *
 */

#include <folly/python/import.h>
#include <thrift/lib/python/capi/iobuf.h>
#include <thrift/lib/python/types.h>

#include <thrift/compiler/test/fixtures/exceptions/src/gen-python-capi/module/thrift_types_api.h>
#include <thrift/compiler/test/fixtures/exceptions/src/gen-python-capi/module/thrift_types_capi.h>


namespace apache {
namespace thrift {
namespace python {
namespace capi {
namespace {
bool ensure_module_imported() {
  static ::folly::python::import_cache_nocapture import((
      ::import_module__thrift_types_capi));
  return import();
}
  static constexpr std::int16_t _fbthrift__Fiery__tuple_pos[1] = {
    1
  };
  static constexpr std::int16_t _fbthrift__Serious__tuple_pos[1] = {
    1
  };
  static constexpr std::int16_t _fbthrift__ComplexFieldNames__tuple_pos[2] = {
    1, 2
  };
  static constexpr std::int16_t _fbthrift__CustomFieldNames__tuple_pos[2] = {
    1, 2
  };
  static constexpr std::int16_t _fbthrift__ExceptionWithPrimitiveField__tuple_pos[2] = {
    1, 2
  };
  static constexpr std::int16_t _fbthrift__ExceptionWithStructuredAnnotation__tuple_pos[2] = {
    1, 2
  };
} // namespace

ExtractorResult<::cpp2::Fiery>
Extractor<::cpp2::Fiery>::operator()(PyObject* obj) {
  int tCheckResult = typeCheck(obj);
  if (tCheckResult != 1) {
      if (tCheckResult == 0) {
        PyErr_SetString(PyExc_TypeError, "Not a Fiery");
      }
      return extractorError<::cpp2::Fiery>(
          "Marshal error: Fiery");
  }
  StrongRef fbThriftData(getExceptionThriftData(obj));
  return Extractor<::apache::thrift::python::capi::ComposedStruct<
      ::cpp2::Fiery>>{}(*fbThriftData);
}

ExtractorResult<::cpp2::Fiery>
Extractor<::apache::thrift::python::capi::ComposedStruct<
    ::cpp2::Fiery>>::operator()(PyObject* fbThriftData) {
  ::cpp2::Fiery cpp;
  std::optional<std::string_view> error;
  Extractor<Bytes>{}.extractInto(
      cpp.message_ref(),
      PyTuple_GET_ITEM(fbThriftData, _fbthrift__Fiery__tuple_pos[0]),
      error);
  if (error) {
    return folly::makeUnexpected(*error);
  }
  return cpp;
}


int Extractor<::cpp2::Fiery>::typeCheck(PyObject* obj) {
  if (!ensure_module_imported()) {
    ::folly::python::handlePythonError(
      "Module module import error");
  }
  int result =
      can_extract__module__Fiery(obj);
  if (result < 0) {
    ::folly::python::handlePythonError(
      "Unexpected type check error: Fiery");
  }
  return result;
}


PyObject* Constructor<::cpp2::Fiery>::operator()(
    const ::cpp2::Fiery& val) {
  if (!ensure_module_imported()) {
    DCHECK(PyErr_Occurred() != nullptr);
    return nullptr;
  }
  Constructor<::apache::thrift::python::capi::ComposedStruct<
        ::cpp2::Fiery>> ctor;
  StrongRef fbthrift_data(ctor(val));
  if (!fbthrift_data) {
    return nullptr;
  }
  return init__module__Fiery(*fbthrift_data);
}

PyObject* Constructor<::apache::thrift::python::capi::ComposedStruct<
        ::cpp2::Fiery>>::operator()(
    FOLLY_MAYBE_UNUSED const ::cpp2::Fiery& val) {
  StrongRef fbthrift_data(createStructTuple(1));
  StrongRef _fbthrift__message(
    Constructor<Bytes>{}
    .constructFrom(val.message_ref()));
  if (!_fbthrift__message ||
      setStructField(
          *fbthrift_data,
          _fbthrift__Fiery__tuple_pos[0],
          *_fbthrift__message) == -1) {
    return nullptr;
  }
  return std::move(fbthrift_data).release();
}


ExtractorResult<::cpp2::Serious>
Extractor<::cpp2::Serious>::operator()(PyObject* obj) {
  int tCheckResult = typeCheck(obj);
  if (tCheckResult != 1) {
      if (tCheckResult == 0) {
        PyErr_SetString(PyExc_TypeError, "Not a Serious");
      }
      return extractorError<::cpp2::Serious>(
          "Marshal error: Serious");
  }
  StrongRef fbThriftData(getExceptionThriftData(obj));
  return Extractor<::apache::thrift::python::capi::ComposedStruct<
      ::cpp2::Serious>>{}(*fbThriftData);
}

ExtractorResult<::cpp2::Serious>
Extractor<::apache::thrift::python::capi::ComposedStruct<
    ::cpp2::Serious>>::operator()(PyObject* fbThriftData) {
  ::cpp2::Serious cpp;
  std::optional<std::string_view> error;
  Extractor<Bytes>{}.extractInto(
      cpp.sonnet_ref(),
      PyTuple_GET_ITEM(fbThriftData, _fbthrift__Serious__tuple_pos[0]),
      error);
  if (error) {
    return folly::makeUnexpected(*error);
  }
  return cpp;
}


int Extractor<::cpp2::Serious>::typeCheck(PyObject* obj) {
  if (!ensure_module_imported()) {
    ::folly::python::handlePythonError(
      "Module module import error");
  }
  int result =
      can_extract__module__Serious(obj);
  if (result < 0) {
    ::folly::python::handlePythonError(
      "Unexpected type check error: Serious");
  }
  return result;
}


PyObject* Constructor<::cpp2::Serious>::operator()(
    const ::cpp2::Serious& val) {
  if (!ensure_module_imported()) {
    DCHECK(PyErr_Occurred() != nullptr);
    return nullptr;
  }
  Constructor<::apache::thrift::python::capi::ComposedStruct<
        ::cpp2::Serious>> ctor;
  StrongRef fbthrift_data(ctor(val));
  if (!fbthrift_data) {
    return nullptr;
  }
  return init__module__Serious(*fbthrift_data);
}

PyObject* Constructor<::apache::thrift::python::capi::ComposedStruct<
        ::cpp2::Serious>>::operator()(
    FOLLY_MAYBE_UNUSED const ::cpp2::Serious& val) {
  StrongRef fbthrift_data(createStructTuple(1));
  StrongRef _fbthrift__sonnet(
    Constructor<Bytes>{}
    .constructFrom(val.sonnet_ref()));
  if (_fbthrift__sonnet.isNone()) {
    Py_INCREF(Py_None);
    PyTuple_SET_ITEM(
      *fbthrift_data,
      _fbthrift__Serious__tuple_pos[0],
      Py_None);
  } else
  if (!_fbthrift__sonnet ||
      setStructField(
          *fbthrift_data,
          _fbthrift__Serious__tuple_pos[0],
          *_fbthrift__sonnet) == -1) {
    return nullptr;
  }
  return std::move(fbthrift_data).release();
}


ExtractorResult<::cpp2::ComplexFieldNames>
Extractor<::cpp2::ComplexFieldNames>::operator()(PyObject* obj) {
  int tCheckResult = typeCheck(obj);
  if (tCheckResult != 1) {
      if (tCheckResult == 0) {
        PyErr_SetString(PyExc_TypeError, "Not a ComplexFieldNames");
      }
      return extractorError<::cpp2::ComplexFieldNames>(
          "Marshal error: ComplexFieldNames");
  }
  StrongRef fbThriftData(getExceptionThriftData(obj));
  return Extractor<::apache::thrift::python::capi::ComposedStruct<
      ::cpp2::ComplexFieldNames>>{}(*fbThriftData);
}

ExtractorResult<::cpp2::ComplexFieldNames>
Extractor<::apache::thrift::python::capi::ComposedStruct<
    ::cpp2::ComplexFieldNames>>::operator()(PyObject* fbThriftData) {
  ::cpp2::ComplexFieldNames cpp;
  std::optional<std::string_view> error;
  Extractor<Bytes>{}.extractInto(
      cpp.error_message_ref(),
      PyTuple_GET_ITEM(fbThriftData, _fbthrift__ComplexFieldNames__tuple_pos[0]),
      error);
  Extractor<Bytes>{}.extractInto(
      cpp.internal_error_message_ref(),
      PyTuple_GET_ITEM(fbThriftData, _fbthrift__ComplexFieldNames__tuple_pos[1]),
      error);
  if (error) {
    return folly::makeUnexpected(*error);
  }
  return cpp;
}


int Extractor<::cpp2::ComplexFieldNames>::typeCheck(PyObject* obj) {
  if (!ensure_module_imported()) {
    ::folly::python::handlePythonError(
      "Module module import error");
  }
  int result =
      can_extract__module__ComplexFieldNames(obj);
  if (result < 0) {
    ::folly::python::handlePythonError(
      "Unexpected type check error: ComplexFieldNames");
  }
  return result;
}


PyObject* Constructor<::cpp2::ComplexFieldNames>::operator()(
    const ::cpp2::ComplexFieldNames& val) {
  if (!ensure_module_imported()) {
    DCHECK(PyErr_Occurred() != nullptr);
    return nullptr;
  }
  Constructor<::apache::thrift::python::capi::ComposedStruct<
        ::cpp2::ComplexFieldNames>> ctor;
  StrongRef fbthrift_data(ctor(val));
  if (!fbthrift_data) {
    return nullptr;
  }
  return init__module__ComplexFieldNames(*fbthrift_data);
}

PyObject* Constructor<::apache::thrift::python::capi::ComposedStruct<
        ::cpp2::ComplexFieldNames>>::operator()(
    FOLLY_MAYBE_UNUSED const ::cpp2::ComplexFieldNames& val) {
  StrongRef fbthrift_data(createStructTuple(2));
  StrongRef _fbthrift__error_message(
    Constructor<Bytes>{}
    .constructFrom(val.error_message_ref()));
  if (!_fbthrift__error_message ||
      setStructField(
          *fbthrift_data,
          _fbthrift__ComplexFieldNames__tuple_pos[0],
          *_fbthrift__error_message) == -1) {
    return nullptr;
  }
  StrongRef _fbthrift__internal_error_message(
    Constructor<Bytes>{}
    .constructFrom(val.internal_error_message_ref()));
  if (!_fbthrift__internal_error_message ||
      setStructField(
          *fbthrift_data,
          _fbthrift__ComplexFieldNames__tuple_pos[1],
          *_fbthrift__internal_error_message) == -1) {
    return nullptr;
  }
  return std::move(fbthrift_data).release();
}


ExtractorResult<::cpp2::CustomFieldNames>
Extractor<::cpp2::CustomFieldNames>::operator()(PyObject* obj) {
  int tCheckResult = typeCheck(obj);
  if (tCheckResult != 1) {
      if (tCheckResult == 0) {
        PyErr_SetString(PyExc_TypeError, "Not a CustomFieldNames");
      }
      return extractorError<::cpp2::CustomFieldNames>(
          "Marshal error: CustomFieldNames");
  }
  StrongRef fbThriftData(getExceptionThriftData(obj));
  return Extractor<::apache::thrift::python::capi::ComposedStruct<
      ::cpp2::CustomFieldNames>>{}(*fbThriftData);
}

ExtractorResult<::cpp2::CustomFieldNames>
Extractor<::apache::thrift::python::capi::ComposedStruct<
    ::cpp2::CustomFieldNames>>::operator()(PyObject* fbThriftData) {
  ::cpp2::CustomFieldNames cpp;
  std::optional<std::string_view> error;
  Extractor<Bytes>{}.extractInto(
      cpp.error_message_ref(),
      PyTuple_GET_ITEM(fbThriftData, _fbthrift__CustomFieldNames__tuple_pos[0]),
      error);
  Extractor<Bytes>{}.extractInto(
      cpp.internal_error_message_ref(),
      PyTuple_GET_ITEM(fbThriftData, _fbthrift__CustomFieldNames__tuple_pos[1]),
      error);
  if (error) {
    return folly::makeUnexpected(*error);
  }
  return cpp;
}


int Extractor<::cpp2::CustomFieldNames>::typeCheck(PyObject* obj) {
  if (!ensure_module_imported()) {
    ::folly::python::handlePythonError(
      "Module module import error");
  }
  int result =
      can_extract__module__CustomFieldNames(obj);
  if (result < 0) {
    ::folly::python::handlePythonError(
      "Unexpected type check error: CustomFieldNames");
  }
  return result;
}


PyObject* Constructor<::cpp2::CustomFieldNames>::operator()(
    const ::cpp2::CustomFieldNames& val) {
  if (!ensure_module_imported()) {
    DCHECK(PyErr_Occurred() != nullptr);
    return nullptr;
  }
  Constructor<::apache::thrift::python::capi::ComposedStruct<
        ::cpp2::CustomFieldNames>> ctor;
  StrongRef fbthrift_data(ctor(val));
  if (!fbthrift_data) {
    return nullptr;
  }
  return init__module__CustomFieldNames(*fbthrift_data);
}

PyObject* Constructor<::apache::thrift::python::capi::ComposedStruct<
        ::cpp2::CustomFieldNames>>::operator()(
    FOLLY_MAYBE_UNUSED const ::cpp2::CustomFieldNames& val) {
  StrongRef fbthrift_data(createStructTuple(2));
  StrongRef _fbthrift__error_message(
    Constructor<Bytes>{}
    .constructFrom(val.error_message_ref()));
  if (!_fbthrift__error_message ||
      setStructField(
          *fbthrift_data,
          _fbthrift__CustomFieldNames__tuple_pos[0],
          *_fbthrift__error_message) == -1) {
    return nullptr;
  }
  StrongRef _fbthrift__internal_error_message(
    Constructor<Bytes>{}
    .constructFrom(val.internal_error_message_ref()));
  if (!_fbthrift__internal_error_message ||
      setStructField(
          *fbthrift_data,
          _fbthrift__CustomFieldNames__tuple_pos[1],
          *_fbthrift__internal_error_message) == -1) {
    return nullptr;
  }
  return std::move(fbthrift_data).release();
}


ExtractorResult<::cpp2::ExceptionWithPrimitiveField>
Extractor<::cpp2::ExceptionWithPrimitiveField>::operator()(PyObject* obj) {
  int tCheckResult = typeCheck(obj);
  if (tCheckResult != 1) {
      if (tCheckResult == 0) {
        PyErr_SetString(PyExc_TypeError, "Not a ExceptionWithPrimitiveField");
      }
      return extractorError<::cpp2::ExceptionWithPrimitiveField>(
          "Marshal error: ExceptionWithPrimitiveField");
  }
  StrongRef fbThriftData(getExceptionThriftData(obj));
  return Extractor<::apache::thrift::python::capi::ComposedStruct<
      ::cpp2::ExceptionWithPrimitiveField>>{}(*fbThriftData);
}

ExtractorResult<::cpp2::ExceptionWithPrimitiveField>
Extractor<::apache::thrift::python::capi::ComposedStruct<
    ::cpp2::ExceptionWithPrimitiveField>>::operator()(PyObject* fbThriftData) {
  ::cpp2::ExceptionWithPrimitiveField cpp;
  std::optional<std::string_view> error;
  Extractor<Bytes>{}.extractInto(
      cpp.message_ref(),
      PyTuple_GET_ITEM(fbThriftData, _fbthrift__ExceptionWithPrimitiveField__tuple_pos[0]),
      error);
  Extractor<int32_t>{}.extractInto(
      cpp.error_code_ref(),
      PyTuple_GET_ITEM(fbThriftData, _fbthrift__ExceptionWithPrimitiveField__tuple_pos[1]),
      error);
  if (error) {
    return folly::makeUnexpected(*error);
  }
  return cpp;
}


int Extractor<::cpp2::ExceptionWithPrimitiveField>::typeCheck(PyObject* obj) {
  if (!ensure_module_imported()) {
    ::folly::python::handlePythonError(
      "Module module import error");
  }
  int result =
      can_extract__module__ExceptionWithPrimitiveField(obj);
  if (result < 0) {
    ::folly::python::handlePythonError(
      "Unexpected type check error: ExceptionWithPrimitiveField");
  }
  return result;
}


PyObject* Constructor<::cpp2::ExceptionWithPrimitiveField>::operator()(
    const ::cpp2::ExceptionWithPrimitiveField& val) {
  if (!ensure_module_imported()) {
    DCHECK(PyErr_Occurred() != nullptr);
    return nullptr;
  }
  Constructor<::apache::thrift::python::capi::ComposedStruct<
        ::cpp2::ExceptionWithPrimitiveField>> ctor;
  StrongRef fbthrift_data(ctor(val));
  if (!fbthrift_data) {
    return nullptr;
  }
  return init__module__ExceptionWithPrimitiveField(*fbthrift_data);
}

PyObject* Constructor<::apache::thrift::python::capi::ComposedStruct<
        ::cpp2::ExceptionWithPrimitiveField>>::operator()(
    FOLLY_MAYBE_UNUSED const ::cpp2::ExceptionWithPrimitiveField& val) {
  StrongRef fbthrift_data(createStructTuple(2));
  StrongRef _fbthrift__message(
    Constructor<Bytes>{}
    .constructFrom(val.message_ref()));
  if (!_fbthrift__message ||
      setStructField(
          *fbthrift_data,
          _fbthrift__ExceptionWithPrimitiveField__tuple_pos[0],
          *_fbthrift__message) == -1) {
    return nullptr;
  }
  StrongRef _fbthrift__error_code(
    Constructor<int32_t>{}
    .constructFrom(val.error_code_ref()));
  if (!_fbthrift__error_code ||
      setStructField(
          *fbthrift_data,
          _fbthrift__ExceptionWithPrimitiveField__tuple_pos[1],
          *_fbthrift__error_code) == -1) {
    return nullptr;
  }
  return std::move(fbthrift_data).release();
}


ExtractorResult<::cpp2::ExceptionWithStructuredAnnotation>
Extractor<::cpp2::ExceptionWithStructuredAnnotation>::operator()(PyObject* obj) {
  int tCheckResult = typeCheck(obj);
  if (tCheckResult != 1) {
      if (tCheckResult == 0) {
        PyErr_SetString(PyExc_TypeError, "Not a ExceptionWithStructuredAnnotation");
      }
      return extractorError<::cpp2::ExceptionWithStructuredAnnotation>(
          "Marshal error: ExceptionWithStructuredAnnotation");
  }
  StrongRef fbThriftData(getExceptionThriftData(obj));
  return Extractor<::apache::thrift::python::capi::ComposedStruct<
      ::cpp2::ExceptionWithStructuredAnnotation>>{}(*fbThriftData);
}

ExtractorResult<::cpp2::ExceptionWithStructuredAnnotation>
Extractor<::apache::thrift::python::capi::ComposedStruct<
    ::cpp2::ExceptionWithStructuredAnnotation>>::operator()(PyObject* fbThriftData) {
  ::cpp2::ExceptionWithStructuredAnnotation cpp;
  std::optional<std::string_view> error;
  Extractor<Bytes>{}.extractInto(
      cpp.message_field_ref(),
      PyTuple_GET_ITEM(fbThriftData, _fbthrift__ExceptionWithStructuredAnnotation__tuple_pos[0]),
      error);
  Extractor<int32_t>{}.extractInto(
      cpp.error_code_ref(),
      PyTuple_GET_ITEM(fbThriftData, _fbthrift__ExceptionWithStructuredAnnotation__tuple_pos[1]),
      error);
  if (error) {
    return folly::makeUnexpected(*error);
  }
  return cpp;
}


int Extractor<::cpp2::ExceptionWithStructuredAnnotation>::typeCheck(PyObject* obj) {
  if (!ensure_module_imported()) {
    ::folly::python::handlePythonError(
      "Module module import error");
  }
  int result =
      can_extract__module__ExceptionWithStructuredAnnotation(obj);
  if (result < 0) {
    ::folly::python::handlePythonError(
      "Unexpected type check error: ExceptionWithStructuredAnnotation");
  }
  return result;
}


PyObject* Constructor<::cpp2::ExceptionWithStructuredAnnotation>::operator()(
    const ::cpp2::ExceptionWithStructuredAnnotation& val) {
  if (!ensure_module_imported()) {
    DCHECK(PyErr_Occurred() != nullptr);
    return nullptr;
  }
  Constructor<::apache::thrift::python::capi::ComposedStruct<
        ::cpp2::ExceptionWithStructuredAnnotation>> ctor;
  StrongRef fbthrift_data(ctor(val));
  if (!fbthrift_data) {
    return nullptr;
  }
  return init__module__ExceptionWithStructuredAnnotation(*fbthrift_data);
}

PyObject* Constructor<::apache::thrift::python::capi::ComposedStruct<
        ::cpp2::ExceptionWithStructuredAnnotation>>::operator()(
    FOLLY_MAYBE_UNUSED const ::cpp2::ExceptionWithStructuredAnnotation& val) {
  StrongRef fbthrift_data(createStructTuple(2));
  StrongRef _fbthrift__message_field(
    Constructor<Bytes>{}
    .constructFrom(val.message_field_ref()));
  if (!_fbthrift__message_field ||
      setStructField(
          *fbthrift_data,
          _fbthrift__ExceptionWithStructuredAnnotation__tuple_pos[0],
          *_fbthrift__message_field) == -1) {
    return nullptr;
  }
  StrongRef _fbthrift__error_code(
    Constructor<int32_t>{}
    .constructFrom(val.error_code_ref()));
  if (!_fbthrift__error_code ||
      setStructField(
          *fbthrift_data,
          _fbthrift__ExceptionWithStructuredAnnotation__tuple_pos[1],
          *_fbthrift__error_code) == -1) {
    return nullptr;
  }
  return std::move(fbthrift_data).release();
}


ExtractorResult<::cpp2::Banal>
Extractor<::cpp2::Banal>::operator()(PyObject* obj) {
  int tCheckResult = typeCheck(obj);
  if (tCheckResult != 1) {
      if (tCheckResult == 0) {
        PyErr_SetString(PyExc_TypeError, "Not a Banal");
      }
      return extractorError<::cpp2::Banal>(
          "Marshal error: Banal");
  }
  return ::cpp2::Banal{};
}

ExtractorResult<::cpp2::Banal>
Extractor<::apache::thrift::python::capi::ComposedStruct<
    ::cpp2::Banal>>::operator()(PyObject* fbThriftData) {
  ::cpp2::Banal cpp;
  (void)fbThriftData;
  return cpp;
}


int Extractor<::cpp2::Banal>::typeCheck(PyObject* obj) {
  if (!ensure_module_imported()) {
    ::folly::python::handlePythonError(
      "Module module import error");
  }
  int result =
      can_extract__module__Banal(obj);
  if (result < 0) {
    ::folly::python::handlePythonError(
      "Unexpected type check error: Banal");
  }
  return result;
}


PyObject* Constructor<::cpp2::Banal>::operator()(
    const ::cpp2::Banal& val) {
  if (!ensure_module_imported()) {
    DCHECK(PyErr_Occurred() != nullptr);
    return nullptr;
  }
  Constructor<::apache::thrift::python::capi::ComposedStruct<
        ::cpp2::Banal>> ctor;
  StrongRef fbthrift_data(ctor(val));
  if (!fbthrift_data) {
    return nullptr;
  }
  return init__module__Banal(*fbthrift_data);
}

PyObject* Constructor<::apache::thrift::python::capi::ComposedStruct<
        ::cpp2::Banal>>::operator()(
    FOLLY_MAYBE_UNUSED const ::cpp2::Banal& val) {
  StrongRef fbthrift_data(createStructTuple(0));
  return std::move(fbthrift_data).release();
}


} // namespace capi
} // namespace python
} // namespace thrift
} // namespace apache
