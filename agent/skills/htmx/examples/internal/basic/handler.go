// Package htmxexample shows full-page and fragment responses for htmx.
package htmxexample

import (
	"bytes"
	"html/template"
	"log/slog"
	"net/http"
	"strings"
)

type handler struct {
	templates *template.Template
}

type usersView struct {
	Query string
	Users []string
}

// NewHandler parses templates and returns the example HTTP handler.
func NewHandler() (http.Handler, error) {
	templates, err := parseTemplates()
	if err != nil {
		return nil, err
	}

	handler := &handler{templates: templates}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", handler.users)
	return mux, nil
}

func (h *handler) users(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query().Get("q")
	view := usersView{
		Query: query,
		Users: filterUsers(query),
	}

	templateName := "users-page"
	if isHTMX(request) {
		templateName = "users-list"
	}
	body, err := executeTemplate(h.templates, templateName, view)
	if err != nil {
		slog.ErrorContext(request.Context(), "render users", "error", err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.Header().Add("Vary", "HX-Request")
	writer.WriteHeader(http.StatusOK)
	if _, writeErr := writer.Write(body); writeErr != nil {
		slog.ErrorContext(request.Context(), "write users", "error", writeErr)
	}
}

func isHTMX(request *http.Request) bool {
	return strings.EqualFold(request.Header.Get("HX-Request"), "true")
}

func executeTemplate(templates *template.Template, name string, data any) ([]byte, error) {
	var buffer bytes.Buffer
	if err := templates.ExecuteTemplate(&buffer, name, data); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func filterUsers(query string) []string {
	users := []string{"Ada Lovelace", "Grace Hopper", "Linus Torvalds"}
	if query == "" {
		return users
	}

	query = strings.ToLower(query)
	filtered := make([]string, 0, len(users))
	for _, user := range users {
		if strings.Contains(strings.ToLower(user), query) {
			filtered = append(filtered, user)
		}
	}
	return filtered
}
