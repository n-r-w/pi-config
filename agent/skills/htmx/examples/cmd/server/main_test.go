package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"example.com/htmx-go-example/internal/htmltest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAppPages(t *testing.T) {
	t.Parallel()

	app, err := newApp()
	require.NoError(t, err)

	tests := []struct {
		name       string
		path       string
		elementTag string
		attributes map[string]string
	}{
		{
			name:       "users page",
			path:       "/users",
			elementTag: "ul",
			attributes: map[string]string{"id": "users-list"},
		},
		{
			name:       "contact form",
			path:       "/contacts/new",
			elementTag: "form",
			attributes: map[string]string{"id": "contact-form"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			request := httptest.NewRequestWithContext(t.Context(), http.MethodGet, test.path, nil)
			response := httptest.NewRecorder()

			app.ServeHTTP(response, request)

			assert.Equal(t, http.StatusOK, response.Code)
			assert.Equal(t, "text/html; charset=utf-8", response.Header().Get("Content-Type"))
			document := htmltest.Parse(t, response.Body.String())
			assert.True(t, document.HasDoctype())
			document.RequireElement(t, test.elementTag, test.attributes)
		})
	}
}

func TestAppStaticAssets(t *testing.T) {
	t.Parallel()

	app, err := newApp()
	require.NoError(t, err)

	tests := []struct {
		name        string
		path        string
		contentType string
	}{
		{
			name:        "local htmx",
			path:        "/static/htmx.min.js",
			contentType: "text/javascript; charset=utf-8",
		},
		{
			name:        "application styles",
			path:        "/static/styles.css",
			contentType: "text/css; charset=utf-8",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			request := httptest.NewRequestWithContext(t.Context(), http.MethodGet, test.path, nil)
			response := httptest.NewRecorder()

			app.ServeHTTP(response, request)

			assert.Equal(t, http.StatusOK, response.Code)
			assert.Equal(t, test.contentType, response.Header().Get("Content-Type"))
			assert.NotEmpty(t, response.Body.Bytes())
		})
	}
}

func TestAppFavicon(t *testing.T) {
	t.Parallel()

	app, err := newApp()
	require.NoError(t, err)
	request := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/favicon.ico", nil)
	response := httptest.NewRecorder()

	app.ServeHTTP(response, request)

	assert.Equal(t, http.StatusNoContent, response.Code)
}

func TestAppContactFormHTMXFlow(t *testing.T) {
	t.Parallel()

	app, err := newApp()
	require.NoError(t, err)

	formRequest := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/contacts/new", nil)
	formResponse := httptest.NewRecorder()
	app.ServeHTTP(formResponse, formRequest)

	result := formResponse.Result()
	t.Cleanup(func() {
		require.NoError(t, result.Body.Close())
	})
	require.Equal(t, http.StatusOK, result.StatusCode)
	document := htmltest.Parse(t, formResponse.Body.String())
	csrf := document.RequireElement(t, "input", map[string]string{"name": "_csrf"})
	token := htmltest.RequireAttribute(t, csrf, "value")
	require.NotEmpty(t, token)

	values := url.Values{
		"_csrf": {token},
		"name":  {"Ada Lovelace"},
	}
	createRequest := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/contacts",
		strings.NewReader(values.Encode()),
	)
	createRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	createRequest.Header.Set("HX-Request", "true")
	createResponse := httptest.NewRecorder()

	app.ServeHTTP(createResponse, createRequest)

	assert.Equal(t, http.StatusOK, createResponse.Code)
	assert.Equal(t, "/contacts", createResponse.Header().Get("HX-Location"))
}
