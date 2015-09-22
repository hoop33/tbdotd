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

	for _, vendor := range models.Vendors {
		go func(vendor models.Vendor) {
			var deal models.Deal
			method := reflect.ValueOf(&vendor).MethodByName(vendor.GetProcessingMethodName())
			if method.IsValid() {
				if payload, err := getUrl(vendor.DealUrl); err == nil {
					values := method.Call([]reflect.Value{reflect.ValueOf(payload)})
					deal = values[0].Interface().(models.Deal)
				}
			}
			// TODO Surely there's a better way to do this
			// We have two if statements, so two elses
			if deal.Title == "" {
				deal = vendor.NotFound()
			}
			results <- deal
		}(vendor)
	}

	// TODO should we use a WaitGroup here instead?
	deals := []models.Deal{}
	for _, _ = range models.Vendors {
		deals = append(deals, <-results)
	}

	sort.Sort(models.ByVendorName(deals))
	return c.Render(deals)
}
