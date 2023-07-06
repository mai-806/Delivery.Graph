#pragma once

#include <userver/clients/dns/component.hpp>
#include <userver/components/component.hpp>
#include <userver/storages/postgres/cluster.hpp>
#include <userver/storages/postgres/component.hpp>
#include <userver/server/handlers/http_handler_json_base.hpp>

#include <views/objects/objects.hpp>


namespace userver::formats::parse {

  deli_main::views::Coordinate Parse(const userver::formats::json::Value &elem,
                                     userver::formats::parse::To<deli_main::views::Coordinate>);

  deli_main::views::v1::order::post::OrderCreationRequest Parse(
          const userver::formats::json::Value &elem,
          userver::formats::parse::To<deli_main::views::v1::order::post::OrderCreationRequest>);
} // namespace userver::formats::parse

namespace userver::formats::serialize {

  json::Value Serialize(const deli_main::views::ErrorResponse &value,
                        serialize::To<json::Value>);

  json::Value Serialize(const deli_main::views::v1::order::post::OrderCreationResponse &value,
                        serialize::To<json::Value>);

} // namespace userver::formats::serialize
