package service_test

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrsufgi/projects-manager/internal/domain"
	"github.com/mrsufgi/projects-manager/internal/domain/mocks"
	"github.com/mrsufgi/projects-manager/internal/projects/service"
)

func String(x string) *string {
	return &x
}

func TestNewProjectService(t *testing.T) {
	type args struct {
		tr domain.ProjectsRepository
	}
	tests := []struct {
		name string
		args args
		want domain.ProjectsService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.NewProjectService(tt.args.tr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProjectService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectsService_SearchProjects(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tr := mocks.NewMockProjectsRepository(ctrl)

	type fields struct {
		tr domain.ProjectsRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    *[]domain.Project
		wantErr bool
	}{
		{"happy search projects", fields{tr: tr}, &[]domain.Project{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := service.NewProjectService(
				tt.fields.tr,
			)
			tr.EXPECT().SearchProjects().Return(&[]domain.Project{}, nil)

			got, err := ts.SearchProjects()
			if (err != nil) != tt.wantErr {
				t.Errorf("projectsService.SearchProjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("projectsService.SearchProjects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectsService_CreateProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tr := mocks.NewMockProjectsRepository(ctrl)

	type fields struct {
		tr domain.ProjectsRepository
	}
	type args struct {
		project domain.Project
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"happy create project", fields{tr: tr}, args{domain.Project{ID: 0, Done: false, Name: String("Test"), Details: String("None")}}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := service.NewProjectService(
				tt.fields.tr,
			)
			tr.EXPECT().CreateProject(tt.args.project).Return(tt.args.project.ID, nil)

			got, err := ts.CreateProject(tt.args.project)
			if (err != nil) != tt.wantErr {
				t.Errorf("projectsService.CreateProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("projectsService.CreateProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectsService_ReadProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tr := mocks.NewMockProjectsRepository(ctrl)

	type fields struct {
		tr domain.ProjectsRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Project
		wantErr bool
	}{
		{"happy read project", fields{tr: tr}, args{id: 0},
			&domain.Project{ID: 0, Done: false, Name: String("Test"), Details: String("None")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := service.NewProjectService(
				tt.fields.tr,
			)

			// note: returning the 'tt.want', simplify the fake data and validation checks
			// the func doesn't alter the result.
			tr.EXPECT().ReadProject(tt.args.id).Return(tt.want, nil)
			got, err := ts.ReadProject(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("projectsService.ReadProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("projectsService.ReadProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectsService_UpdateProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tr := mocks.NewMockProjectsRepository(ctrl)

	type fields struct {
		tr domain.ProjectsRepository
	}
	type args struct {
		id      int
		project domain.Project
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{"happy update project", fields{tr: tr}, args{id: 0, project: domain.Project{Done: true}}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := service.NewProjectService(
				tt.fields.tr,
			)
			tr.EXPECT().UpdateProject(tt.args.id, tt.args.project).Return(int64(1), nil)
			got, err := ts.UpdateProject(tt.args.id, tt.args.project)
			if (err != nil) != tt.wantErr {
				t.Errorf("projectsService.UpdateProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("projectsService.UpdateProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_projectsService_DeleteProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tr := mocks.NewMockProjectsRepository(ctrl)

	type fields struct {
		tr domain.ProjectsRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{"happy delete project", fields{tr: tr}, args{id: 0}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := service.NewProjectService(
				tt.fields.tr,
			)
			tr.EXPECT().DeleteProject(tt.args.id).Return(int64(1), nil)
			got, err := ts.DeleteProject(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("projectsService.UpdateProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("projectsService.UpdateProject() = %v, want %v", got, tt.want)
			}
		})
	}
}
