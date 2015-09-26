package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"time"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/revel/revel"
)

const VENDOR_FILE = "conf/vendors.json"

type App struct {
	*revel.Controller
}

func getUrl(url string, timeout time.Duration) ([]byte, error) {
	revel.INFO.Printf("Retrieving %s", url)
	var contents []byte

	client := http.Client{
		Timeout: timeout,
	}
	response, err := client.Get(url)
	if err == nil {
		contents, err = ioutil.ReadAll(response.Body)
		defer response.Body.Close()
	}
	if err == nil {
		revel.INFO.Printf("Retrieved %s", url)
		revel.TRACE.Printf("%s response: %s", url, contents)
	} else {
		revel.ERROR.Printf("%s error: %s", url, err.Error())
	}
	return contents, err
}

func LoadVendors() {
	contents, err := ioutil.ReadFile(VENDOR_FILE)
	if err != nil {
		revel.ERROR.Fatal("Error loading vendors from %s: %v", VENDOR_FILE, err)
	} else {
		json.Unmarshal(contents, &models.Vendors)
		revel.INFO.Printf("Loaded vendors from %s", VENDOR_FILE)
		revel.INFO.Printf("Vendor contents:\n%s", string(contents))
	}
}

func (c App) Index() revel.Result {
	results := make(chan models.Deal)

	timeout := revel.Config.IntDefault("url.timeout", 5)
	revel.INFO.Printf("URL timeout set to %ds", timeout)

	urlTimeout := time.Duration(timeout) * time.Second

	for _, vendor := range models.Vendors {
		go func(vendor models.Vendor) {
			var deal models.Deal
			method := reflect.ValueOf(&vendor).MethodByName(vendor.GetProcessingMethodName())
			if method.IsValid() {
				if payload, err := getUrl(vendor.DealUrl, urlTimeout); err == nil {
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
