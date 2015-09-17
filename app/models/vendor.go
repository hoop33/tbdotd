package models

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type Vendor struct {
	Name    string
	HomeUrl string
	DealUrl string
}

func getUrl(url string) ([]byte, error) {
	var contents []byte
	response, err := http.Get(url)
	if err == nil {
		contents, err = ioutil.ReadAll(response.Body)
		defer response.Body.Close()
	}
	return contents, err
}

func (vendor *Vendor) NotFound() Deal {
	return Deal{
		Vendor:   vendor,
		Title:    "No Results",
		ImageUrl: "",
		Url:      vendor.HomeUrl,
	}
}

func (vendor *Vendor) Apress() Deal {
	contents, err := getUrl(vendor.DealUrl)
	if err == nil {
		rss := struct {
			Channel struct {
				Items []struct {
					Title string `xml:"title"`
					Link  string `xml:"link"`
					Sku   string `xml:"sku"`
				} `xml:"item"`
			} `xml:"channel"`
		}{}
		xml.Unmarshal(contents, &rss)
		if len(rss.Channel.Items) > 0 {
			item := rss.Channel.Items[0]
			return Deal{
				Vendor:   vendor,
				Title:    item.Title,
				ImageUrl: item.Sku,
				Url:      item.Link,
			}
		}
	}
	return vendor.NotFound()
}

func (vendor *Vendor) Springer() Deal {
	return vendor.Apress()
}
