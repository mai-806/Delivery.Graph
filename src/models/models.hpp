#pragma once

#include <userver/formats/json/value.hpp>
#include <userver/storages/postgres/io/enum_types.hpp>
#include <userver/storages/postgres/io/json_types.hpp>


namespace deli_main::models {

  struct Coordinate {
    double latitude;
    double longitude;
  };

  enum class OrderStatus {
    kOrderStatusNew,
    kOrderStatusWaiting,
    kOrderStatusInProgress,
    kOrderStatusDelivered,
    kOrderStatusCanceled
  };

  struct Order {
    int64_t id;
    Coordinate start_point;
    Coordinate end_point;
    OrderStatus status;
    int64_t customer;
    std::optional<int64_t> courier;
  };


} // namespace deli_main::models



namespace userver::storages::postgres::io {

  template<>
  struct CppToUserPg<deli_main::models::Coordinate> {
    static constexpr DBTypeName postgres_name = "deli_main.coordinate_v1";
  };

  template<>
  struct CppToUserPg<deli_main::models::Order> {
    static constexpr DBTypeName postgres_name = "deli_main.order_v1";
  };

  template<>
  struct CppToUserPg<deli_main::models::OrderStatus>
          : EnumMappingBase<deli_main::models::OrderStatus> {
    static constexpr DBTypeName postgres_name = "deli_main.order_status";
    static constexpr EnumeratorList enumerators{
            {EnumType::kOrderStatusNew,        "new"},
            {EnumType::kOrderStatusWaiting,    "waiting"},
            {EnumType::kOrderStatusInProgress, "in_progress"},
            {EnumType::kOrderStatusDelivered,  "delivered"},
            {EnumType::kOrderStatusCanceled,   "canceled"}};
  };

} // namespace userver::storages::postgres::io
