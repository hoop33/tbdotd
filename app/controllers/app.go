package controllers

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"time"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

var vendors = []models.Vendor{
	{
		Name:    "Apress",
		HomeUrl: "http://www.apress.com/",
		DealUrl: "http://www.apress.com/",
	},
	{
		Name:    "Springer",
		HomeUrl: "http://www.apress.com/",
		DealUrl: "http://www.apress.com/",
	},
	{
		Name:    "InformIT",
		HomeUrl: "http://www.informit.com/",
		DealUrl: "http://www.informit.com/deals/deal_rss.aspx",
	},
	{
		Name:    "InformIT Video",
		HomeUrl: "http://www.informit.com/",
		DealUrl: "http://www.informit.com/deals/video/deal_rss.aspx",
	},
	{
		Name:    "Peachpit",
		HomeUrl: "http://www.peachpit.com/",
		DealUrl: "http://www.peachpit.com/deals/deal_rss.aspx",
	},
	{
		Name:    "Peachpit Video",
		HomeUrl: "http://www.peachpit.com/",
		DealUrl: "http://www.peachpit.com/deals/video/deal_rss.aspx",
	},
	{
		Name:    "O'Reilly",
		HomeUrl: "http://www.oreilly.com/",
		DealUrl: "http://feeds.feedburner.com/oreilly/ebookdealoftheday",
	},
	{
		Name:    "O'Reilly Business",
		HomeUrl: "http://www.oreilly.com/",
		DealUrl: "http://feeds.feedburner.com/oreilly/mspebookdeal",
	},
	{
		Name:    "O'Reilly Video",
		HomeUrl: "http://www.oreilly.com/",
		DealUrl: "http://feeds.feedburner.com/oreilly/videodealoftheweek",
	},
	{
		Name:    "Manning",
		HomeUrl: "https://manning.com/",
		DealUrl: "https://manning.com/dotd",
	},
	{
		Name:    "Packt Publishing",
		HomeUrl: "https://www.packtpub.com/",
		DealUrl: "https://www.packtpub.com/books/deal-of-the-day",
	},
	{
		Name:    "Packt Publishing Video",
		HomeUrl: "https://www.packtpub.com/",
		DealUrl: "https://www.packtpub.com/videos/deal-of-the-week",
	},
}

func getUrl(url string) ([]byte, error) {
	revel.INFO.Printf("Retrieving %s", url)
	var contents []byte

	// TODO put timeout in configuration file
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}
	response, err := client.Get(url)
	if err == nil {
		contents, err = ioutil.ReadAll(response.Body)
		defer response.Body.Close()
	}
	if err == nil {
		revel.INFO.Printf("Retrieved %s", url)
		revel.TRACE.Printf("%s returned %s", url, contents)
	} else {
		revel.ERROR.Printf("%s got error '%s'", url, err.Error())
	}
	return contents, err
}

func (c App) Index() revel.Result {
	results := make(chan models.Deal)

	for _, vendor := range vendors {
		go func(vendor models.Vendor) {
			var deal models.Deal
			method := reflect.ValueOf(&vendor).MethodByName(vendor.GetProcessingMethodName())
			if method.IsValid() {
				payload, err := getUrl(vendor.DealUrl)
				if err == nil {
					values := method.Call([]reflect.Value{reflect.ValueOf(payload)})
					deal = values[0].Interface().(models.Deal)
				}
			} else {
				deal = vendor.NotFound()
			}
			results <- deal
		}(vendor)
	}

	// TODO should we use a WaitGroup here instead?
	deals := []models.Deal{}
	for _, _ = range vendors {
		deals = append(deals, <-results)
	}

	sort.Sort(models.ByVendorName(deals))
	return c.Render(deals)
}
