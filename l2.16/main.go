package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

var visited = make(map[string]bool)

func main() {
	startURL := flag.String("url", "", "URL of the website to download")
	outputDir := flag.String("output", "site", "Directory to save website")
	depth := flag.Int("depth", 1, "Recursion depth")
	flag.Parse()

	if *startURL == "" {
		fmt.Println("Usage: go run main.go -url <URL> [-output <directory>] [-depth <n>]")
		os.Exit(1)
	}

	parsedURL, err := url.Parse(*startURL)
	if err != nil {
		fmt.Println("Invalid URL:", err)
		os.Exit(1)
	}

	err = os.MkdirAll(*outputDir, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create output directory:", err)
		os.Exit(1)
	}

	err = crawl(parsedURL, *outputDir, *depth)
	if err != nil {
		fmt.Println("Error during crawl:", err)
		os.Exit(1)
	}

	fmt.Println("Website downloaded successfully.")
}

func crawl(u *url.URL, outputDir string, depth int) error {
	if depth < 0 || visited[u.String()] {
		return nil
	}
	visited[u.String()] = true

	fmt.Println("Downloading:", u.String())

	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	localPath := filepath.Join(outputDir, localFileName(u))
	os.MkdirAll(filepath.Dir(localPath), os.ModePerm)
	err = os.WriteFile(localPath, body, 0644)
	if err != nil {
		return err
	}

	// Парсинг HTML и извлечение ссылок
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return err
	}

	links := extractLinks(doc, u)

	for _, link := range links {
		parsed, err := u.Parse(link)
		if err != nil {
			continue
		}
		if !sameDomain(u, parsed) {
			continue
		}
		err = crawl(parsed, outputDir, depth-1)
		if err != nil {
			fmt.Println("Failed to download:", parsed.String(), err)
		}
	}

	return nil
}

func localFileName(u *url.URL) string {
	path := u.Path
	if path == "" || strings.HasSuffix(path, "/") {
		path += "index.html"
	}
	return filepath.FromSlash(u.Host + path)
}

func sameDomain(base, other *url.URL) bool {
	return base.Hostname() == other.Hostname()
}

func extractLinks(n *html.Node, base *url.URL) []string {
	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "href" || attr.Key == "src" {
					links = append(links, attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return links
}
