/*
 *  Copyright (c) Meta Platforms, Inc. and affiliates.
 *
 *  This source code is licensed under the MIT license found in the LICENSE
 *  file in the root directory of this source tree.
 *
 */

/*
 *  THIS FILE IS AUTOGENERATED. DO NOT MODIFY IT; ALL CHANGES WILL BE LOST IN
 *  VAIN.
 *
 *  @generated
 */
#pragma once

#include "mcrouter/lib/carbon/test/gen/gen-cpp2/B_types.h"
#include "mcrouter/lib/carbon/test/gen/gen-cpp2/B.h"

namespace carbon {
namespace test {
namespace B {

template <class ServerOnRequest>
class BServerOnRequestThrift : public thrift::BSvIf {
 public:
  BServerOnRequestThrift(
      std::unordered_map<
          folly::EventBase*,
          std::shared_ptr<ServerOnRequest>> serverOnRequestMap)
      : serverOnRequestMap_(std::move(serverOnRequestMap)) {}
  void async_eb_testB(
      apache::thrift::HandlerCallbackPtr<carbon::test::B::TestBReply> callback,
      const carbon::test::B::TestBRequest& request) override final {
    onRequestThriftHelper<std::remove_reference_t<decltype(request)>>(
        request, std::move(callback));
  }
  void async_eb_mcVersion(
      apache::thrift::HandlerCallbackPtr<facebook::memcache::McVersionReply> callback,
      const facebook::memcache::McVersionRequest& request) override final {
    onRequestThriftHelper<std::remove_reference_t<decltype(request)>>(
        request, std::move(callback));
  }
 // Return this factory instead of BAsyncProcessor from getProcessor(),
 // so that we don't use the default statically registered handlers
 class BAsyncProcessorCustomHandlers
     : public thrift::BAsyncProcessor {
  public:
   explicit BAsyncProcessorCustomHandlers(thrift::BSvIf* svif)
        : BAsyncProcessor(svif) {
     clearEventHandlers();
   }
 };
 std::unique_ptr<apache::thrift::AsyncProcessor> getProcessor()
     override final {
   return std::make_unique<BAsyncProcessorCustomHandlers>(this);
 }

 private:
  std::unordered_map<
     folly::EventBase*,
     std::shared_ptr<ServerOnRequest>>
     serverOnRequestMap_;
  static thread_local ServerOnRequest* serverOnRequest_;

  template <class Request>
   void onRequestThriftHelper(
       const Request& request,
       apache::thrift::HandlerCallbackPtr<typename Request::reply_type>
           callback) {
     getServerOnRequest(callback->getEventBase())
         ->onRequestThrift(
             std::move(callback),
             std::move(const_cast<std::remove_const_t<Request>&>(request)));
   }

 // Returns the ServerOnRequest* associated with this evb
  ServerOnRequest* getServerOnRequest(folly::EventBase* evb) {
    if (serverOnRequest_ == nullptr) {
      auto it = serverOnRequestMap_.find(evb);
      CHECK(it != serverOnRequestMap_.end());
      serverOnRequest_ = it->second.get();
    }
    return serverOnRequest_;
  }
};

template <class ServerOnRequest>
thread_local ServerOnRequest*
    BServerOnRequestThrift<ServerOnRequest>::serverOnRequest_{nullptr};

} // namespace B
} // namespace test
} // namespace carbon