package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mrsufgi/projects-manager/internal/domain"

	log "github.com/sirupsen/logrus"
)

type pgEventsRepository struct {
	conn *sqlx.DB
}

// NewPgEventsRepository will create an object that represent the EventsRepository interface
func NewPgEventsRepository(conn *sqlx.DB) domain.EventsRepository {
	r := &pgEventsRepository{
		conn: conn,
	}

	return r
}

// TODO: Add Context (and use user_id in search)
func (tr *pgEventsRepository) SearchEvents(p domain.SearchEventsInput) (*[]domain.Event, error) {
	// TODO: Better querybuilding
	var where string
	if p.Name != "" {
		where = fmt.Sprintf("WHERE name='%s' ORDER BY timestamp DESC", p.Name)
	}
	query := fmt.Sprintf("SELECT event_id, name, timestamp from events %s", where)
	event := &[]domain.Event{}
	if err := tr.conn.Select(event, query); err != nil {
		log.Errorf("query error: %v", err)
		return nil, err
	}
	return event, nil
}

func (tr *pgEventsRepository) ReadEvent(id int) (*domain.Event, error) {
	query := "SELECT event_id, name, timestamp from events WHERE event_id = $1"
	event := &domain.Event{}
	if err := tr.conn.Get(event, query, id); err != nil {
		log.Errorf("query error: %v", err)
		return nil, err
	}
	return event, nil
}

// Note: for simplicity timestamp is automatically created by the DB and not using event_ts
func (tr *pgEventsRepository) CreateEvent(event domain.Event) (int, error) {
	query := `INSERT INTO events (name) VALUES ($1) RETURNING event_id`
	var id int
	if err := tr.conn.QueryRow(query, event.Name).Scan(&id); err != nil {
		log.Errorf("query error: %v", err)
		return -1, err
	}
	return id, nil
}
