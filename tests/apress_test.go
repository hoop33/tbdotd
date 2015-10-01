package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/stretchr/testify/assert"
)

var apress = models.VendorWithName("Apress")

func TestApressParsesContents(t *testing.T) {
	filename := "apress.html"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		deal := apress.Apress(contents)
		assert.Equal(t, "Android Best Practices", deal.Title)
		assert.Equal(t, "http://www.apress.com/dailydeals/index/view/id/1719/url/aHR0cDovL3d3dy5hcHJlc3MuY29tLzk3ODE0MzAyNTg1NzU=/", deal.Url)
	}
}

func TestApressEmptyReturnsNoResults(t *testing.T) {
	deal := apress.Apress([]byte{})
	assert.Equal(t, "No Results", deal.Title)
}
