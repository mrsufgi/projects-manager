package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mrsufgi/projects-manager/internal/domain"
	"github.com/mrsufgi/projects-manager/pkg/transport"

	log "github.com/sirupsen/logrus"
)

type ProjectsHandler struct {
	TService domain.ProjectsService
}

func NewProjectsHandler(r *httprouter.Router, ts domain.ProjectsService) *ProjectsHandler {
	handler := &ProjectsHandler{
		TService: ts,
	}

	// CRUD
	r.GET("/projects", handler.searchProjects)
	r.GET("/projects/:id", handler.readProject)
	r.POST("/projects", handler.createProject)
	r.PUT("/projects/:id", handler.updateProject)
	r.DELETE("/projects/:id", handler.deleteProject)

	return handler
}

func (p *ProjectsHandler) searchProjects(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	queryValues := r.URL.Query()
	rec, err := p.TService.SearchProjects(domain.SearchProjectsInput{Name: queryValues.Get("name"), URL: queryValues.Get("url")})
	if err != nil {
		herr := &transport.ResponseError{HTTPStatus: http.StatusBadRequest, Code: 40001, Message: "unable to search projects"}
		herr.WriteToResponse(w)
		return
	}

	if err := json.NewEncoder(w).Encode(rec); err != nil {
		log.Errorf("unable to encode response %v", err)
	}
}

func (p *ProjectsHandler) readProject(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Errorf("unable to convert the string into int.  %v", err)
	}

	rec, err := p.TService.ReadProject(id)
	if err != nil {
		log.Errorf("unable to read project %v", err)
		herr := &transport.ResponseError{HTTPStatus: http.StatusBadRequest, Code: 40001, Message: "Unable to read project"}
		herr.WriteToResponse(w)
		return
	}

	if err := json.NewEncoder(w).Encode(rec); err != nil {
		log.Errorf("unable to encode response %v", err)
	}
}

func (p *ProjectsHandler) createProject(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var project domain.Project

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		log.Errorf("unable to parse the body.  %v", err)
	}

	id, err := p.TService.CreateProject(project)
	if err != nil {
		log.Errorf("unable to read project %v", err)
		herr := &transport.ResponseError{HTTPStatus: http.StatusBadRequest, Code: 40001, Message: "Unable to create project"}
		herr.WriteToResponse(w)
		return
	}

	res := transport.ResponseMessage{
		ID:      int64(id),
		Message: "Project created successfully",
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Errorf("unable to encode response %v", err)
	}
}

func (p *ProjectsHandler) updateProject(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Errorf("unable to convert the string into int.  %v", err)
	}

	var project domain.Project

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		log.Errorf("unable to parse the body.  %v", err)
	}

	affected, err := p.TService.UpdateProject(id, project)
	// TODO: Create HTTP Error pkg
	if err != nil {
		log.Errorf("unable to update project %v", err)
		herr := &transport.ResponseError{HTTPStatus: http.StatusBadRequest, Code: 40001, Message: "Unable to update project"}
		herr.WriteToResponse(w)
		return
	}

	// TODO: throw on not found in update query
	if affected == 0 {
		log.Errorf("unable to update project %v", err)
		herr := &transport.ResponseError{HTTPStatus: http.StatusBadRequest, Code: 40400, Message: "Unable to update project, nothing was changed."}
		herr.WriteToResponse(w)
		return
	}

	msg := fmt.Sprintf("project updated successfully. total rows affected %v", affected)

	res := transport.ResponseMessage{
		ID:      int64(id),
		Message: msg,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Errorf("unable to encode response %v", err)
	}
}

func (p *ProjectsHandler) deleteProject(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Errorf("unable to convert the string into int.  %v", err)
	}

	affected, err := p.TService.DeleteProject(id)
	if err != nil {
		log.Errorf("unable to delete project %v", err)
		herr := &transport.ResponseError{HTTPStatus: http.StatusBadRequest, Code: 40001, Message: "Unable to delete project"}
		herr.WriteToResponse(w)
		return
	}

	msg := fmt.Sprintf("project deleted successfully. total rows affected %v", affected)

	res := transport.ResponseMessage{
		ID:      int64(id),
		Message: msg,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Errorf("unable to encode response %v", err)
	}
}
