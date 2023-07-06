DROP SCHEMA IF EXISTS deli_main CASCADE;

CREATE SCHEMA IF NOT EXISTS deli_main;

CREATE TYPE deli_main.order_status AS ENUM (
    'new',
    'waiting',
    'in_progress',
    'delivered',
    'canceled'
    );

CREATE TYPE deli_main.coordinate_v1 AS
(
    latitude  FLOAT,
    longitude FLOAT
);

CREATE TABLE IF NOT EXISTS deli_main.orders
(
    id          BIGSERIAL PRIMARY KEY,
    start_point deli_main.coordinate_v1 NOT NULL,
    end_point   deli_main.coordinate_v1 NOT NULL,
    status      deli_main.order_status DEFAULT 'new',
    customer    BIGINT                  NOT NULL,
    courier     BIGINT,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS ix_deli_main_orders_customer ON deli_main.orders (customer);
CREATE INDEX IF NOT EXISTS ix_deli_main_orders_courier ON deli_main.orders (courier);


CREATE TYPE deli_main.order_v1 AS
(
    id          BIGINT,
    start_point deli_main.coordinate_v1,
    end_point   deli_main.coordinate_v1,
    status      deli_main.order_status,
    customer    BIGINT,
    courier     BIGINT
);
