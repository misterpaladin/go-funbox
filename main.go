package funpics

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var providers = []func() string{
	provider0,
	provider1,
}

// Get get random picture url
func Get() (url string) {
	return providers[1]()
}

// GetFrom get from specific provider
func GetFrom(provider int) (url string) {
	if len(providers)-1 < provider {
		fmt.Println("No such provider")
		return ""
	}

	return providers[provider]()
}

func provider0() (url string) {
	var first = 1
	var last int

	pagesDoc, _ := goquery.NewDocument("http://vse-shutochki.ru/kartinki-prikolnye")
	pagesDoc.Find(".pagination ul li a").Each(func(i int, s *goquery.Selection) {
		if i > 0 {
			last, _ = strconv.Atoi(s.Text())
		}
	})

	page := random(first, last)
	picturesDoc, _ := goquery.NewDocument("http://vse-shutochki.ru/kartinki-prikolnye/" + strconv.Itoa(page))

	pictures := make([]string, 0)

	picturesDoc.Find(".post .hidden-phone").Each(func(i int, s *goquery.Selection) {
		page, _ := s.Attr("src")
		pictures = append(pictures, page)
	})

	url = pictures[random(0, len(pictures)-1)]

	return url
}

func provider1() (url string) {
	var first = 1
	var last int

	pagesDoc, _ := goquery.NewDocument("https://bugaga.ru/tags/прикольные+картинки/")
	pagesDoc.Find(".navigation a").Each(func(i int, s *goquery.Selection) {
		if i > 0 {
			last, _ = strconv.Atoi(s.Text())
		}
	})

	page := random(first, last)
	picturesDoc, _ := goquery.NewDocument("https://bugaga.ru/tags/прикольные+картинки/page/" + strconv.Itoa(page))

	pictures := make([]string, 0)

	picturesDoc.Find(".w_news .w_cntn a.highslide").Each(func(i int, s *goquery.Selection) {
		page, _ := s.Attr("href")
		pictures = append(pictures, page)
	})

	url = pictures[random(0, len(pictures)-1)]

	return url
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
