package domain

import "time"

type Event struct {
	ID        int        `db:"event_id" json:"id,omitempty"`
	Name      *string    `db:"name" json:"name"`
	ProjectID int        `db:"project_id" json:"project_id"`
	Timestamp *time.Time `db:"timestamp" json:"timestamp"`
}

type SearchEventsInput struct {
	Name string
}

type LogEventInput struct {
	RepoURL       string
	CommitMessage string
	// Timestamp *time.Time
}

//go:generate mockgen -destination=mocks/mock_events_repository.go -package=mocks . EventsRepository
type EventsRepository interface {
	SearchEvents(p SearchEventsInput) (*[]Event, error)
	CreateEvent(event Event) (int, error)
	ReadEvent(id int) (*Event, error)
}

//go:generate mockgen -destination=mocks/mock_events_service.go -package=mocks . EventsService
type EventsService interface {
	SearchEvents(p SearchEventsInput) (*[]Event, error)
	LogEvent(p LogEventInput) (int, error)
	ReadEvent(id int) (*Event, error)
}
