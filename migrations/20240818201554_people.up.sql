BEGIN;
-- migrate:up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE people (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    password VARCHAR(64),
    phone VARCHAR(20) NOT NULL,
    preferences JSONB,
    onduty BOOLEAN DEFAULT NULL,
    city VARCHAR(100) NOT NULL,
    state CHAR(2) NOT NULL,
    zipcode VARCHAR(10),
    created TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_people_uuid ON people (uuid);

COMMIT;