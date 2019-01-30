package funpics

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var pictureProviders = []func() string{
	picProvider0,
	picProvider1,
}

var jokeProviders = []func() string{
	jokeProvider0,
}

// Picture get random picture url
func Picture() (url string) {
	var rand = 0
	if len(pictureProviders) > 1 {
		rand = random(0, len(pictureProviders)-1)
	}

	return pictureProviders[rand]()
}

// Joke get random joke text
func Joke() (text string) {
	var rand = 0
	if len(jokeProviders) > 1 {
		rand = random(0, len(jokeProviders)-1)
	}
	return jokeProviders[rand]()
}

// PictureFrom get picture from specific provider
func PictureFrom(provider int) (url string) {
	if len(pictureProviders)-1 < provider {
		fmt.Println("No such provider")
		return ""
	}

	return pictureProviders[provider]()
}

// JokeFrom get picture from specific provider
func JokeFrom(provider int) (url string) {
	if len(jokeProviders)-1 < provider {
		fmt.Println("No such provider")
		return ""
	}

	return jokeProviders[provider]()
}

func jokeProvider0() (text string) {
	texts := make([]string, 0)

	textDoc, _ := goquery.NewDocument("https://pda.anekdot.ru/random/anekdot/")
	textDoc.Find(".topicbox .text").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		fmt.Println(text)
		texts = append(texts, text)
	})

	text = texts[random(0, len(texts)-1)]

	return text
}

func picProvider0() (url string) {
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

func picProvider1() (url string) {
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
