CREATE SCHEMA IF NOT EXISTS delivery_service;
SET search_path = delivery_service;

CREATE TABLE IF NOT EXISTS deliveries (
    id TEXT PRIMARY KEY,
    order_id TEXT NOT NULL,
    delivery_person_id TEXT NOT NULL,
    status TEXT NOT NULL,
    current_location TEXT NOT NULL,
    assigned_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS delivery_persons (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    status TEXT NOT NULL
);
