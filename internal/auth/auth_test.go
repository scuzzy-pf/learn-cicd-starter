package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	type test struct {
		header http.Header
		want   string
	}
	bads := map[string]test{
		"bearer": {
			header: map[string][]string{"Authorization": {"Bearer asdfkejasdfjejsdf"}},
			want:   "",
		},
		"auth empty": {
			header: map[string][]string{"Authorization": {""}},
			want:   "",
		},
		"auth none": {
			header: map[string][]string{},
			want:   "",
		},
	}
	goods := map[string]test{
		"correct": {
			header: map[string][]string{"Authorization": {"ApiKey asdfkejasdfjejsdf"}},
			want:   "asdfkejasdfjejsdf",
		},
	}

	for name, bad := range bads {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(bad.header)
			if err == nil {
				t.Fatalf("expected error but did not get one.")
			}
			if bad.want != got {
				t.Fatalf("wanted: %s, got: %s", bad.want, got)
			}
		})
	}
	for name, good := range goods {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(good.header)
			if err != nil {
				t.Fatal("error not expected")
			}
			if good.want != got {
				t.Fatalf("wanted: %s, got: %s", good.want, got)
			}
		})
	}
}
