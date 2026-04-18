package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		header  http.Header
		want    string
		wantErr bool
	}{
		{
			name: "Valid header",
			header: http.Header{
				"Authorization": []string{"ApiKey QWxhZGRpbjpvcGVu"},
			},
			want:    "QWxhZGRpbjpvcGVu",
			wantErr: false,
		},
		{
			name:    "Empty header",
			header:  http.Header{},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.header)

			if (err != nil) != tt.wantErr {
				t.Fatalf("error mismatch: got %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
