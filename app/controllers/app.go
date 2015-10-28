package controllers

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"sync"
	"time"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

var deals = make(map[string]models.Deal)

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

func (c App) Index() revel.Result {
	timeout := revel.Config.IntDefault("url.timeout", 5)
	revel.INFO.Printf("URL timeout set to %ds", timeout)

	urlTimeout := time.Duration(timeout) * time.Second

	var waitGroup sync.WaitGroup
	for _, vendor := range models.Vendors {
		for sourceName, source := range vendor.Sources {
			waitGroup.Add(1)
			go func(vendor models.Vendor, sourceName string, source models.Source) {
				defer waitGroup.Done()

				deal := deals[sourceName]
				if time.Now().After(deal.ExpirationDate) {
					method := reflect.ValueOf(&source).MethodByName(source.GetProcessingMethodName(sourceName))
					if method.IsValid() {
						if payload, err := getUrl(source.PayloadUrl, urlTimeout); err == nil {
							values := method.Call([]reflect.Value{reflect.ValueOf(&vendor), reflect.ValueOf(payload)})
							deal := values[0].Interface().(models.Deal)
							deal.Vendor = &vendor
							deal.Source = &source
							deal.Name = sourceName
							deals[sourceName] = deal
						}
					}
				} else {
					revel.INFO.Printf("Using %s from cache", sourceName)
				}
				if deal.Title == "" {
					deal = source.NotFound()
				}
			}(vendor, sourceName, source)
		}
	}

	waitGroup.Wait()
	return c.Render(deals)
}
