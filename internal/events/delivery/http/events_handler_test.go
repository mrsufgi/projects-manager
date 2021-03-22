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
	tdh "github.com/mrsufgi/projects-manager/internal/events/delivery/http"
)

func TestNewEventsHandler(t *testing.T) {
	type args struct {
		r  *httprouter.Router
		ts domain.EventsService
	}
	tests := []struct {
		name string
		args args
		want *tdh.EventsHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tdh.NewEventsHandler(tt.args.r, tt.args.ts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventsHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createReadRequest(t *testing.T, id int) *http.Request {
	r, err := http.NewRequest("GET", fmt.Sprintf("/events/%d", id), nil)
	if err != nil {
		t.Fatal(err)
	}
	return r
}
func TestEventsHandler_readEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	router := httprouter.New()
	s := mocks.NewMockEventsService(ctrl)
	w := httptest.NewRecorder()
	r := createReadRequest(t, 1)

	type fields struct {
		TService domain.EventsService
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
		{"happy read event call", fields{TService: s}, args{w: w, r: r}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tdh.NewEventsHandler(router, tt.fields.TService)
			s.EXPECT().ReadEvent(gomock.Any())
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
