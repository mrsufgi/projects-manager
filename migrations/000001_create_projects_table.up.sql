CREATE TABLE IF NOT EXISTS projects (
  project_id serial PRIMARY KEY,
  name VARCHAR(240) NOT NULL UNIQUE,
  vertical VARCHAR(240) NOT NULL,
  event VARCHAR(240) NOT NULL,
  url VARCHAR(240) NOT NULL UNIQUE,
  credentials JSON DEFAULT '{}'::json,
  created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX projects_url_index ON projects (url);

CREATE OR REPLACE FUNCTION update_timestamp() RETURNS TRIGGER
  LANGUAGE plpgsql AS
    $$
      BEGIN
          NEW.updated_at = CURRENT_TIMESTAMP;
          RETURN NEW;
      END;
    $$;

CREATE TRIGGER update_project_timestamp
BEFORE UPDATE ON projects
FOR EACH ROW EXECUTE PROCEDURE update_timestamp();