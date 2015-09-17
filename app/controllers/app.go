package controllers

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

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
		DealUrl: "http://www.apress.com/index.php/dailydeals/index/rss",
	},
	{
		Name:    "Springer",
		HomeUrl: "http://www.apress.com/",
		DealUrl: "http://www.apress.com/index.php/dailydealsspringer/index/rss",
	},
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

func (c App) Index() revel.Result {
	results := make(chan models.Deal)

	for _, vendor := range vendors {
		go func(vendor models.Vendor) {
			var deal models.Deal
			method := reflect.ValueOf(&vendor).MethodByName(strings.Replace(vendor.Name, "'", "", -1))
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
	// TODO sort!
	deals := []models.Deal{}
	for _, _ = range vendors {
		deals = append(deals, <-results)
	}

	return c.Render(deals)
}
