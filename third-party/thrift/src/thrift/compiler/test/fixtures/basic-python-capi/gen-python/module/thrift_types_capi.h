
/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT
 *  @generated
 *
 */

#pragma once

#include <thrift/lib/python/capi/constructor.h>
#include <thrift/lib/python/capi/extractor.h>

#include <thrift/compiler/test/fixtures/basic-python-capi/src/gen-cpp2/module_types.h>

namespace apache {
namespace thrift {
namespace python {
namespace capi {
template <>
struct Extractor<::test::fixtures::basic-python-capi::MyStruct>
    : public BaseExtractor<::test::fixtures::basic-python-capi::MyStruct> {
  ExtractorResult<::test::fixtures::basic-python-capi::MyStruct> operator()(PyObject* obj);
  int typeCheck(PyObject* obj);
};

template <>
struct Constructor<::test::fixtures::basic-python-capi::MyStruct> {
  PyObject* operator()(::test::fixtures::basic-python-capi::MyStruct&& val);
};

template <>
struct Extractor<::test::fixtures::basic-python-capi::MyDataItem>
    : public BaseExtractor<::test::fixtures::basic-python-capi::MyDataItem> {
  ExtractorResult<::test::fixtures::basic-python-capi::MyDataItem> operator()(PyObject* obj);
  int typeCheck(PyObject* obj);
};

template <>
struct Constructor<::test::fixtures::basic-python-capi::MyDataItem> {
  PyObject* operator()(::test::fixtures::basic-python-capi::MyDataItem&& val);
};

template <>
struct Extractor<::test::fixtures::basic-python-capi::MyUnion>
    : public BaseExtractor<::test::fixtures::basic-python-capi::MyUnion> {
  ExtractorResult<::test::fixtures::basic-python-capi::MyUnion> operator()(PyObject* obj);
  int typeCheck(PyObject* obj);
};

template <>
struct Constructor<::test::fixtures::basic-python-capi::MyUnion> {
  PyObject* operator()(::test::fixtures::basic-python-capi::MyUnion&& val);
};

template <>
struct Extractor<::test::fixtures::basic-python-capi::MyEnum>
    : public BaseExtractor<::test::fixtures::basic-python-capi::MyEnum> {
  ExtractorResult<::test::fixtures::basic-python-capi::MyEnum> operator()(PyObject* obj);
  int typeCheck(PyObject* obj);
};

template <>
struct Constructor<::test::fixtures::basic-python-capi::MyEnum> {
  PyObject* operator()(::test::fixtures::basic-python-capi::MyEnum&& val);
};

} // namespace capi
} // namespace python
} // namespace thrift
} // namespace apache
