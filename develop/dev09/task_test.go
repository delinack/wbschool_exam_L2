package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestExtractLinks(t *testing.T) {
	htmlStr := `<html>
<head>
  <title>Test Page</title>
</head>
<body>
  <a href="https://example.com/page1">Link 1</a>
  <a href="https://example.com/page2">Link 2</a>
  <a href="https://example.com/page3">Link 3</a>
</body>
</html>`

	doc, _ := html.Parse(strings.NewReader(htmlStr))
	links := extractLinks(doc)

	if len(links) != 3 {
		t.Errorf("Expected 3 links, found %d", len(links))
	}

	expectedLinks := []string{
		"https://example.com/page1",
		"https://example.com/page2",
		"https://example.com/page3",
	}

	for i, link := range links {
		if link != expectedLinks[i] {
			t.Errorf("Expected link %s, found %s", expectedLinks[i], link)
		}
	}
}

func TestExtractLinksNoLinks(t *testing.T) {
	htmlStr := `<html>
<head>
  <title>Test Page</title>
</head>
<body>
  <p>No links here!</p>
</body>
</html>`

	doc, _ := html.Parse(strings.NewReader(htmlStr))
	links := extractLinks(doc)

	if len(links) != 0 {
		t.Errorf("Expected 0 links, found %d", len(links))
	}
}

func TestExtractLinksNestedLinks(t *testing.T) {
	htmlStr := `<html>
<head>
  <title>Test Page</title>
</head>
<body>
  <div>
    <p>
      <a href="https://example.com/page1">Link 1</a>
    </p>
  </div>
  <a href="https://example.com/page2">Link 2</a>
</body>
</html>`

	doc, _ := html.Parse(strings.NewReader(htmlStr))
	links := extractLinks(doc)

	if len(links) != 2 {
		t.Errorf("Expected 2 links, found %d", len(links))
	}

	expectedLinks := []string{
		"https://example.com/page1",
		"https://example.com/page2",
	}

	for i, link := range links {
		if link != expectedLinks[i] {
			t.Errorf("Expected link %s, found %s", expectedLinks[i], link)
		}
	}
}
