// +build live

package repository_test

// NOT a unit test suit, uses a real db. prevent it from running with unit.
// improve test cases to make it run with unit\live\clearing db (not seeding!)
// requires a more complex setup (migrations etc.)
// these tests mostly helped with initial implementation

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mrsufgi/projects-manager/internal/domain"
	repository "github.com/mrsufgi/projects-manager/internal/projects/repository/pg"
	helpers "github.com/mrsufgi/projects-manager/pkg/helpers"
	log "github.com/sirupsen/logrus"
)

func String(x string) *string {
	return &x
}

func getConn() *sqlx.DB {
	conn, err := sqlx.Connect("postgres", helpers.GetConnectionString())
	if err != nil {
		log.Fatalln(err)
	}
	return conn
}

// TODO: create function to create PG connection from env variables so it works with docker/local pg.
func TestNewPgProjectsRepository(t *testing.T) {
	type args struct {
		conn *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want domain.ProjectsRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repository.NewPgProjectsRepository(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPgProjectsRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgProjectsRepository_ReadProject(t *testing.T) {
	tr := repository.NewPgProjectsRepository(getConn())

	type args struct {
		id int
	}
	tests := []struct {
		name    string
		tr      domain.ProjectsRepository
		args    args
		want    *domain.Project
		wantErr bool
	}{
		{"happy project spec", tr, args{id: 0}, &domain.Project{ID: 0, Done: false, Name: String("ori"), Details: String("mehhh")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.ReadProject(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgProjectRepository.ReadProject() error = %#v, wantErr %#v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgProjectRepository.ReadProject() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_pgProjectsRepository_SearchProjects(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		id int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Project
		wantErr bool
	}{
		{"happy project spec", fields{conn: getConn()}, args{}, []domain.Project{{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := repository.NewPgProjectsRepository(
				tt.fields.conn,
			)
			got, err := tr.SearchProjects()
			if (err != nil) != tt.wantErr {
				t.Errorf("pgProjectsRepository.SearchProjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgProjectsRepository.SearchProjects() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO: id is serial, running in parallel provide unexpected result.
// either create id externally (and mock it) or add a read by id and compare result.
func Test_pgProjectsRepository_CreateProject(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		project domain.Project
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"happy project spec", fields{conn: getConn()}, args{domain.Project{Name: String("test"), Details: String("test")}}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := repository.NewPgProjectsRepository(
				tt.fields.conn,
			)
			got, err := tr.CreateProject(tt.args.project)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgProjectsRepository.CreateProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("pgProjectsRepository.CreateProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgProjectsRepository_UpdateProject(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		id      int
		project domain.Project
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{"happy project spec", fields{conn: getConn()}, args{id: 1,
			project: domain.Project{Name: String("updated"), Details: String("updated"), Done: true}}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := repository.NewPgProjectsRepository(
				tt.fields.conn,
			)
			got, err := tr.UpdateProject(tt.args.id, tt.args.project)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgProjectsRepository.UpdateProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgProjectsRepository.UpdateProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgProjectsRepository_DeleteProject(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		id int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{"happy project spec", fields{conn: getConn()}, args{id: 1}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := repository.NewPgProjectsRepository(
				tt.fields.conn,
			)
			got, err := tr.DeleteProject(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgProjectsRepository.DeleteProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("pgProjectsRepository.DeleteProject() = %v, want %v", got, tt.want)
			}
		})
	}
}
