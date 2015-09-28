package models

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"time"

	"github.com/hoop33/tbdotd/Godeps/_workspace/src/github.com/revel/revel"
)

const VENDOR_FILE = "conf/vendors.json"

type Vendor struct {
	Name       string
	HomeUrl    string
	DealUrl    string
	DateFormat string
}

var Vendors []Vendor

var removeChars = []string{"'", " "}

func init() {
	// TODO there has to be a better way
	// This is to find the file whether we running or testing
	if !(loadVendors(VENDOR_FILE) || loadVendors(fmt.Sprintf("../%s", VENDOR_FILE))) {
		revel.ERROR.Fatal("Error loading vendors from %s", VENDOR_FILE)
	}
}

func loadVendors(path string) bool {
	contents, err := ioutil.ReadFile(path)
	if err == nil {
		json.Unmarshal(contents, &Vendors)
		revel.INFO.Printf("Loaded vendors from %s", VENDOR_FILE)
		revel.TRACE.Printf("Vendor contents:\n%s", string(contents))
		return true
	}
	return false
}

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
		ImageUrl: "/public/img/notfound.png",
		Url:      vendor.HomeUrl,
	}
}

func (vendor *Vendor) Apress(payload []byte) Deal {
	re := regexp.MustCompile("\\<h2 class=\"icon\"\\>Deal of the Day\\</h2\\>(?s:.+?)\\<img .+?src=\"(.+?)\"(?s:.+?)\\<a href=\"(.+?)\"\\>(.+?)\\</a\\>")
	matches := re.FindSubmatch(payload)
	if matches != nil {
		return Deal{
			Vendor:   vendor,
			Title:    string(matches[3]),
			ImageUrl: string(matches[1]),
			Url:      string(matches[2]),
		}
	}
	return vendor.NotFound()
}

func (vendor *Vendor) Springer(payload []byte) Deal {
	re := regexp.MustCompile("\\<h2 class=\"icon\"\\>Springer Daily Deal\\</h2\\>(?s:.+?)\\<img .+?src=\"(.+?)\"(?s:.+?)\\<a href=\"(.+?)\"\\>(.+?)\\</a\\>")
	matches := re.FindSubmatch(payload)
	if matches != nil {
		return Deal{
			Vendor:   vendor,
			Title:    string(matches[3]),
			ImageUrl: string(matches[1]),
			Url:      string(matches[2]),
		}
	}
	return vendor.NotFound()
}

func (vendor *Vendor) InformIT(payload []byte) Deal {
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
	xml.Unmarshal(payload, &rss)
	if rss.Channel.Item.Title != "" {
		item := rss.Channel.Item

		date, _ := time.Parse(vendor.DateFormat, item.Date)
		return Deal{
			Vendor:   vendor,
			Title:    item.Title,
			ImageUrl: fmt.Sprintf("%sShowCover.aspx?isbn=%s&type=f", vendor.HomeUrl, item.Isbn),
			Url:      item.Link,
			Date:     date,
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

func (vendor *Vendor) PacktPublishing(payload []byte) Deal {
	re := regexp.MustCompile("\\<div class=\"dotd-main-book-image(?s:.+?)\\<img src=\"(.+?)\"(?s:.+?)\\<div class=\"dotd-main-book-title\"\\>(?s:.+?)\\<h2\\>(.+?)\\</h2\\>")
	matches := re.FindSubmatch(payload)
	if matches != nil {
		return Deal{
			Vendor:   vendor,
			Title:    string(matches[2]),
			ImageUrl: fmt.Sprintf("http:%s", matches[1]),
			Url:      vendor.DealUrl,
		}
	}
	return vendor.NotFound()
}

func (vendor *Vendor) PacktPublishingVideo(payload []byte) Deal {
	re := regexp.MustCompile("\\<div class=\"dotd-main-book-image(?s:.+?)\\<img src=\"(.+?)\"(?s:.+?)\\<div class=\"dotw-heading\"\\>(?s:\\s+?)(.+?)(?s:\\s+?)\\</div\\>")
	matches := re.FindSubmatch(payload)
	if matches != nil {
		return Deal{
			Vendor:   vendor,
			Title:    string(matches[2]),
			ImageUrl: fmt.Sprintf("http:%s", matches[1]),
			Url:      vendor.DealUrl,
		}
	}
	return vendor.NotFound()
}

func VendorWithName(name string) Vendor {
	for _, vendor := range Vendors {
		if vendor.Name == name {
			return vendor
		}
	}
	return Vendor{}
}
