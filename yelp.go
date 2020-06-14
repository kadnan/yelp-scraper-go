package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func getListing(listingURL string) []string {
	var links []string
	//HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	request, err := http.NewRequest("GET", listingURL, nil)
	if err != nil {
		fmt.Println(err)
	}

	//Setting headers
	request.Header.Set("pragma", "no-cache")
	request.Header.Set("cache-control", "no-cache")
	request.Header.Set("dnt", "1")
	request.Header.Set("upgrade-insecure-requests", "1")
	request.Header.Set("referer", "https://www.yelp.com/")
	resp, err := client.Do(request)
	if resp.StatusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(resp.Body)

		if err != nil {
			fmt.Println(err)
		}

		doc.Find(".lemon--ul__373c0__1_cxs a").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			link = "https://yelp.com/" + link

			// Make sure you we only fetch correct URL with corresponding title
			if strings.Contains(link, "biz/") {
				text := s.Text()
				if text != "" && text != "more" { //to avoid unecessary links
					links = append(links, link)
				}
			}

		})
	}

	return links
}

func main() {
	m := getListing("https://www.yelp.com/search?cflt=mobilephonerepair&find_loc=San+Francisco%2C+CA")
	fmt.Println(strings.Join(m, "\n"))
}
