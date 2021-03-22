package domain

import (
	"encoding/json"
	"time"
)

type Project struct {
	ID          int              `db:"project_id" json:"id"`
	Name        *string          `db:"name" json:"name"`
	Vertical    *string          `db:"vertical" json:"vertical"`
	Event       *string          `db:"event" json:"event"`
	URL         *string          `db:"url" json:"url"`
	Credentials *json.RawMessage `db:"credentials" json:"credentials,omitempty"`
	CreatedAt   *time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time       `db:"updated_at" json:"updated_at"`
}
type SearchProjectsInput struct {
	Name string
	URL  string
}

//go:generate mockgen -destination=mocks/mock_projects_repository.go -package=mocks . ProjectsRepository
type ProjectsRepository interface {
	SearchProjects(p SearchProjectsInput) (*[]Project, error)
	CreateProject(project Project) (int, error)
	ReadProject(id int) (*Project, error)
	UpdateProject(id int, project Project) (int64, error)
	DeleteProject(id int) (int64, error) // TODO: project archive (soft delete)
}

//go:generate mockgen -destination=mocks/mock_projects_service.go -package=mocks . ProjectsService
type ProjectsService interface {
	SearchProjects(p SearchProjectsInput) (*[]Project, error)
	CreateProject(project Project) (int, error)
	ReadProject(id int) (*Project, error)
	UpdateProject(id int, project Project) (int64, error)
	DeleteProject(id int) (int64, error)
}
