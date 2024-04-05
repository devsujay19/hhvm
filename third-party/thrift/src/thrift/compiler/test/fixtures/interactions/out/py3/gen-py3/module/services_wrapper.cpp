/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

#include <thrift/compiler/test/fixtures/interactions/gen-py3/module/services_wrapper.h>
#include <thrift/compiler/test/fixtures/interactions/gen-py3/module/services_api.h>
#include <thrift/lib/cpp2/async/AsyncProcessor.h>

namespace cpp2 {

MyServiceWrapper::MyServiceWrapper(PyObject *obj, folly::Executor* exc)
  : if_object(obj), executor(exc)
  {
    import_test__fixtures__interactions__module__services();
  }


void MyServiceWrapper::async_tm_foo(
  std::unique_ptr<apache::thrift::HandlerCallback<void>> callback) {
  auto ctx = callback->getRequestContext();
  folly::via(
    this->executor,
    [this, ctx,
     callback = std::move(callback)    ]() mutable {
        auto [promise, future] = folly::makePromiseContract<folly::Unit>();
        call_cy_MyService_foo(
            this->if_object,
            ctx,
            std::move(promise)        );
        std::move(future).via(this->executor).thenTry([callback = std::move(callback)](folly::Try<folly::Unit>&& t) {
          (void)t;
          callback->complete(std::move(t));
        });
    });
}
void MyServiceWrapper::async_tm_interact(
  std::unique_ptr<apache::thrift::HandlerCallback<void>> callback
    , int32_t arg
) {
  auto ctx = callback->getRequestContext();
  folly::via(
    this->executor,
    [this, ctx,
     callback = std::move(callback),
arg    ]() mutable {
        auto [promise, future] = folly::makePromiseContract<folly::Unit>();
        call_cy_MyService_interact(
            this->if_object,
            ctx,
            std::move(promise),
            arg        );
        std::move(future).via(this->executor).thenTry([callback = std::move(callback)](folly::Try<folly::Unit>&& t) {
          (void)t;
          callback->complete(std::move(t));
        });
    });
}
void MyServiceWrapper::async_tm_interactFast(
  std::unique_ptr<apache::thrift::HandlerCallback<int32_t>> callback) {
  auto ctx = callback->getRequestContext();
  folly::via(
    this->executor,
    [this, ctx,
     callback = std::move(callback)    ]() mutable {
        auto [promise, future] = folly::makePromiseContract<int32_t>();
        call_cy_MyService_interactFast(
            this->if_object,
            ctx,
            std::move(promise)        );
        std::move(future).via(this->executor).thenTry([callback = std::move(callback)](folly::Try<int32_t>&& t) {
          (void)t;
          callback->complete(std::move(t));
        });
    });
}
void MyServiceWrapper::async_tm_serialize(
  std::unique_ptr<apache::thrift::HandlerCallback<apache::thrift::ResponseAndServerStream<int32_t,int32_t>>> callback) {
  auto ctx = callback->getRequestContext();
  folly::via(
    this->executor,
    [this, ctx,
     callback = std::move(callback)    ]() mutable {
        auto [promise, future] = folly::makePromiseContract<apache::thrift::ResponseAndServerStream<int32_t,int32_t>>();
        call_cy_MyService_serialize(
            this->if_object,
            ctx,
            std::move(promise)        );
        std::move(future).via(this->executor).thenTry([callback = std::move(callback)](folly::Try<apache::thrift::ResponseAndServerStream<int32_t,int32_t>>&& t) {
          (void)t;
          callback->complete(std::move(t));
        });
    });
}
std::unique_ptr<MyServiceSvIf::MyInteractionIf> MyServiceWrapper::createMyInteraction() {
  throw std::runtime_error("Py3 server doesn't support interactions.");
}
std::unique_ptr<MyServiceSvIf::MyInteractionFastIf> MyServiceWrapper::createMyInteractionFast() {
  throw std::runtime_error("Py3 server doesn't support interactions.");
}
std::unique_ptr<MyServiceSvIf::SerialInteractionIf> MyServiceWrapper::createSerialInteraction() {
  throw std::runtime_error("Py3 server doesn't support interactions.");
}
std::shared_ptr<apache::thrift::ServerInterface> MyServiceInterface(PyObject *if_object, folly::Executor *exc) {
  return std::make_shared<MyServiceWrapper>(if_object, exc);
}
folly::SemiFuture<folly::Unit> MyServiceWrapper::semifuture_onStartServing() {
  auto [promise, future] = folly::makePromiseContract<folly::Unit>();
  call_cy_MyService_onStartServing(
      this->if_object,
      std::move(promise)
  );
  return std::move(future);
}
folly::SemiFuture<folly::Unit> MyServiceWrapper::semifuture_onStopRequested() {
  auto [promise, future] = folly::makePromiseContract<folly::Unit>();
  call_cy_MyService_onStopRequested(
      this->if_object,
      std::move(promise)
  );
  return std::move(future);
}
} // namespace cpp2
