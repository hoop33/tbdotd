package models

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Deal struct {
	Name       string
	DateFormat string
	PayloadUrl string
	Regex      string
}

var XML_REGEXES = map[*regexp.Regexp][]byte{
	regexp.MustCompile(" & "): []byte(" &amp; "),
}

var TITLE_REGEXES = []*regexp.Regexp{
	regexp.MustCompile("^.+?::"),
	regexp.MustCompile("^#.+?:"),
	regexp.MustCompile("\\[.+?\\]$"),
}

var removeChars = []string{"'", " "}

func cleanTitle(title string) string {
	title = strings.TrimSpace(title)
	for _, re := range TITLE_REGEXES {
		title = re.ReplaceAllString(title, "")
	}
	return title
}

func cleanXML(xml []byte) []byte {
	for re, repl := range XML_REGEXES {
		xml = re.ReplaceAll(xml, repl)
	}
	return xml
}

func (deal *Deal) GetProcessingMethodName() string {
	name := deal.Name
	for _, r := range removeChars {
		name = strings.Replace(name, r, "", -1)
	}
	return name
}

func (deal *Deal) NotFound() Result {
	return Result{
		Title:          "No Results",
		ImageUrl:       "/public/img/notfound.png",
		ExpirationDate: time.Now(),
		// TODO Url
	}
}

func (deal *Deal) Apress(vendor *Vendor, payload []byte) Result {
	re := regexp.MustCompile(deal.Regex)
	matches := re.FindSubmatch(payload)
	if matches != nil {
		return Result{
			Title:    cleanTitle(string(matches[3])),
			ImageUrl: string(matches[1]),
			Url:      string(matches[2]),
		}
	} else {
		return deal.NotFound()
	}
}

func (deal *Deal) Springer(vendor *Vendor, payload []byte) Result {
	return deal.Apress(vendor, payload)
}

func (deal *Deal) InformIT(vendor *Vendor, payload []byte) Result {
	rss := struct {
		Channel struct {
			Item struct {
				Title string `xml:"title"`
				Link  string `xml:"link"`
				Isbn  string `xml:"guid"`
				Date  string `xml:"pubDate"`
			} `xml:"item"`
		} `xml:"channel"`
	}{}
	xml.Unmarshal(cleanXML(payload), &rss)
	if rss.Channel.Item.Title != "" {
		item := rss.Channel.Item

		date, err := time.Parse(deal.DateFormat, item.Date)
		if err == nil {
			date = date.Add(24 * time.Hour)
		}
		return Result{
			Title:          cleanTitle(item.Title),
			ImageUrl:       fmt.Sprintf("%sShowCover.aspx?isbn=%s&type=f", vendor.Url, item.Isbn),
			Url:            strings.TrimSpace(item.Link),
			ExpirationDate: date,
		}
	} else {
		return deal.NotFound()
	}
}

func (deal *Deal) InformITVideo(vendor *Vendor, payload []byte) Result {
	return deal.InformIT(vendor, payload)
}

func (deal *Deal) Peachpit(vendor *Vendor, payload []byte) Result {
	return deal.InformIT(vendor, payload)
}

func (deal *Deal) PeachpitVideo(vendor *Vendor, payload []byte) Result {
	return deal.InformIT(vendor, payload)
}

func (deal *Deal) OReilly(vendor *Vendor, payload []byte) Result {
	rss := struct {
		Entries []struct {
			Title   string `xml:"title"`
			Link    string `xml:"id"`
			Content string `xml:"content"`
		} `xml:"entry"`
	}{}
	xml.Unmarshal(payload, &rss)
	if len(rss.Entries) > 0 {
		item := rss.Entries[0]

		imageUrl := ""
		re := regexp.MustCompile("img src=[\"'](.+?)[\"']")
		matches := re.FindStringSubmatch(item.Content)
		if matches != nil {
			imageUrl = matches[1]
		}

		return Result{
			Title:    cleanTitle(item.Title),
			ImageUrl: imageUrl,
			Url:      item.Link,
		}
	} else {
		return deal.NotFound()
	}
}

func (deal *Deal) OReillyBusiness(vendor *Vendor, payload []byte) Result {
	return deal.OReilly(vendor, payload)
}

func (deal *Deal) OReillyVideo(vendor *Vendor, payload []byte) Result {
	return deal.OReilly(vendor, payload)
}

func (deal *Deal) Manning(vendor *Vendor, payload []byte) Result {
	re := regexp.MustCompile(deal.Regex)
	matches := re.FindSubmatch(payload)
	if matches != nil {
		return Result{
			Title:    cleanTitle(string(matches[2])),
			ImageUrl: string(matches[3]),
			Url:      fmt.Sprintf("%s%s", vendor.Url, strings.TrimPrefix(string(matches[1]), "/")),
		}
	} else {
		return deal.NotFound()
	}
}

func (deal *Deal) PacktPublishing(vendor *Vendor, payload []byte) Result {
	re := regexp.MustCompile(deal.Regex)
	matches := re.FindSubmatch(payload)
	if matches != nil {
		return Result{
			Title:    cleanTitle(strings.TrimSpace(string(matches[2]))),
			ImageUrl: fmt.Sprintf("http:%s", matches[1]),
			Url:      deal.PayloadUrl,
		}
	} else {
		return deal.NotFound()
	}
}

func (deal *Deal) PacktPublishingVideo(vendor *Vendor, payload []byte) Result {
	return deal.PacktPublishing(vendor, payload)
}
