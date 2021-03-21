CREATE TABLE IF NOT EXISTS events (
  event_id serial PRIMARY KEY,
  name VARCHAR(240) NOT NULL,
  project_id int NOT NULL,
  timestamp timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT projects_project_id_foreign FOREIGN KEY (project_id) REFERENCES projects (project_id) 
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
);