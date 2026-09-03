package formexample

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"example.com/htmx-go-example/internal/htmltest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const testCSRFToken = "csrf-token"

func TestGetContactForm(t *testing.T) {
	t.Parallel()

	handler, csrf, _ := newTestHandler(t)
	csrf.EXPECT().Token(gomock.Any()).Return(testCSRFToken)
	request := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/contacts/new", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "text/html; charset=utf-8", response.Header().Get("Content-Type"))

	document := htmltest.Parse(t, response.Body.String())
	assert.True(t, document.HasDoctype())
	form := document.RequireElement(t, "form", map[string]string{"id": "contact-form"})
	assert.Equal(t, "/contacts", htmltest.RequireAttribute(t, form, "action"))
	assert.Equal(t, "post", htmltest.RequireAttribute(t, form, "method"))
	assert.Equal(t, "/contacts", htmltest.RequireAttribute(t, form, "hx-post"))
	assert.Equal(t, "find button", htmltest.RequireAttribute(t, form, "hx-disabled-elt"))

	csrfInput := document.RequireElement(t, "input", map[string]string{"name": "_csrf"})
	assert.Equal(t, testCSRFToken, htmltest.RequireAttribute(t, csrfInput, "value"))
}

func TestInvalidContactSubmission(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		hxRequest    bool
		wantStatus   int
		wantDocument bool
	}{
		{
			name:         "normal request",
			hxRequest:    false,
			wantStatus:   http.StatusUnprocessableEntity,
			wantDocument: true,
		},
		{
			name:         "htmx request",
			hxRequest:    true,
			wantStatus:   http.StatusOK,
			wantDocument: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			handler, csrf, _ := newTestHandler(t)
			gomock.InOrder(
				csrf.EXPECT().Valid(gomock.Any()).Return(true),
				csrf.EXPECT().Token(gomock.Any()).Return(testCSRFToken),
			)
			request := newFormRequest(t, url.Values{
				"_csrf": {testCSRFToken},
				"name":  {""},
			}, test.hxRequest)
			response := httptest.NewRecorder()

			handler.ServeHTTP(response, request)

			assert.Equal(t, test.wantStatus, response.Code)
			assert.Equal(t, "HX-Request", response.Header().Get("Vary"))

			document := htmltest.Parse(t, response.Body.String())
			assert.Equal(t, test.wantDocument, document.HasDoctype())
			field := document.RequireElement(t, "input", map[string]string{"id": "contact-name"})
			assert.Equal(t, "true", htmltest.RequireAttribute(t, field, "aria-invalid"))
			errorID := htmltest.RequireAttribute(t, field, "aria-describedby")
			document.RequireElement(t, "p", map[string]string{
				"id":   errorID,
				"role": "alert",
			})
		})
	}
}

func TestSuccessfulContactSubmission(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		hxRequest      bool
		wantStatus     int
		wantLocation   string
		wantHXLocation string
	}{
		{
			name:           "normal request",
			hxRequest:      false,
			wantStatus:     http.StatusSeeOther,
			wantLocation:   "/contacts",
			wantHXLocation: "",
		},
		{
			name:           "htmx request",
			hxRequest:      true,
			wantStatus:     http.StatusOK,
			wantLocation:   "",
			wantHXLocation: "/contacts",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			handler, csrf, creator := newTestHandler(t)
			csrf.EXPECT().Valid(gomock.Any()).Return(true)
			creator.EXPECT().CreateContact(gomock.Any(), "Ada Lovelace").Return(nil)
			request := newFormRequest(t, url.Values{
				"_csrf": {testCSRFToken},
				"name":  {"Ada Lovelace"},
			}, test.hxRequest)
			response := httptest.NewRecorder()

			handler.ServeHTTP(response, request)

			assert.Equal(t, test.wantStatus, response.Code)
			assert.Equal(t, test.wantLocation, response.Header().Get("Location"))
			assert.Equal(t, test.wantHXLocation, response.Header().Get("HX-Location"))

			destination := test.wantLocation
			if destination == "" {
				destination = test.wantHXLocation
			}
			destinationRequest := httptest.NewRequestWithContext(t.Context(), http.MethodGet, destination, nil)
			destinationResponse := httptest.NewRecorder()
			handler.ServeHTTP(destinationResponse, destinationRequest)
			assert.Equal(t, http.StatusOK, destinationResponse.Code)
			assert.True(t, htmltest.Parse(t, destinationResponse.Body.String()).HasDoctype())
		})
	}
}

func TestContactSubmissionRejectsInvalidCSRF(t *testing.T) {
	t.Parallel()

	for _, hxRequest := range []bool{false, true} {
		handler, csrf, _ := newTestHandler(t)
		csrf.EXPECT().Valid(gomock.Any()).Return(false)
		request := newFormRequest(t, url.Values{
			"_csrf": {"invalid-token"},
			"name":  {"Ada Lovelace"},
		}, hxRequest)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		assert.Equal(t, http.StatusForbidden, response.Code)
	}
}

func newTestHandler(t *testing.T) (http.Handler, *MockCSRFProtector, *MockContactCreator) {
	t.Helper()

	controller := gomock.NewController(t)
	csrf := NewMockCSRFProtector(controller)
	creator := NewMockContactCreator(controller)
	handler, err := NewHandler(csrf, creator)
	require.NoError(t, err)
	return handler, csrf, creator
}

func newFormRequest(t *testing.T, values url.Values, hxRequest bool) *http.Request {
	t.Helper()

	request := httptest.NewRequestWithContext(
		t.Context(),
		http.MethodPost,
		"/contacts",
		strings.NewReader(values.Encode()),
	)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if hxRequest {
		request.Header.Set("HX-Request", "true")
	}
	return request
}
