CREATE TABLE IF NOT EXISTS events (
  event_id serial PRIMARY KEY,
  name VARCHAR(240) NOT NULL,
  timestamp timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
);