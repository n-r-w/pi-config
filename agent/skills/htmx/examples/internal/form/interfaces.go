package formexample

import (
	"context"
	"net/http"
)

//go:generate go tool mockgen -source=interfaces.go -destination=interfaces_mock_test.go -package=formexample

// CSRFProtector supplies and validates CSRF tokens for form requests.
type CSRFProtector interface {
	Token(request *http.Request) string
	Valid(request *http.Request) bool
}

// ContactCreator creates a contact from validated form input.
type ContactCreator interface {
	CreateContact(ctx context.Context, name string) error
}
