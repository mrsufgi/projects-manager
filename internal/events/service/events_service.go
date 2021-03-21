package service

import (
	"fmt"
	"regexp"

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
	projects, err := ts.ps.SearchProjects(domain.SearchProjectsInput{URL: p.RepoURL}) // get project by git-url (and get project id)
	if err != nil || len(*projects) == 0 {
		log.Infof("unable to find projects with event: %v %v", err, p)
		return -1, err
	}

	project := (*projects)[0] // only one project per URL
	fmt.Println(p.CommitMessage, *project.Event)
	matched, err := regexp.MatchString(p.CommitMessage, *project.Event)
	if err != nil {
		log.Errorf("unable to match commit message: %s", err)
		return -1, err
	}

	if matched {
		e := domain.Event{Name: project.Event, ProjectID: project.ID}
		res, err := ts.er.CreateEvent(e)
		if res == -1 {
			log.Infof("unable to create event: %v", e)
			return -1, err
		}
		log.Debugf("event logged succesfully: %#v", e)
		return res, nil
	}

	log.Debugf("no event logged, project event not matched: %#v %#v", p, project)
	return -1, nil
}
