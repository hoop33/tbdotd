package models

import (
	"encoding/xml"
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
				Guid  string `xml:guid"`
			} `xml:"item"`
		} `xml:"channel"`
	}{}
	xml.Unmarshal(payload, &rss)
	if rss.Channel.Item.Title != "" {
		item := rss.Channel.Item
		return Deal{
			Vendor:   vendor,
			Title:    item.Title,
			ImageUrl: item.Guid,
			Url:      item.Link,
		}
	}
	return vendor.NotFound()
}

func (vendor *Vendor) InformITVideo(payload []byte) Deal {
	return vendor.InformIT(payload)
}
