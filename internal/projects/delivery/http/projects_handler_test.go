package http_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/mrsufgi/projects-manager/internal/domain"
	"github.com/mrsufgi/projects-manager/internal/domain/mocks"
	tdh "github.com/mrsufgi/projects-manager/internal/projects/delivery/http"
)

func TestNewProjectsHandler(t *testing.T) {
	type args struct {
		r  *httprouter.Router
		ts domain.ProjectsService
	}
	tests := []struct {
		name string
		args args
		want *tdh.ProjectsHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tdh.NewProjectsHandler(tt.args.r, tt.args.ts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProjectsHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createReadRequest(t *testing.T, id int) *http.Request {
	r, err := http.NewRequest("GET", fmt.Sprintf("/projects/%d", id), nil)
	if err != nil {
		t.Fatal(err)
	}
	return r
}
func TestProjectsHandler_readProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	router := httprouter.New()
	s := mocks.NewMockProjectsService(ctrl)
	w := httptest.NewRecorder()
	r := createReadRequest(t, 1)

	type fields struct {
		TService domain.ProjectsService
	}
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		ps httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"happy read project call", fields{TService: s}, args{w: w, r: r}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tdh.NewProjectsHandler(router, tt.fields.TService)
			s.EXPECT().ReadProject(gomock.Any())
			router.ServeHTTP(tt.args.w, tt.args.r)
			resp := w.Result()
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				t.Errorf("Unexpected status code %d", resp.StatusCode)
			}
		})
	}
}

// TODO: add missing API tests
