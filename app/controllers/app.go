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

var results = make(map[string]models.Result)

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
		for _, deal := range vendor.Deals {
			waitGroup.Add(1)
			go func(vendor models.Vendor, deal models.Deal) {
				defer waitGroup.Done()

				result := results[deal.Name]
				if time.Now().After(result.ExpirationDate) {

					// TODO determine whether current deal has expired
					method := reflect.ValueOf(&deal).MethodByName(deal.GetProcessingMethodName())
					if method.IsValid() {
						if payload, err := getUrl(deal.PayloadUrl, urlTimeout); err == nil {
							values := method.Call([]reflect.Value{reflect.ValueOf(&vendor), reflect.ValueOf(payload)})
							result := values[0].Interface().(models.Result)
							result.Vendor = &vendor
							result.Deal = &deal
							results[deal.Name] = result
						}
					}
				} else {
					revel.INFO.Printf("Pulling %s from cache", deal.Name)
				}
				// TODO Surely there's a better way to do this
				// We have two if statements, so two elses
				//if deal.Title == "" {
				//deal.NotFound()
				//}
			}(vendor, deal)
		}
	}

	waitGroup.Wait()
	//sort.Sort(models.ByVendorName(results))
	return c.Render(results)
}
