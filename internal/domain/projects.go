package domain

import "time"

type Project struct {
	ID        int        `db:"project_id" json:"id,omitempty"`
	Name      *string    `db:"name" json:"name"`
	Verical   *string    `db:"vertical" json:"vertical"`
	Event     *string    `db:"event" json:"event"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

//go:generate mockgen -destination=mocks/mock_projects_repository.go -package=mocks . ProjectsRepository
type ProjectsRepository interface {
	SearchProjects() (*[]Project, error)
	CreateProject(project Project) (int, error)
	ReadProject(id int) (*Project, error)
	UpdateProject(id int, project Project) (int64, error)
	DeleteProject(id int) (int64, error) // TODO: project archive (soft delete)
}

//go:generate mockgen -destination=mocks/mock_projects_service.go -package=mocks . ProjectsService
type ProjectsService interface {
	SearchProjects() (*[]Project, error)
	CreateProject(project Project) (int, error)
	ReadProject(id int) (*Project, error)
	UpdateProject(id int, project Project) (int64, error)
	DeleteProject(id int) (int64, error)
}
