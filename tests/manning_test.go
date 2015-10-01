package tests

import (
  "io/ioutil"
  "testing"

  "github.com/hoop33/tbdotd/app/models"
  "github.com/stretchr/testify/assert"
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
    assert.Equal(t, "https://manning.com/books/hapi-js-in-action", deal.Url)
  }
}

func TestManningEmptyReturnsNoResults(t *testing.T) {
  deal := manning.Manning([]byte{})
  assert.Equal(t, "No Results", deal.Title)
}

