package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/larien/potato/service"
	"github.com/larien/potato/utils/request/params"
)

func TestGetPotatoes(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		potatoes service.Potatoes
		wantCode int
		wantBody string
	}{
		{
			name: "Failed_to_list",
			args: args{
				r: &http.Request{
					URL: &url.URL{},
				},
			},
			potatoes: mockPotatoes{
				fnList: func(params params.QueryParams) ([]service.Potato, error) {
					return nil, errors.New("failed to list")
				},
			},
			wantCode: 500,
			wantBody: `{"error": "could not list"}`,
		},
		{
			name: "GET /potatoes",
			args: args{
				r: &http.Request{
					URL: &url.URL{},
				},
			},
			potatoes: mockPotatoes{
				fnList: func(params params.QueryParams) ([]service.Potato, error) {
					return []service.Potato{
						{
							Name: "potato1",
						},
					}, nil
				},
			},
			wantCode: 200,
			wantBody: `[{"name":"potato1","added_at":"0001-01-01T00:00:00Z","last_modified_at":"0001-01-01T00:00:00Z"}]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceNew = func() service.Potatoes {
				return tt.potatoes
			}
			w := httptest.NewRecorder()
			GetPotatoes(w, tt.args.r)
			if tt.wantCode != w.Code {
				t.Errorf("GetPotatoes() = %v, want %v", w.Code, tt.wantCode)
			}
			if tt.wantBody != w.Body.String() {
				t.Errorf("GetPotatoes() = %v, want %v", w.Body.String(), tt.wantBody)
			}
		})
	}
}

func TestGetPotatoByID(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		potatoes service.Potatoes
		wantCode int
		wantBody string
	}{
		{
			name: "Not found",
			args: args{
				r: httptest.NewRequest(http.MethodGet, "/potatoes/potato1", nil),
			},
			potatoes: mockPotatoes{
				fnGet: func(id string) service.Potato {
					return service.Potato{}
				},
			},
			wantCode: 404,
			wantBody: `{"error": "potato not found"}`,
		},
		{
			name: "GET /potatoes/potato1",
			args: args{
				r: httptest.NewRequest(http.MethodGet, "/potatoes/potato1", nil),
			},
			potatoes: mockPotatoes{
				fnGet: func(id string) service.Potato {
					return service.Potato{
						Name: "potato1",
					}
				},
			},
			wantCode: 200,
			wantBody: `{"name":"potato1","added_at":"0001-01-01T00:00:00Z","last_modified_at":"0001-01-01T00:00:00Z"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceNew = func() service.Potatoes {
				return tt.potatoes
			}
			w := httptest.NewRecorder()
			GetPotatoByID(w, tt.args.r)
			if tt.wantCode != w.Code {
				t.Errorf("GetPotatoByID() = %v, want %v", w.Code, tt.wantCode)
			}
			if tt.wantBody != w.Body.String() {
				t.Errorf("GetPotatoByID() = %v, want %v", w.Body.String(), tt.wantBody)
			}
		})
	}
}

