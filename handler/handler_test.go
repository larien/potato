package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPotatos(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody string
		potatos  map[string]Potato
	}{
		{
			name:     "GET /potatos",
			args:     args{},
			wantCode: 200,
			wantBody: `{"potato1":{"name":"potato1","added_at":"0001-01-01T00:00:00Z","last_modified_at":"0001-01-01T00:00:00Z"}}`,
			potatos: map[string]Potato{
				"potato1": {
					Name: "potato1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			potatos = tt.potatos
			GetPotatos(w, tt.args.r)
			if tt.wantCode != w.Code {
				t.Errorf("GetPotatos() = %v, want %v", w.Code, tt.wantCode)
			}
			if tt.wantBody != w.Body.String() {
				t.Errorf("GetPotatos() = %v, want %v", w.Body.String(), tt.wantBody)
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
		wantCode int
		wantBody string
		potatos  map[string]Potato
	}{
		{
			name: "Not found",
			args: args{
				r: httptest.NewRequest(http.MethodGet, "/potatos/potato1", nil),
			},
			wantCode: 404,
			wantBody: `{"error": "potato not found"}`,
			potatos:  nil,
		},
		{
			name: "GET /potatos/potato1",
			args: args{
				r: httptest.NewRequest(http.MethodGet, "/potatos/potato1", nil),
			},
			wantCode: 200,
			wantBody: `{"name":"potato1","added_at":"0001-01-01T00:00:00Z","last_modified_at":"0001-01-01T00:00:00Z"}`,
			potatos: map[string]Potato{
				"": {
					Name: "potato1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			potatos = tt.potatos
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
