package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mrsufgi/projects-manager/internal/domain"

	log "github.com/sirupsen/logrus"
)

type pgProjectsRepository struct {
	conn *sqlx.DB
}

// NewPgProjectsRepository will create an object that represent the X interface
func NewPgProjectsRepository(conn *sqlx.DB) domain.ProjectsRepository {
	r := &pgProjectsRepository{
		conn: conn,
	}

	return r
}

// TODO: Add Context and Search Params
func (tr *pgProjectsRepository) SearchProjects(p domain.SearchProjectsInput) (*[]domain.Project, error) {
	// TODO: Better querybuilding
	var where string

	// TODO: Support by name
	if p.URL != "" {
		where = fmt.Sprintf("WHERE url='%s'", p.URL)
	}
	query := fmt.Sprintf("SELECT project_id, name, vertical, event, url, credentials, created_at, updated_at from projects %s", where)
	project := &[]domain.Project{}
	if err := tr.conn.Select(project, query); err != nil {
		log.Errorf("query error: %v", err)
		return nil, err
	}
	return project, nil
}

func (tr *pgProjectsRepository) ReadProject(id int) (*domain.Project, error) {
	query := "SELECT project_id, name, vertical, event, url, credentials, created_at, updated_at from projects WHERE project_id = $1"
	project := &domain.Project{}
	if err := tr.conn.Get(project, query, id); err != nil {
		log.Errorf("query error: %v", err)
		return nil, err
	}
	return project, nil
}

func (tr *pgProjectsRepository) CreateProject(project domain.Project) (int, error) {
	query := `INSERT INTO projects (name, vertical, event, url, credentials) VALUES ($1, $2, $3, $4, $5) RETURNING project_id`
	var id int
	if err := tr.conn.QueryRow(query, project.Name, project.Vertical, project.Event, project.URL, project.Credentials).Scan(&id); err != nil {
		log.Errorf("query error: %v", err)
		return -1, err
	}
	return id, nil
}

func (tr *pgProjectsRepository) UpdateProject(id int, project domain.Project) (int64, error) {
	query := `UPDATE projects 
		SET name=COALESCE($2, name), vertical=COALESCE($3, vertical), event=COALESCE($4, event), url=COALESCE($5, url) WHERE project_id=$1`
	res, err := tr.conn.Exec(query, id, project.Name, project.Vertical, project.Event, project.URL)

	if err != nil {
		log.Errorf("query error: %v", err)
		return -1, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("error checking affected rows: %v", err)
	}

	return rowsAffected, nil
}

func (tr *pgProjectsRepository) DeleteProject(id int) (int64, error) {
	query := `DELETE FROM projects WHERE project_id = $1`
	res, err := tr.conn.Exec(query, id)

	if err != nil {
		log.Errorf("query error: %v", err)
		return -1, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("error checking affected rows: %v", err)
	}

	return rowsAffected, nil
}
