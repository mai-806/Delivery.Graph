#pragma once


namespace deli_main::models::sql {
  const constexpr char *kInsertOrder = "INSERT INTO deli_main.orders "
                                       "(start_point, end_point, status, customer, created_at, updated_at) "
                                       "VALUES "
                                       "($1.start_point, $1.end_point, $1.status, $1.customer, "
                                       "CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) "
                                       "RETURNING id;";
} // namespace deli_main::models::sql
