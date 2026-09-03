package htmxexample

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/htmx-go-example/internal/htmltest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUsersHandlerRepresentations(t *testing.T) {
	t.Parallel()

	handler, err := NewHandler()
	require.NoError(t, err)

	tests := []struct {
		name         string
		hxRequest    string
		wantDocument bool
	}{
		{
			name:         "full page",
			hxRequest:    "",
			wantDocument: true,
		},
		{
			name:         "htmx fragment",
			hxRequest:    "TRUE",
			wantDocument: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			request := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/users", nil)
			if test.hxRequest != "" {
				request.Header.Set("HX-Request", test.hxRequest)
			}
			response := httptest.NewRecorder()

			handler.ServeHTTP(response, request)

			result := response.Result()
			t.Cleanup(func() {
				require.NoError(t, result.Body.Close())
			})

			assert.Equal(t, http.StatusOK, result.StatusCode)
			assert.Equal(t, "text/html; charset=utf-8", result.Header.Get("Content-Type"))
			assert.Equal(t, "HX-Request", result.Header.Get("Vary"))

			document := htmltest.Parse(t, response.Body.String())
			assert.Equal(t, test.wantDocument, document.HasDoctype())
			document.RequireElement(t, "ul", map[string]string{"id": "users-list"})
		})
	}
}

func TestUsersHandlerEscapesQuery(t *testing.T) {
	t.Parallel()

	handler, err := NewHandler()
	require.NoError(t, err)

	const query = `<script>alert(1)</script>`
	request := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodGet,
		"/users?q=%3Cscript%3Ealert%281%29%3C%2Fscript%3E",
		nil,
	)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	document := htmltest.Parse(t, response.Body.String())
	search := document.RequireElement(t, "input", map[string]string{"id": "search"})
	assert.Equal(t, query, htmltest.RequireAttribute(t, search, "value"))
	scripts := document.Elements("script", nil)
	require.Len(t, scripts, 1)
	assert.Equal(t, "/static/htmx.min.js", htmltest.RequireAttribute(t, scripts[0], "src"))
	assert.Empty(t, htmltest.Text(scripts[0]))
}

func TestFilterUsers(t *testing.T) {
	t.Parallel()

	assert.Equal(t, []string{"Ada Lovelace"}, filterUsers("ada"))
	assert.Empty(t, filterUsers("unknown"))
}
