#pragma once

#include <userver/clients/dns/component.hpp>
#include <userver/components/component.hpp>
#include <userver/storages/postgres/cluster.hpp>
#include <userver/storages/postgres/component.hpp>
#include <userver/server/handlers/http_handler_json_base.hpp>

#include <views/objects/parsers.hpp>

#include <components/requester.hpp>

namespace deli_main::views::v1::order::post {

  class Handler : public userver::server::handlers::HttpHandlerJsonBase {
  public:
    static constexpr std::string_view kName = "v1-order-post-handler";

    using Request = OrderCreationRequest;
    using Response400 = ErrorResponse;
    using Response200 = OrderCreationResponse;

    Handler(const userver::components::ComponentConfig &config,
            const userver::components::ComponentContext &component_context);

    userver::formats::json::Value HandleRequestJsonThrow(
            const userver::server::http::HttpRequest &, const userver::formats::json::Value &json,
            userver::server::request::RequestContext &) const override;

  private:
    const components::Requester &requester_;
  };
} // namespace deli_main::views::v1::order::post
