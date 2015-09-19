package models

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
)

type Vendor struct {
	Name    string
	HomeUrl string
	DealUrl string
}

var removeChars = []string{"'", " "}

func (vendor *Vendor) GetProcessingMethodName() string {
	name := vendor.Name
	for _, r := range removeChars {
		name = strings.Replace(name, r, "", -1)
	}
	return name
}

func (vendor *Vendor) NotFound() Deal {
	return Deal{
		Vendor:   vendor,
		Title:    "No Results",
		ImageUrl: "",
		Url:      vendor.HomeUrl,
	}
}

func (vendor *Vendor) Apress(payload []byte) Deal {
	rss := struct {
		Channel struct {
			Items []struct {
				Title string `xml:"title"`
				Link  string `xml:"link"`
				Sku   string `xml:"sku"`
			} `xml:"item"`
		} `xml:"channel"`
	}{}
	xml.Unmarshal(payload, &rss)
	if len(rss.Channel.Items) > 0 {
		item := rss.Channel.Items[0]
		return Deal{
			Vendor:   vendor,
			Title:    item.Title,
			ImageUrl: item.Sku,
			Url:      item.Link,
		}
	}
	return vendor.NotFound()
}

func (vendor *Vendor) Springer(payload []byte) Deal {
	return vendor.Apress(payload)
}

func (vendor *Vendor) InformIT(payload []byte) Deal {
	rss := struct {
		Channel struct {
			Item struct {
				Title string `xml:"title"`
				Link  string `xml:"link"`
				Isbn  string `xml:"guid"`
			} `xml:"item"`
		} `xml:"channel"`
	}{}
	xml.Unmarshal(payload, &rss)
	if rss.Channel.Item.Title != "" {
		item := rss.Channel.Item
		return Deal{
			Vendor:   vendor,
			Title:    item.Title,
			ImageUrl: fmt.Sprintf("%sShowCover.aspx?isbn=%s&type=f", vendor.HomeUrl, item.Isbn),
			Url:      item.Link,
		}
	}
	return vendor.NotFound()
}

func (vendor *Vendor) InformITVideo(payload []byte) Deal {
	return vendor.InformIT(payload)
}

func (vendor *Vendor) Peachpit(payload []byte) Deal {
	return vendor.InformIT(payload)
}

func (vendor *Vendor) PeachpitVideo(payload []byte) Deal {
	return vendor.InformIT(payload)
}

func (vendor *Vendor) OReilly(payload []byte) Deal {
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
			Vendor:   vendor,
			Title:    item.Title,
			ImageUrl: imageUrl,
			Url:      item.Link,
		}
	}
	return vendor.NotFound()
}

func (vendor *Vendor) OReillyBusiness(payload []byte) Deal {
	return vendor.OReilly(payload)
}

func (vendor *Vendor) OReillyVideo(payload []byte) Deal {
	return vendor.OReilly(payload)
}

func (vendor *Vendor) Manning(payload []byte) Deal {
	re := regexp.MustCompile("\\<div class=\"title\">Deal of the Day\\</div\\>(?s:.+?)\\<a href=\"(.+?)\"(?s:.+?)\\<div class=\"product-placeholder-title\"\\>\\s*(.+?)\\s*\\<(?s:.+?)\\<div style=\"background-image: url\\('(.+?)'\\)")
	matches := re.FindSubmatch(payload)
	if matches != nil {
		return Deal{
			Vendor:   vendor,
			Title:    string(matches[2]),
			ImageUrl: string(matches[3]),
			Url:      fmt.Sprintf("%s%s", vendor.HomeUrl, strings.TrimPrefix(string(matches[1]), "/")),
		}
	}
	return vendor.NotFound()
}
