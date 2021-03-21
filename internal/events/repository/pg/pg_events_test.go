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
	"github.com/mrsufgi/events-manager/internal/domain"
	repository "github.com/mrsufgi/events-manager/internal/events/repository/pg"
	helpers "github.com/mrsufgi/events-manager/pkg/helpers"
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
func TestNewPgEventsRepository(t *testing.T) {
	type args struct {
		conn *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want domain.EventsRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repository.NewPgEventsRepository(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPgEventsRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgEventsRepository_ReadEvent(t *testing.T) {
	tr := repository.NewPgEventsRepository(getConn())

	type args struct {
		id int
	}
	tests := []struct {
		name    string
		tr      domain.EventsRepository
		args    args
		want    *domain.Event
		wantErr bool
	}{
		{"happy event spec", tr, args{id: 0}, &domain.Event{ID: 0, Done: false, Name: String("ori"), Details: String("mehhh")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.ReadEvent(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgEventRepository.ReadEvent() error = %#v, wantErr %#v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgEventRepository.ReadEvent() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_pgEventsRepository_SearchEvents(t *testing.T) {
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
		want    []domain.Event
		wantErr bool
	}{
		{"happy event spec", fields{conn: getConn()}, args{}, []domain.Event{{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := repository.NewPgEventsRepository(
				tt.fields.conn,
			)
			got, err := tr.SearchEvents()
			if (err != nil) != tt.wantErr {
				t.Errorf("pgEventsRepository.SearchEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgEventsRepository.SearchEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO: id is serial, running in parallel provide unexpected result.
// either create id externally (and mock it) or add a read by id and compare result.
func Test_pgEventsRepository_CreateEvent(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		event domain.Event
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"happy event spec", fields{conn: getConn()}, args{domain.Event{Name: String("test"), Details: String("test")}}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := repository.NewPgEventsRepository(
				tt.fields.conn,
			)
			got, err := tr.CreateEvent(tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgEventsRepository.CreateEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("pgEventsRepository.CreateEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgEventsRepository_UpdateEvent(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		id    int
		event domain.Event
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{"happy event spec", fields{conn: getConn()}, args{id: 1,
			event: domain.Event{Name: String("updated"), Details: String("updated"), Done: true}}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := repository.NewPgEventsRepository(
				tt.fields.conn,
			)
			got, err := tr.UpdateEvent(tt.args.id, tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgEventsRepository.UpdateEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgEventsRepository.UpdateEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgEventsRepository_DeleteEvent(t *testing.T) {
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
		{"happy event spec", fields{conn: getConn()}, args{id: 1}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := repository.NewPgEventsRepository(
				tt.fields.conn,
			)
			got, err := tr.DeleteEvent(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgEventsRepository.DeleteEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("pgEventsRepository.DeleteEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
