package app

import (
	"fmt"
	"github.com/hoop33/tbdotd/Godeps/_workspace/src/github.com/revel/revel"
)

func init() {
	revel.OnAppStart(func() {
		fmt.Println("Go to /@tests to run the tests.")
	})
}
