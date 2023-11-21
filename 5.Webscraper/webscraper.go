package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	// Specify the URL of the website to scrape
	url := "https://www.w3schools.com/html/" // Replace with the actual URL

	// Make an HTTP request to the website
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer response.Body.Close()

	// Check if the request was successful (status code 200)
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code", response.Status)
		return
	}

	// Parse the HTML content of the response
	doc, err := html.Parse(response.Body)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	// Extract and display article titles
	articleTitles := extractArticleTitles(doc)
	fmt.Println("Article Titles:")
	for _, title := range articleTitles {
		fmt.Println("- " + title)
	}
}

func extractArticleTitles(n *html.Node) []string {
	var titles []string

	var visitNode func(*html.Node)
	visitNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "h2" {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.TextNode {
					titles = append(titles, c.Data)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visitNode(c)
		}
	}

	visitNode(n)
	return titles
}
