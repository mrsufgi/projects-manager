package service

import (
	"github.com/mrsufgi/projects-manager/internal/domain"
	log "github.com/sirupsen/logrus"
)

type eventsService struct {
	er domain.EventsRepository
}

func NewEventService(er domain.EventsRepository) domain.EventsService {
	return &eventsService{
		er: er,
	}
}

func (ts *eventsService) SearchEvents(p domain.SearchEventsInput) (*[]domain.Event, error) {
	res, err := ts.er.SearchEvents(p)
	if len(*res) == 0 {
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

func (ts *eventsService) AddEvent(event domain.Event) (int, error) {
	res, err := ts.er.AddEvent(event)
	if res == -1 {
		log.Infof("unable to create event: %v", event)
		return -1, err
	}
	return res, nil
}
