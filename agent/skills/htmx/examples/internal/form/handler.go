// Package formexample shows form handling for htmx and normal HTTP requests.
package formexample

import (
	"bytes"
	"errors"
	"html/template"
	"log/slog"
	"net/http"
	"strings"
)

var (
	errCSRFProtectorRequired  = errors.New("csrf protector is required")
	errContactCreatorRequired = errors.New("contact creator is required")
)

type handler struct {
	csrf      CSRFProtector
	creator   ContactCreator
	templates *template.Template
}

type contactView struct {
	CSRFToken string
	Name      string
	NameError string
}

// NewHandler parses templates and returns form example HTTP handler.
func NewHandler(csrf CSRFProtector, creator ContactCreator) (http.Handler, error) {
	if csrf == nil {
		return nil, errCSRFProtectorRequired
	}
	if creator == nil {
		return nil, errContactCreatorRequired
	}

	templates, err := parseTemplates()
	if err != nil {
		return nil, err
	}

	handler := &handler{
		csrf:      csrf,
		creator:   creator,
		templates: templates,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /contacts", handler.contacts)
	mux.HandleFunc("GET /contacts/new", handler.newContact)
	mux.HandleFunc("POST /contacts", handler.createContact)
	return mux, nil
}

func (h *handler) contacts(writer http.ResponseWriter, request *http.Request) {
	h.render(writer, request, http.StatusOK, "contacts-page", contactView{})
}

func (h *handler) newContact(writer http.ResponseWriter, request *http.Request) {
	view := contactView{
		CSRFToken: h.csrf.Token(request),
		Name:      "",
		NameError: "",
	}
	h.render(writer, request, http.StatusOK, "contact-page", view)
}

func (h *handler) createContact(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Vary", "HX-Request")

	if err := request.ParseForm(); err != nil {
		http.Error(writer, "invalid form", http.StatusBadRequest)
		return
	}
	if !h.csrf.Valid(request) {
		http.Error(writer, "forbidden", http.StatusForbidden)
		return
	}

	name := strings.TrimSpace(request.FormValue("name"))
	if name == "" {
		h.renderInvalidContact(writer, request)
		return
	}
	if err := h.creator.CreateContact(request.Context(), name); err != nil {
		slog.ErrorContext(request.Context(), "create contact", "error", err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}

	if isHTMX(request) {
		writer.Header().Set("HX-Location", "/contacts")
		writer.WriteHeader(http.StatusOK)
		return
	}
	http.Redirect(writer, request, "/contacts", http.StatusSeeOther)
}

func (h *handler) renderInvalidContact(writer http.ResponseWriter, request *http.Request) {
	view := contactView{
		CSRFToken: h.csrf.Token(request),
		Name:      request.FormValue("name"),
		NameError: "Name is required.",
	}
	if isHTMX(request) {
		h.render(writer, request, http.StatusOK, "contact-form", view)
		return
	}
	h.render(writer, request, http.StatusUnprocessableEntity, "contact-page", view)
}

func (h *handler) render(
	writer http.ResponseWriter,
	request *http.Request,
	status int,
	templateName string,
	view contactView,
) {
	var buffer bytes.Buffer
	if err := h.templates.ExecuteTemplate(&buffer, templateName, view); err != nil {
		slog.ErrorContext(request.Context(), "render contact", "error", err)
		http.Error(writer, "internal server error", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.WriteHeader(status)
	if _, err := buffer.WriteTo(writer); err != nil {
		slog.ErrorContext(request.Context(), "write contact", "error", err)
	}
}

func isHTMX(request *http.Request) bool {
	return strings.EqualFold(request.Header.Get("HX-Request"), "true")
}
