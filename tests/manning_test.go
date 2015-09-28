package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/hoop33/tbdotd/app/models"
)

var manning = models.VendorWithName("Manning")

func TestManningParsesContents(t *testing.T) {
	filename := "manning.html"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		deal := manning.Manning(contents)
		assert.Equal(t, "hapi.js in Action", deal.Title)
	}
}

func TestManningEmptyReturnsNoResults(t *testing.T) {
	deal := manning.Manning([]byte{})
	assert.Equal(t, "No Results", deal.Title)
}
