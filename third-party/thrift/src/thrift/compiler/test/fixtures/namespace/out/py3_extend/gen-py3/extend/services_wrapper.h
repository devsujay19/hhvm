/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

#pragma once
#if __has_include(<thrift/compiler/test/fixtures/namespace/gen-cpp2/ExtendTestService.h>)
#include <thrift/compiler/test/fixtures/namespace/gen-cpp2/ExtendTestService.h>
#else
#include <thrift/compiler/test/fixtures/namespace/gen-cpp2/extend_handlers.h>
#endif
#include <thrift/compiler/test/fixtures/namespace/gen-py3/hsmodule/services_wrapper.h>
#include <folly/python/futures.h>
#include <Python.h>

#include <memory>

namespace cpp2 {

class ExtendTestServiceWrapper : public ::cpp2::HsTestServiceWrapper, virtual public ExtendTestServiceSvIf {
  public:
    explicit ExtendTestServiceWrapper(PyObject *if_object, folly::Executor *exc);
    void async_tm_check(apache::thrift::HandlerCallbackPtr<bool> callback
        , std::unique_ptr<::cpp2::HsFoo> struct1
    ) override;
folly::SemiFuture<folly::Unit> semifuture_onStartServing() override;
folly::SemiFuture<folly::Unit> semifuture_onStopRequested() override;
};

std::shared_ptr<apache::thrift::ServerInterface> ExtendTestServiceInterface(PyObject *if_object, folly::Executor *exc);
} // namespace cpp2
