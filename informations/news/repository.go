package news

import (
	"fmt"
	"nearby/models"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const GOOGLE_URL = "https://google.com/search"

func NewsRepositoryFactory() func(city string) ([]models.News, error) {
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	return func(city string) ([]models.News, error) {
		request := createRequest(city)

		response, err := client.Do(request)
		if err != nil {
			return []models.News{}, err
		}

		if response.StatusCode != http.StatusOK {
			return []models.News{}, fmt.Errorf("invalid status code for news, %d", response.StatusCode)
		}

		defer response.Body.Close()

		document, _ := goquery.NewDocumentFromReader(response.Body)
		elements := document.Find("div.MkXWrd")

		news := []models.News{}

		for index := range elements.Nodes {
			item := elements.Eq(index)

			news = append(news, parseElement(item))
		}

		return news, nil
	}
}

func createRequest(city string) *http.Request {
	url := fmt.Sprintf("%s?q=%s", GOOGLE_URL, city)
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	addHeadersFor(request)

	return request
}

func addHeadersFor(request *http.Request) {
	request.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Add("accept-language", "fr,fr-FR;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	request.Header.Add("cache-control", "no-cache")
	request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36")
}

func parseElement(item *goquery.Selection) models.News {
	link := strings.TrimSpace(item.Find("a").AttrOr("href", ""))
	title, _ := item.Find("div.tNxQIb").Html()

	return models.News{
		Link:  link,
		Title: title,
	}
}
