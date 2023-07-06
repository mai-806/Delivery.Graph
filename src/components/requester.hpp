#pragma once

#include <userver/components/loggable_component_base.hpp>
#include <userver/dynamic_config/source.hpp>
#include <userver/storages/postgres/cluster.hpp>


namespace deli_main::components {


  class Requester : public userver::components::LoggableComponentBase {
  public:
    static constexpr auto kName = "requester";

    Requester(const userver::components::ComponentConfig &config,
              const userver::components::ComponentContext &context);

    template<class ReturnType, class... Args, class... QueryArgs>
    ReturnType DoDBQuery(ReturnType (*query)(Args...), QueryArgs... args) const {
      return query(cluster_, args...);
    }

  private:
    userver::storages::postgres::ClusterPtr cluster_;
  };


} // deli_main::components
