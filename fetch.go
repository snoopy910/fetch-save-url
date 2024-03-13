package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func fetchURL(urlStr string) error {
	resp, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	fileName := strings.Replace(u.Hostname(), ".", "_", -1) + ".html"
	err = os.WriteFile(fileName, body, 0644)
	if err != nil {
		return err
	}

	return nil
}

func fetchMetadata(urlStr string) error {
	resp, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	numLinks := len(doc.Find("a").Nodes)
	numImages := len(doc.Find("img").Nodes)

	fmt.Printf("site: %s\nnum_links: %d\nimages: %d\nlast_fetch: %s\n", urlStr, numLinks, numImages, time.Now().Format(time.RFC1123))

	return nil
}

func main() {
	fetchFlag := flag.Bool("metadata", false, "Fetch metadata")
	flag.Parse()

	urls := flag.Args()
	if len(urls) == 0 {
		fmt.Println("No URLs provided.")
		return
	}

	for _, urlStr := range urls {
		if *fetchFlag {
			err := fetchMetadata(urlStr)
			if err != nil {
				fmt.Printf("Error fetching metadata for %s: %v\n", urlStr, err)
			}
		} else {
			err := fetchURL(urlStr)
			if err != nil {
				fmt.Printf("Error fetching %s: %v\n", urlStr, err)
			}
			if err != nil {
				fmt.Printf("Error downloading assets: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Web page and assets saved successfully!")
		}
	}
}
