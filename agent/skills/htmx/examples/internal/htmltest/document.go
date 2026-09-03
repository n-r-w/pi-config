// Package htmltest provides semantic HTML assertions for example tests.
package htmltest

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

// Document contains parsed HTML response nodes.
type Document struct {
	root *html.Node
}

// Parse parses an HTML document or fragment.
func Parse(t testing.TB, markup string) *Document {
	t.Helper()

	root, err := html.Parse(strings.NewReader(markup))
	require.NoError(t, err)
	return &Document{root: root}
}

// HasDoctype reports whether response contains an HTML doctype.
func (document *Document) HasDoctype() bool {
	return hasNode(document.root, func(node *html.Node) bool {
		return node.Type == html.DoctypeNode && node.Data == "html"
	})
}

// Elements returns elements matching tag and attributes.
func (document *Document) Elements(tag string, attributes map[string]string) []*html.Node {
	elements := make([]*html.Node, 0)
	visit(document.root, func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == tag && hasAttributes(node, attributes) {
			elements = append(elements, node)
		}
	})
	return elements
}

// RequireElement returns exactly one matching element.
func (document *Document) RequireElement(
	t testing.TB,
	tag string,
	attributes map[string]string,
) *html.Node {
	t.Helper()

	elements := document.Elements(tag, attributes)
	require.Len(t, elements, 1)
	return elements[0]
}

// RequireAttribute returns required element attribute value.
func RequireAttribute(t testing.TB, node *html.Node, name string) string {
	t.Helper()

	value, exists := Attribute(node, name)
	require.True(t, exists, "attribute %q is required", name)
	return value
}

// Attribute returns element attribute value.
func Attribute(node *html.Node, name string) (string, bool) {
	for _, attribute := range node.Attr {
		if attribute.Key == name {
			return attribute.Val, true
		}
	}
	return "", false
}

// Text returns normalized text content from node and its descendants.
func Text(node *html.Node) string {
	var text strings.Builder
	visit(node, func(current *html.Node) {
		if current.Type == html.TextNode {
			text.WriteString(current.Data)
			text.WriteByte(' ')
		}
	})
	return strings.Join(strings.Fields(text.String()), " ")
}

func hasAttributes(node *html.Node, attributes map[string]string) bool {
	for name, expected := range attributes {
		actual, exists := Attribute(node, name)
		if !exists || actual != expected {
			return false
		}
	}
	return true
}

func hasNode(root *html.Node, match func(node *html.Node) bool) bool {
	matched := false
	visit(root, func(node *html.Node) {
		if match(node) {
			matched = true
		}
	})
	return matched
}

func visit(node *html.Node, visitNode func(node *html.Node)) {
	visitNode(node)
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		visit(child, visitNode)
	}
}
