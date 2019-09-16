package main 

import(
	"net/url"
	"github.com/PuerkitoBio/goquery"
	"strings"
        "math/rand"
)

//TODO: Use regex matching for more consistent behavior
func VocaGet(producer string) (songLink string) {
	const BASE_URL = "https://vocaloid.fandom.com/wiki/Special:Search?query="
        priority := map[string]int {"youtube":5, "soundcloud":4, "nicovideo":3, "piapro":2, "bilibili":1}

        links := make([]string, 0)

	producer = url.QueryEscape(producer)
	page, err := goquery.NewDocument(BASE_URL + producer)
	HandleError(err, "Error constructing Goquery document from URL " + BASE_URL + producer)

        //result-link on the first page gets only the links from the search
	prodUrl, found := page.Find("a.result-link").First().Attr("href")
	if !found {
            return "Sorry, I couldn't find anything for " + producer + "!"
	}

	page, err = goquery.NewDocument(prodUrl)
	HandleError(err, "Error constructing Goquery document from URL" + prodUrl)

        //center gets all links to one song
	page.Find("center").Each(func(index int, item *goquery.Selection) {
                var highest int
                var bestLink string

		item.Find("a").Each(func(subindex int, subitem *goquery.Selection) {
			link, _ := subitem.Attr("href")
                        divided := strings.Split(link, ".")

                        if priority[divided[1]] > highest {
                            highest = priority[divided[1]]
                            bestLink = link
                        }
		})
                links = append(links, bestLink)
	})
        if len(links) == 0 {
            return "Sorry, I couldn't find anything for " + producer + "!"
        }
	return links[rand.Intn(len(links))]
}

func VocaRand() (songLink string) {
	const QUERY_URL = "https://vocaloid.fandom.com/wiki/Special:Random"

        links := make([]string, 0)

	for len(links) == 0 {
		page, err := goquery.NewDocument(QUERY_URL)
		HandleError(err, "Failed to construct document")

		page.Find("a").Each(func(index int, item *goquery.Selection) {
			link, _ := item.Attr("href")

			if strings.Contains(link, "youtube.com/watch") {
				links = append(links, link)
			}
		})
	}

	return songLink
}
