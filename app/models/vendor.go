package models

import "encoding/xml"

type Vendor struct {
	Name    string
	HomeUrl string
	DealUrl string
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
