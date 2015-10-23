package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/stretchr/testify/assert"
)

func TestSpringerParsesContents(t *testing.T) {
	filename := "apress.html"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		vendor := models.Vendors["Apress"]
		source := vendor.Sources["Springer"]
		deal := source.Springer(&vendor, contents)
		assert.Equal(t, "Concurrent Programming: Algorithms, Principles, and Foundations", deal.Title)
		assert.Equal(t, "http://www.apress.com/dailydealsspringer/index/view/id/1673/url/aHR0cDovL3d3dy5hcHJlc3MuY29tLzk3ODM2NDIzMjAyNjI=/", deal.Url)
	}
}

func TestSpringerEmptyReturnsNoResults(t *testing.T) {
	vendor := models.Vendors["Apress"]
	source := vendor.Sources["Springer"]
	deal := source.Springer(&vendor, []byte{})
	assert.Equal(t, "No Deal", deal.Title)
}
