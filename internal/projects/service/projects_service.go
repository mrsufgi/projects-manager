package service

import (
	"github.com/mrsufgi/projects-manager/internal/domain"
	"github.com/pusher/pusher-http-go"
	log "github.com/sirupsen/logrus"
)

type projectsService struct {
	tr domain.ProjectsRepository
	pc *pusher.Client
}

func NewProjectService(tr domain.ProjectsRepository, pc *pusher.Client) domain.ProjectsService {
	return &projectsService{
		tr: tr,
		pc: pc,
	}
}

// TODO: add search variables
func (ts *projectsService) SearchProjects() (*[]domain.Project, error) {
	res, err := ts.tr.SearchProjects()
	if len(*res) == 0 {
		log.Info("no projects found")
	}
	return res, err
}

func (ts *projectsService) CreateProject(project domain.Project) (int, error) {
	res, err := ts.tr.CreateProject(project)
	if res == -1 {
		log.Infof("unable to create project: %v", project)
		return -1, err
	}
	return res, nil
}

func (ts *projectsService) ReadProject(id int) (*domain.Project, error) {
	res, err := ts.tr.ReadProject(id)
	if res == nil {
		log.Infof("unable to find project: %v", id)
		return nil, err
	}
	data := map[string]string{"message": "hello world"}
	ts.pc.Trigger("my-channel", "my-event", data)
	return res, nil
}

func (ts *projectsService) UpdateProject(id int, project domain.Project) (int64, error) {
	res, err := ts.tr.UpdateProject(id, project)
	if res != 1 || err != nil {
		log.Infof("unable to update project: %v, %v", id, project)
		return 0, err
	}
	return res, err
}

func (ts *projectsService) DeleteProject(id int) (int64, error) {
	res, err := ts.tr.DeleteProject(id)
	if res != 1 || err != nil {
		log.Infof("unable to delete project: %v", id)
		return 0, err
	}
	return res, err
}
