// Package main runs browser-ready htmx examples.
package main

import (
	"context"
	"crypto/rand"
	"embed"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"

	basicexample "example.com/htmx-go-example/internal/basic"
	formexample "example.com/htmx-go-example/internal/form"
)

const (
	defaultAddress  = "127.0.0.1:8080"
	readTimeout     = 10 * time.Second
	readHeadTimeout = 5 * time.Second
	writeTimeout    = 10 * time.Second
	idleTimeout     = 60 * time.Second
)

//go:embed static/*
var staticFiles embed.FS

type csrfProtection struct {
	mutex  sync.Mutex
	tokens map[string]struct{}
}

var _ formexample.CSRFProtector = (*csrfProtection)(nil)

func newCSRFProtection() *csrfProtection {
	return &csrfProtection{
		mutex:  sync.Mutex{},
		tokens: make(map[string]struct{}),
	}
}

func (protection *csrfProtection) Token(_ *http.Request) string {
	token := rand.Text()

	protection.mutex.Lock()
	defer protection.mutex.Unlock()
	protection.tokens[token] = struct{}{}
	return token
}

func (protection *csrfProtection) Valid(request *http.Request) bool {
	provided := request.FormValue("_csrf")

	protection.mutex.Lock()
	defer protection.mutex.Unlock()
	_, valid := protection.tokens[provided]
	delete(protection.tokens, provided)
	return valid
}

type contactCreator struct{}

var _ formexample.ContactCreator = contactCreator{}

func (contactCreator) CreateContact(ctx context.Context, _ string) error {
	slog.InfoContext(ctx, "contact created")
	return nil
}

func main() {
	address := flag.String("addr", defaultAddress, "HTTP listen address")
	flag.Parse()

	app, err := newApp()
	if err != nil {
		slog.Error("build application", "error", err)
		os.Exit(1)
	}

	server := new(http.Server)
	server.Addr = *address
	server.Handler = app
	server.ReadTimeout = readTimeout
	server.ReadHeaderTimeout = readHeadTimeout
	server.WriteTimeout = writeTimeout
	server.IdleTimeout = idleTimeout

	slog.Info("htmx example server started", "url", "http://"+*address)
	if serveErr := server.ListenAndServe(); !errors.Is(serveErr, http.ErrServerClosed) {
		slog.Error("serve HTTP", "error", serveErr)
		os.Exit(1)
	}
}

func newApp() (http.Handler, error) {
	users, err := basicexample.NewHandler()
	if err != nil {
		return nil, fmt.Errorf("build users example: %w", err)
	}

	csrf := newCSRFProtection()
	contacts, err := formexample.NewHandler(csrf, contactCreator{})
	if err != nil {
		return nil, fmt.Errorf("build contacts example: %w", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /favicon.ico", func(writer http.ResponseWriter, _ *http.Request) {
		writer.WriteHeader(http.StatusNoContent)
	})
	mux.Handle("GET /static/", http.FileServer(http.FS(staticFiles)))
	mux.Handle("/users", users)
	mux.Handle("/contacts", contacts)
	mux.Handle("/contacts/", contacts)
	mux.HandleFunc("GET /{$}", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/users", http.StatusSeeOther)
	})
	return mux, nil
}
