package controllers

import (
	"reflect"

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
}

func (c App) Index() revel.Result {
	results := make(chan models.Deal)

	for _, vendor := range vendors {
		go func() {
			var deal models.Deal
			method := reflect.ValueOf(&vendor).MethodByName(vendor.Name)
			if method.IsValid() {
				values := method.Call([]reflect.Value{})
				deal = values[0].Interface().(models.Deal)
			} else {
				deal = models.Deal{
					Vendor:   &vendor,
					Title:    "Not Found",
					ImageUrl: "",
					Url:      "#",
				}
			}
			results <- deal
		}()
	}

	deals := []models.Deal{}
	for _, _ = range vendors {
		deals = append(deals, <-results)
	}

	return c.Render(deals)
}
