package service_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mrsufgi/projects-manager/internal/domain"
	"github.com/mrsufgi/projects-manager/internal/domain/mocks"
	"github.com/mrsufgi/projects-manager/internal/events/service"
)

func String(x string) *string {
	return &x
}

func TestNewEventService(t *testing.T) {
	type args struct {
		tr domain.EventsRepository
	}
	tests := []struct {
		name string
		args args
		want domain.EventsService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.NewEventService(tt.args.tr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventsService_SearchEvents(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tr := mocks.NewMockEventsRepository(ctrl)

	type fields struct {
		tr domain.EventsRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    *[]domain.Event
		wantErr bool
	}{
		{"happy search events", fields{tr: tr}, &[]domain.Event{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := service.NewEventService(
				tt.fields.tr,
			)
			tr.EXPECT().SearchEvents(nil).Return(&[]domain.Event{}, nil)

			got, err := ts.SearchEvents()
			if (err != nil) != tt.wantErr {
				t.Errorf("eventsService.SearchEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventsService.SearchEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventsService_LogEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tr := mocks.NewMockEventsRepository(ctrl)

	timestamp := &time.Time{}
	type fields struct {
		tr domain.EventsRepository
	}
	type args struct {
		event domain.Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"happy create event", fields{tr: tr}, args{domain.Event{ID: 0, Name: String("Test"), Timestamp: timestamp}}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := service.NewEventService(
				tt.fields.tr,
			)
			tr.EXPECT().CreateEvent(tt.args.event).Return(tt.args.event.ID, nil)

			got, err := ts.LogEvent(tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("eventsService.LogEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("eventsService.LogEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventsService_ReadEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tr := mocks.NewMockEventsRepository(ctrl)

	type fields struct {
		tr domain.EventsRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Event
		wantErr bool
	}{
		{"happy read event", fields{tr: tr}, args{id: 0},
			&domain.Event{ID: 0, Name: String("Test")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := service.NewEventService(
				tt.fields.tr,
			)

			// note: returning the 'tt.want', simplify the fake data and validation checks
			// the func doesn't alter the result.
			tr.EXPECT().ReadEvent(tt.args.id).Return(tt.want, nil)
			got, err := ts.ReadEvent(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("eventsService.ReadEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventsService.ReadEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
