-- Up Migration
BEGIN;

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    message TEXT NOT NULL,
    city VARCHAR(255),
    state CHAR(2),
    zipcode CHAR(5),
    caller VARCHAR(255) NOT NULL,
    contact VARCHAR(255),
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    updated TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted TIMESTAMP
);

-- Create index on uuid
CREATE INDEX idx_messages_uuid ON messages (uuid);

COMMIT;