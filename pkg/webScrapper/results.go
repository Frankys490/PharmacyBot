package webScrapper

import (
	"fmt"
	"github.com/gocolly/colly"
)

func SearchResults(URL string) map[string][]string {
	results := make(map[string][]string)
	c := colly.NewCollector()

	c.OnHTML("img + a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		results[e.Text] = append(results[e.Text], link)
	})

	err := c.Visit(URL)
	if err != nil {
		panic(err)
	}
	return results
}

func GetInstruction(URL string) []string {
	var instructions []string
	c := colly.NewCollector()

	c.OnHTML("", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	err := c.Visit(URL)
	if err != nil {
		panic(err)
	}
	return instructions
}

/*func GetDescription(URL string) {
	var descriptionFromURL []string
	var descriptionFinal []string
	c := colly.NewCollector()

	c.OnHTML("ul[class=structure-list] li", func(e *colly.HTMLElement) {
		descriptionFromURL = append(descriptionFromURL, e.Text)
	})

	err := c.Visit(URL)
	if err != nil {
		panic(err)
	}
	for _, elem := range descriptionFromURL {
		elem = strings.Trim(elem, " \n")

		descriptionFinal = append(descriptionFinal, elem)
	}
}
*/
