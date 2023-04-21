package lorawan

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_RegisterDevEUI(t *testing.T) {
	t.Run("200 response registers deveuid", func(t *testing.T) {
		mux := http.NewServeMux()
		mockServer := httptest.NewServer(mux)

		mux.HandleFunc(registerDevEUIEndpoint, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "POST", r.Method)
			w.WriteHeader(200)
		})

		c := client{
			client:  http.DefaultClient,
			baseUrl: mockServer.URL,
		}

		success, err := c.RegisterDevEUI("ABCDEF")
		assert.Nil(t, err)
		assert.True(t, success)
	})

	t.Run("422 response fails to registers deveuid", func(t *testing.T) {
		mux := http.NewServeMux()
		mockServer := httptest.NewServer(mux)

		mux.HandleFunc(registerDevEUIEndpoint, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "POST", r.Method)
			w.WriteHeader(422)
		})

		c := client{
			client:  http.DefaultClient,
			baseUrl: mockServer.URL,
		}

		success, err := c.RegisterDevEUI("ABCDEF")
		assert.Nil(t, err)
		assert.False(t, success)
	})

	t.Run("non-200 or non-422 response returns an error", func(t *testing.T) {
		mux := http.NewServeMux()
		mockServer := httptest.NewServer(mux)

		mux.HandleFunc(registerDevEUIEndpoint, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "POST", r.Method)
			w.WriteHeader(418)
		})

		c := client{
			client:  http.DefaultClient,
			baseUrl: mockServer.URL,
		}

		success, err := c.RegisterDevEUI("ABCDEF")
		assert.NotNil(t, err)
		assert.False(t, success)
	})

	t.Run("empty deveuid causes error", func(t *testing.T) {
		c := client{}
		success, err := c.RegisterDevEUI("")

		assert.False(t, success)
		assert.NotNil(t, err)
	})
}
