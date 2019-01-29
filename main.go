package funpics

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var first int
var last int
var updatedAt int64 = 0

// Get get random picture url
func Get() (url string) {
	if time.Now().Unix()-updatedAt > 600 {
		fmt.Println("Getting new pages")
		updatedAt = time.Now().Unix()
		first, last = getPicPages()
	}

	page := random(first, last)
	doc, _ := goquery.NewDocument("http://vse-shutochki.ru/kartinki-prikolnye/" + strconv.Itoa(page))

	pictures := make([]string, 0)

	doc.Find(".post .hidden-phone").Each(func(i int, s *goquery.Selection) {
		page, _ := s.Attr("src")
		pictures = append(pictures, page)
	})

	url = pictures[random(0, len(pictures)-1)]

	return url
}

func getPicPages() (first int, last int) {
	doc, _ := goquery.NewDocument("http://vse-shutochki.ru/kartinki-prikolnye")
	doc.Find(".pagination ul li a").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			first, _ = strconv.Atoi(s.Text())
		} else {
			last, _ = strconv.Atoi(s.Text())
		}
	})

	return first, last
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
