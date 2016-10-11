// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import "time"

const IssueUrl = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
    TotalCount int 'json:"total_count"'
    Item []*Issue
}

type Issue struct {
    Number  int
    HTMLURL string 'json:"html_url"'
    Title   string
    State   string
    User    *User
    CreatedAt time.Time 'json:"created_at"'
    Body string // in Markdown format
}

type User struct {
    Login string
    HTMLRURL string 'json:"html_url"'
}
