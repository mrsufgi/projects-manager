package service

import (
	"github.com/mrsufgi/projects-manager/internal/domain"
	log "github.com/sirupsen/logrus"
)

type eventsService struct {
	er domain.EventsRepository
	ps domain.ProjectsService
}

func NewEventService(er domain.EventsRepository, ps domain.ProjectsService) domain.EventsService {
	return &eventsService{
		er: er,
		ps: ps,
	}
}

func (ts *eventsService) SearchEvents(p domain.SearchEventsInput) (*[]domain.Event, error) {
	res, err := ts.er.SearchEvents(p)
	if res == nil || len(*res) == 0 {
		log.Info("no events found")
	}
	return res, err
}

func (ts *eventsService) ReadEvent(id int) (*domain.Event, error) {
	res, err := ts.er.ReadEvent(id)
	if res == nil {
		log.Infof("unable to find event: %v", id)
		return nil, err
	}

	return res, nil
}

// LogEvent find the project
func (ts *eventsService) LogEvent(p domain.LogEventInput) (int, error) {
	// p := ts.ps.SearchProjects() // get project by git-url (and get project id)
	// if project not found (no listeners, dont log)
	e := domain.Event{Name: &p.CommitMessage}
	res, err := ts.er.CreateEvent(e)
	if res == -1 {
		log.Infof("unable to create event: %v", e)
		return -1, err
	}
	log.Debugf("event logged succesfully: %#v", p)
	return res, nil
}
