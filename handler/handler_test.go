package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPotato(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody string
	}{
		{
			name: "GET /",
			args: args{
				r: httptest.NewRequest(http.MethodGet, "/", nil),
			},
			wantCode: 200,
			wantBody: `{"content": "potato"}`,
		},
		{
			name: "POST /",
			args: args{
				r: httptest.NewRequest(http.MethodPost, "/", nil),
			},
			wantCode: 405,
			wantBody: `{"error": "method not allowed"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			GetPotato(w, tt.args.r)
			if tt.wantCode != w.Code {
				t.Errorf("GetPotato() = %v, want %v", w.Code, tt.wantCode)
			}
			if tt.wantBody != w.Body.String() {
				t.Errorf("GetPotato() = %v, want %v", w.Body.String(), tt.wantBody)
			}
		})
	}
}
