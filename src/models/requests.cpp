#include "requests.hpp"

#include <models/sql/all.hpp>

namespace deli_main::models::requests {

  int64_t InsertOrder(const userver::storages::postgres::ClusterPtr& cluster,
                      const models::Order &order) {
    const auto &result =
            cluster->Execute(userver::storages::postgres::ClusterHostType::kMaster,
                             sql::kInsertOrder,
                             order);
    return result.AsSingleRow<int64_t>();
  }

} // namespace deli_main::models::requests