func TestCreatePotato(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		potatoes service.Potatoes
		wantCode int
		wantBody string
	}{
		{
			name: "Failed_to_decode",
			args: args{
				r: &http.Request{
					URL: &url.URL{},
					Body: ioutil.NopCloser(
						strings.NewReader(`test`)),
				},
			},
			wantCode: 400,
			wantBody: `{"error": "could not decode request body"}`,
		},
		{
			name: "Already_exists",
			args: args{
				r: &http.Request{
					URL: &url.URL{},
					Body: ioutil.NopCloser(
						strings.NewReader(`{"name": "potato1"}`)),
				},
			},
			potatoes: mockPotatoes{
				fnCreate: func(potato service.Potato) error {
					return service.ErrAlreadyExists
				},
			},
			wantCode: 400,
			wantBody: `{"error": "potato already exists"}`,
		},
		{
			name: "Failed_to_create",
			args: args{
				r: &http.Request{
					URL: &url.URL{},
					Body: ioutil.NopCloser(
						strings.NewReader(`{"name": "potato1"}`)),
				},
			},
			potatoes: mockPotatoes{
				fnCreate: func(potato service.Potato) error {
					return errors.New("failed to create")
				},
			},
			wantCode: 500,
			wantBody: `{"error": "failed to create"}`,
		},
		{
			name: "Success",
			args: args{
				r: &http.Request{
					URL: &url.URL{},
					Body: ioutil.NopCloser(
						strings.NewReader(`{"name": "potato1"}`)),
				},
			},
			potatoes: mockPotatoes{
				fnCreate: func(potato service.Potato) error {
					return nil
				},
			},
			wantCode: 201,
			wantBody: `{"name":"potato1"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceNew = func() service.Potatoes {
				return tt.potatoes
			}
			w := httptest.NewRecorder()
			CreatePotato(w, tt.args.r)
			if tt.wantCode != w.Code {
				t.Errorf("CreatePotato() = %v, want %v", w.Code, tt.wantCode)
			}
			if tt.wantBody != w.Body.String() {
				t.Errorf("CreatePotato() = %v, want %v", w.Body.String(), tt.wantBody)
			}
		})
	}
}

func TestUpdatePotato(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		potatoes service.Potatoes
		wantCode int
		wantBody string
	}{
		{
			name: "Failed_to_decode",
			args: args{
				r: &http.Request{
					URL: &url.URL{},
					Body: ioutil.NopCloser(
						strings.NewReader(`test`)),
				},
			},
			wantCode: 400,
			wantBody: `{"error": "could not decode request body"}`,
		},
		{
			name: "Not_found",
			args: args{
				r: &http.Request{
					URL: &url.URL{},
					Body: ioutil.NopCloser(
						strings.NewReader(`{"name": "potato1"}`)),
				},
			},
			potatoes: mockPotatoes{
				fnUpdate: func(potato service.Potato) error {
					return service.ErrNotFound
				},
			},
			wantCode: 404,
			wantBody: `{"error": "potato not found"}`,
		},
		{
			name: "Failed_to_update",
			args: args{
				r: &http.Request{
					URL: &url.URL{},
					Body: ioutil.NopCloser(
						strings.NewReader(`{"name": "potato1"}`)),
				},
			},
			potatoes: mockPotatoes{
				fnUpdate: func(potato service.Potato) error {
					return errors.New("failed to update")
				},
			},
			wantCode: 500,
			wantBody: `{"error": "failed to update"}`,
		},
		{
			name: "Success",
			args: args{
				r: &http.Request{
					URL: &url.URL{},
					Body: ioutil.NopCloser(
						strings.NewReader(`{"name": "potato1"}`)),
				},
			},
			potatoes: mockPotatoes{
				fnUpdate: func(potato service.Potato) error {
					return nil
				},
			},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceNew = func() service.Potatoes {
				return tt.potatoes
			}
			w := httptest.NewRecorder()
			UpdatePotato(w, tt.args.r)
			if tt.wantCode != w.Code {
				t.Errorf("UpdatePotato() = %v, want %v", w.Code, tt.wantCode)
			}
			if tt.wantBody != w.Body.String() {
				t.Errorf("UpdatePotato() = %v, want %v", w.Body.String(), tt.wantBody)
			}
		})
	}
}

func TestDeletePotato(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		potatoes service.Potatoes
		wantCode int
		wantBody string
	}{
		{
			name: "Not_found",
			args: args{
				r: &http.Request{
					URL: &url.URL{},
				},
			},
			potatoes: mockPotatoes{
				fnDelete: func(id string) error {
					return service.ErrNotFound
				},
			},
			wantCode: 404,
			wantBody: `{"error": "potato not found"}`,
		},
		{
			name: "Failed_to_delete",
			args: args{
				r: &http.Request{
					URL: &url.URL{},
				},
			},
			potatoes: mockPotatoes{
				fnDelete: func(id string) error {
					return errors.New("failed to delete")
				},
			},
			wantCode: 500,
			wantBody: `{"error": "failed to delete"}`,
		},
		{
			name: "Success",
			args: args{
				r: &http.Request{
					URL: &url.URL{},
				},
			},
			potatoes: mockPotatoes{
				fnDelete: func(id string) error {
					return nil
				},
			},
			wantCode: 204,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceNew = func() service.Potatoes {
				return tt.potatoes
			}
			w := httptest.NewRecorder()
			DeletePotato(w, tt.args.r)
			if tt.wantCode != w.Code {
				t.Errorf("DeletePotato() = %v, want %v", w.Code, tt.wantCode)
			}
			if tt.wantBody != w.Body.String() {
				t.Errorf("DeletePotato() = %v, want %v", w.Body.String(), tt.wantBody)
			}
		})
	}
}
