package presentation

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetToken(t *testing.T) {
	t.Run("Tests not valid user", func(t *testing.T) {
		endpointWithQueryParams := fmt.Sprintf("/token?uid=%d&password=%s", 1, "pass")
		request, _ := http.NewRequest(http.MethodGet, endpointWithQueryParams, nil)
		response := httptest.NewRecorder()

		IDPHandler(response, request)

		got := response.Result().StatusCode
		want := 404
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
