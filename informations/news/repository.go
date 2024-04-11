package news

import (
	"fmt"
	"nearby/models"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const BING_URL = "https://www.bing.com/news/search"

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
		elements := document.Find("div.news-card")
		news := newsFromElements(elements)

		return news, nil
	}
}

func createRequest(city string) *http.Request {
	url := fmt.Sprintf("%s?q=%s", BING_URL, city)
	request, _ := http.NewRequest(http.MethodGet, url, nil)

	return request
}

func newsFromElements(elements *goquery.Selection) []models.News {
	news := []models.News{}

	for index := range elements.Nodes {
		item := elements.Eq(index)

		news = append(news, parse(item))
	}

	return news
}

func parse(item *goquery.Selection) models.News {
	href := item.AttrOr("data-url", "")
	title := item.AttrOr("data-title", "")

	return models.News{
		Link:  href,
		Title: title,
	}
}
