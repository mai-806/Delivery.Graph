#pragma once

#include <cinttypes>
#include <string>

namespace deli_main::views {

  struct Coordinate {
    double lon;
    double lat;
  };

  struct ErrorResponse {
    std::string message;
  };

  namespace v1::order::post {
    struct OrderCreationRequest {
      int64_t customer_id;
      Coordinate start;
      Coordinate finish;
    };

  struct OrderCreationResponse {
    int64_t order_id;
  };
  }

} // namespace deli_main::views
