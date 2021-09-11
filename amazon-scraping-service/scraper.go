package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type product struct {
	Name         string `json:"name"`
	ImageURL     string `json:"imageURL"`
	Description  string `json:"description"`
	Price        string `json:"price"`
	TotalRatings string `json:"totalRatings"`
}

type responsePayload struct {
	URL       string  `json:"url"`
	Timestamp string  `json:"timestamp"`
	Product   product `json:"product"`
}

func scrapeAmazonUrl(newUrl *urlStruct) (responsePayload, error) {

	var Response responsePayload
	Response.URL = newUrl.URL
	Response.Product.Name = "Not Found"
	Response.Product.ImageURL = "Not Found"
	Response.Product.Description = "Not Found"
	Response.Product.Price = "Not Found"
	Response.Product.TotalRatings = "Not Found"
	var newError error = nil

	//Verify newUrl is only from www.amazon.com
	match, _ := regexp.MatchString("www.amazon.com", newUrl.URL)

	if match == false {
		newError = errors.New("invalid Input : The url should only be for the domain 'www.amazon.com'")
	}

	c := colly.NewCollector(
		// Visit only these domains
		colly.AllowedDomains("www.amazon.com"),
	)

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		Response.Timestamp = time.Now().Local().String()
		fmt.Println("Visiting ", r.URL.String())
	})

	c.OnHTML(`span[id=productTitle]`, func(e *colly.HTMLElement) {
		Response.Product.Name = strings.Trim(e.Text, "\n")
		fmt.Println("Name ", Response.Product.Name)
	})

	c.OnHTML(`div[id=imgTagWrapperId]`, func(e *colly.HTMLElement) {
		Response.Product.ImageURL = e.ChildAttr("img", "src")

		fmt.Println("ImageUrl ", Response.Product.ImageURL)
	})

	c.OnHTML(`div[id=feature-bullets]`, func(e *colly.HTMLElement) {
		descriptionList := e.ChildTexts("ul>li:nth-child(n+4)>span")
		Response.Product.Description = ""
		for _, element := range descriptionList {
			Response.Product.Description = Response.Product.Description + "\n" + element
		}

		fmt.Println("Description ", Response.Product.Description)
	})

	c.OnHTML(`div[id="olp_feature_div"]`, func(e *colly.HTMLElement) {
		Response.Product.Price = e.ChildText("div:nth-child(4)>span>a>span:nth-child(2)")

		fmt.Println("Price ", Response.Product.Price)
	})

	c.OnHTML(`div[id="price"]`, func(e *colly.HTMLElement) {
		Response.Product.Price = e.ChildText("table>tbody>tr>td>span[id=priceblock_ourprice]")

		fmt.Println("Price ", Response.Product.Price)
	})

	c.OnHTML(`div[id=averageCustomerReviews]`, func(e *colly.HTMLElement) {
		Response.Product.TotalRatings = e.ChildText("span:nth-child(3)>a")

		fmt.Println("TotalRatings ", Response.Product.TotalRatings)
	})
	c.OnScraped(func(r *colly.Response) {
		persistenceResponse := callPersistence(&Response)
		fmt.Println("Response from persistence : ", persistenceResponse)
	})

	c.Visit(newUrl.URL)

	return Response, newError

}
