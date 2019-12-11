package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/kaepa3/KobaiKyobai/config"
	"github.com/kaepa3/KobaiKyobai/record"
)

const (
	AnalyzeURL = `https://競売公売.com/auction/find?pid=14`
)

var conf config.Config
var dynamicConf config.DynamicConfig

func main() {
	conf, dynamicConf, _ = config.ReadAllConfig("config.toml", "dynamic.toml")

	records := CreateItemRecords()
	sendRecord := selectSendRecord(records)
	if len(sendRecord) != 0 {
		newRecord := config.DynamicConfig{BeforeNewestRecord: sendRecord[0]}
		config.WriteConfig("dynamic.toml", newRecord)
		nortify(sendRecord)
	}
}
func nortify(records []record.Record) {
	fmt.Println("output\n")
	fmt.Printf("%v", records)
}
func selectSendRecord(records []record.Record) []record.Record {
	result := make([]record.Record, 0, 10)
	for _, v := range records {
		if v.Name == dynamicConf.BeforeNewestRecord.Name {
			break
		}
		result = append(result, v)
	}
	return result
}

var re0 = regexp.MustCompile(`^(\d+)月(\d+)日`)
var re1 = regexp.MustCompile(`^(\d+)月(\d+)日$`)
var re2 = regexp.MustCompile(`^残り(\d+)日$`)

func createName(s *goquery.Selection) string {
	info := s.Find(`div.info`)
	name := info.Find(`a`)
	return name.Text()
}
func createTime(s *goquery.Selection) time.Time {
	attention := s.Find(`div.info ul li.attention`)
	reList := []*regexp.Regexp{re0, re1, re2}
	var endtime time.Time
	now := time.Now()
	for _, v := range reList {
		find := v.FindStringSubmatch(attention.Text())
		if 3 == len(find) {
			month, _ := strconv.Atoi(find[1])
			day, _ := strconv.Atoi(find[2])
			endtime = time.Date(now.Year(), time.Month(month), (day), 0, 0, 0, 0, time.Local)
		} else if len(find) == 2 {
			day, _ := strconv.Atoi(find[1])
			bufTime := now.AddDate(0, 0, -day)
			endtime = time.Date(now.Year(), now.Month(), bufTime.Day(), 0, 0, 0, 0, time.Local)
		}
	}
	return endtime
}
func CreateItemRecords() []record.Record {
	records := make([]record.Record, 0, 100)
	if doc, err := goquery.NewDocument(AnalyzeURL); err == nil {
		selection := doc.Find(`div.content-list`)
		list := selection.Find(`article`)
		list.Each(func(index int, s *goquery.Selection) {
			name := createName(s)
			endTime := createTime(s)
			rec := record.Record{name, endTime}
			records = append(records, rec)
		})
	}
	return records
}
