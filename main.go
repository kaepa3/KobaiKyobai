package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/PuerkitoBio/goquery"
)

type Record struct {
	Name string
	date time.Time
}

type Config struct {
	LastDate time.Time
}

var config Config

func main() {
	url := `https://競売公売.com/?pid=14`
	re := regexp.MustCompile(`^(\d+)月(\d+)日`)
	readConfig()

	if doc, err := goquery.NewDocument(url); err == nil {
		selection := doc.Find(`div.content-list`)
		list := selection.Find(`article`)

		list.Each(func(index int, s *goquery.Selection) {
			info := s.Find(`div.info`)

			name := info.Find(`a`)
			fmt.Printf("name:%s\n", name.Text())
			attension := info.Find(`ul li.attention`)

			find := re.FindStringSubmatch(attension.Text())

			now := time.Now()
			month, _ := strconv.Atoi(find[1])
			day, _ := strconv.Atoi(find[2])

			time := time.Date(now.Year(), time.Month(month), (day), 0, 0, 0, 0, time.Local)
			fmt.Println(time.String())
		})
	}
}

func readConfig() {
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		fmt.Println(err)
	}
}
