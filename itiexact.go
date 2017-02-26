package itiexact

import (
	"net/http"
	"strings"
)

// ExactMatcher Store the path to match with the request Path
type ExactMatcher struct {
	path string
}

// New is the constructor to ItiExact
func New(path string) *ExactMatcher {
	p := strings.ToLower(strings.Trim(path, "/"))
	return &ExactMatcher{
		path: p,
	}
}

// Match returns if the request can be handled by this Route.
func (e *ExactMatcher) Match(req *http.Request) bool {
	path := strings.ToLower(strings.Trim(req.URL.Path, "/"))
	if path == e.path {
		return true
	}
	return false
}
