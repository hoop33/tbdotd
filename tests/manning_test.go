package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/stretchr/testify/assert"
)

func TestManningParsesContents(t *testing.T) {
	filename := "manning.html"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		vendor := models.Vendors["Manning"]
		source := vendor.Sources["Manning"]
		deal := source.Manning(&vendor, contents)
		assert.Equal(t, "hapi.js in Action", deal.Title)
		assert.Equal(t, "https://manning.com/books/hapi-js-in-action", deal.Url)
	}
}

func TestManningEmptyReturnsNoResults(t *testing.T) {
	vendor := models.Vendors["Manning"]
	source := vendor.Sources["Manning"]
	deal := source.Manning(&vendor, []byte{})
	assert.Equal(t, "No Deal", deal.Title)
}
