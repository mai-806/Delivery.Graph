#include "requester.hpp"

#include <userver/components/component_config.hpp>
#include <userver/components/component_context.hpp>
#include <userver/components/statistics_storage.hpp>
#include <userver/dynamic_config/storage/component.hpp>
#include <userver/storages/postgres/component.hpp>

#include <common/consts.hpp>


namespace deli_main::components {

    Requester::Requester(const userver::components::ComponentConfig& config,
                             const userver::components::ComponentContext& context)
            : LoggableComponentBase(config, context),
              cluster_(context.FindComponent<userver::components::Postgres>(common::consts::kPgClusterName)
                               .GetCluster()) {}

} // namespace deli_main::components
