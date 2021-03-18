package repository

import (
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

func (tr *pgProjectsRepository) SearchProjects() (*[]domain.Project, error) {
	query := "SELECT project_id, done, name, details from projects"
	project := &[]domain.Project{}
	if err := tr.conn.Select(project, query); err != nil {
		log.Errorf("query error: %v", err)
		return nil, err
	}
	return project, nil
}

func (tr *pgProjectsRepository) ReadProject(id int) (*domain.Project, error) {
	query := "SELECT project_id, done, name, details from projects WHERE project_id = $1"
	project := &domain.Project{}
	if err := tr.conn.Get(project, query, id); err != nil {
		log.Errorf("query error: %v", err)
		return nil, err
	}
	return project, nil
}

func (tr *pgProjectsRepository) CreateProject(project domain.Project) (int, error) {
	query := `INSERT INTO projects (name, details) VALUES ($1, $2) RETURNING project_id`
	var id int
	if err := tr.conn.QueryRow(query, project.Name, project.Details).Scan(&id); err != nil {
		log.Errorf("query error: %v", err)
		return -1, err
	}
	return id, nil
}

func (tr *pgProjectsRepository) UpdateProject(id int, project domain.Project) (int64, error) {
	query := `UPDATE projects SET name=COALESCE($2, name), details=COALESCE($3, details), done=COALESCE($4, done) WHERE project_id=$1`
	res, err := tr.conn.Exec(query, id, project.Name, project.Details, project.Done)

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
