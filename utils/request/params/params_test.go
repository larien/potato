package params

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want QueryParams
	}{
		{
			name: "should_return_QueryParams_with_default_values_because_no_query_params_are_provided",
			args: args{
				r: &http.Request{
					URL: &url.URL{
						RawQuery: "",
					},
				},
			},
			want: QueryParams{
				Search: Search{
					ItemsPerPage: defaultItemsPerPage,
					Page:         defaultPage,
				},
			},
		},
		{
			name: "should_return_QueryParams_with_default_values_because_of_invalid_query_params",
			args: args{
				r: &http.Request{
					URL: &url.URL{
						RawQuery: "itemsPerPage=abc&page=def",
					},
				},
			},
			want: QueryParams{
				Search: Search{
					ItemsPerPage: defaultItemsPerPage,
					Page:         defaultPage,
				},
			},
		},
		{
			name: "should_return_QueryParams_with_custom_itemsperpage",
			args: args{
				r: &http.Request{
					URL: &url.URL{
						RawQuery: "itemsPerPage=10",
					},
				},
			},
			want: QueryParams{
				Search: Search{
					ItemsPerPage: 10,
					Page:         defaultPage,
				},
			},
		},
		{
			name: "should_return_QueryParams_with_custom_page",
			args: args{
				r: &http.Request{
					URL: &url.URL{
						RawQuery: "page=10",
					},
				},
			},
			want: QueryParams{
				Search: Search{
					ItemsPerPage: defaultItemsPerPage,
					Page:         10,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
