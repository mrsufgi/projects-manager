package http

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/google/go-github/github"
	"github.com/julienschmidt/httprouter"
	"github.com/mrsufgi/projects-manager/internal/domain"
	"github.com/mrsufgi/projects-manager/pkg/transport"
	log "github.com/sirupsen/logrus"
)

type EventsHandler struct {
	TService domain.EventsService
}

func NewEventsHandler(r *httprouter.Router, ts domain.EventsService) *EventsHandler {
	handler := &EventsHandler{
		TService: ts,
	}

	// CRUD
	r.GET("/events", handler.searchEvents)
	r.GET("/events/:id", handler.readEvent)

	// WEBHOOK
	r.POST("/events/webhook/github", handler.githubWebhook)
	return handler
}

func (p *EventsHandler) searchEvents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	queryValues := r.URL.Query()
	rec, err := p.TService.SearchEvents(domain.SearchEventsInput{Name: queryValues.Get("name")})
	if err != nil {
		herr := &transport.ResponseError{HTTPStatus: http.StatusBadRequest, Code: 40001, Message: "unable to search events"}
		herr.WriteToResponse(w)
		return
	}

	if err := json.NewEncoder(w).Encode(rec); err != nil {
		log.Errorf("unable to encode response %v", err)
	}
}

func (p *EventsHandler) readEvent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Errorf("unable to convert the string into int.  %v", err)
	}

	rec, err := p.TService.ReadEvent(id)
	if err != nil {
		log.Errorf("unable to read event %v", err)
		herr := &transport.ResponseError{HTTPStatus: http.StatusBadRequest, Code: 40001, Message: "Unable to read event"}
		herr.WriteToResponse(w)
		return
	}

	if err := json.NewEncoder(w).Encode(rec); err != nil {
		log.Errorf("unable to encode response %v", err)
	}
}

// TODO: Error status returning to github (not sure if needed)
func (p *EventsHandler) githubWebhook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	payload, err := github.ValidatePayload(r, []byte("test")) // TODO: use secret from config
	if err != nil {
		log.Errorf("unable to validate request body: %s\n", err)
		return
	}
	defer r.Body.Close()

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Errorf("unable to parse webhook: %s\n", err)
		return
	}
	var id int = -1
	switch e := event.(type) {
	case *github.PushEvent:
		matched, err := regexp.MatchString("refs/tags", *e.Ref)
		if err != nil {
			log.Errorf("unable to match ref: %s", err)
			return
		}

		if matched {
			id, err = p.TService.LogEvent(domain.LogEventInput{RepoURL: e.GetRepo().GetHTMLURL(), CommitMessage: e.GetHeadCommit().GetMessage()})
			if err != nil {
				log.Errorf("unable to read event %v", err)
				herr := &transport.ResponseError{HTTPStatus: http.StatusBadRequest, Code: 40001, Message: "Unable to update event"}
				herr.WriteToResponse(w)
				return
			}
		}

	default:
		log.Printf("unsupported event type: %s", github.WebHookType(r))
	}

	// TODO: maybe redundant
	res := transport.ResponseMessage{
		ID:      int64(id),
		Message: "hook processed succesfully",
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Errorf("unable to encode response %v", err)
	}
}
