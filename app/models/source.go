package models

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Source struct {
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

func cleanTitle(title string) string {
	title = strings.TrimSpace(title)
	for _, re := range TITLE_REGEXES {
		title = re.ReplaceAllString(title, "")
	}
	return strings.TrimSpace(title)
}

func cleanXML(xml []byte) []byte {
	for re, repl := range XML_REGEXES {
		xml = re.ReplaceAll(xml, repl)
	}
	return xml
}

func nextDay(date time.Time) time.Time {
	return date.Add(24 * time.Hour)
}

func nextWeek(date time.Time) time.Time {
	return date.Add(7 * 24 * time.Hour)
}

func (source *Source) GetProcessingMethodName(sourceName string) string {
	name := sourceName
	for _, r := range []string{"'", " "} {
		name = strings.Replace(name, r, "", -1)
	}
	return name
}

func (source *Source) NotFound() Deal {
	return Deal{
		Title:          "No Deal",
		ImageUrl:       "/public/img/notfound.png",
		ExpirationDate: time.Now(),
		// TODO Url
	}
}

func (source *Source) Apress(vendor *Vendor, payload []byte) Deal {
	re := regexp.MustCompile(source.Regex)
	matches := re.FindSubmatch(payload)
	if matches != nil {
		return Deal{
			Title:    cleanTitle(string(matches[3])),
			ImageUrl: string(matches[1]),
			Url:      string(matches[2]),
		}
	} else {
		return source.NotFound()
	}
}

func (source *Source) Springer(vendor *Vendor, payload []byte) Deal {
	return source.Apress(vendor, payload)
}

func (source *Source) InformIT(vendor *Vendor, payload []byte) Deal {
	deal := source.informITCommon(vendor, payload)
	deal.ExpirationDate = nextDay(deal.ExpirationDate)
	return deal
}

func (source *Source) InformITVideo(vendor *Vendor, payload []byte) Deal {
	deal := source.informITCommon(vendor, payload)
	deal.ExpirationDate = nextWeek(deal.ExpirationDate)
	return deal
}

func (source *Source) informITCommon(vendor *Vendor, payload []byte) Deal {
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

		date, _ := time.Parse(source.DateFormat, item.Date)
		return Deal{
			Title:          cleanTitle(item.Title),
			ImageUrl:       fmt.Sprintf("%sShowCover.aspx?isbn=%s&type=f", vendor.Url, item.Isbn),
			Url:            strings.TrimSpace(item.Link),
			ExpirationDate: date,
		}
	} else {
		return source.NotFound()
	}
}

func (source *Source) Peachpit(vendor *Vendor, payload []byte) Deal {
	deal := source.informITCommon(vendor, payload)
	deal.ExpirationDate = nextWeek(deal.ExpirationDate)
	return deal
}

func (source *Source) PeachpitVideo(vendor *Vendor, payload []byte) Deal {
	deal := source.informITCommon(vendor, payload)
	deal.ExpirationDate = nextWeek(deal.ExpirationDate)
	return deal
}

func (source *Source) OReilly(vendor *Vendor, payload []byte) Deal {
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

		return Deal{
			Title:    cleanTitle(item.Title),
			ImageUrl: imageUrl,
			Url:      item.Link,
		}
	} else {
		return source.NotFound()
	}
}

func (source *Source) OReillyBusiness(vendor *Vendor, payload []byte) Deal {
	return source.OReilly(vendor, payload)
}

func (source *Source) OReillyVideo(vendor *Vendor, payload []byte) Deal {
	return source.OReilly(vendor, payload)
}

func (source *Source) Manning(vendor *Vendor, payload []byte) Deal {
	re := regexp.MustCompile(source.Regex)
	matches := re.FindSubmatch(payload)
	if matches != nil {
		return Deal{
			Title:    cleanTitle(string(matches[2])),
			ImageUrl: string(matches[3]),
			Url:      fmt.Sprintf("%s%s", vendor.Url, strings.TrimPrefix(string(matches[1]), "/")),
		}
	} else {
		return source.NotFound()
	}
}

func (source *Source) PacktPublishing(vendor *Vendor, payload []byte) Deal {
	re := regexp.MustCompile(source.Regex)
	matches := re.FindSubmatch(payload)
	if matches != nil {
		return Deal{
			Title:    cleanTitle(strings.TrimSpace(string(matches[2]))),
			ImageUrl: fmt.Sprintf("http:%s", matches[1]),
			Url:      source.PayloadUrl,
		}
	} else {
		return source.NotFound()
	}
}

func (source *Source) PacktPublishingVideo(vendor *Vendor, payload []byte) Deal {
	return source.PacktPublishing(vendor, payload)
}
