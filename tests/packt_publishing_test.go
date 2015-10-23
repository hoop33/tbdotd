package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/stretchr/testify/assert"
)

func TestPacktPublishingParsesContents(t *testing.T) {
	filename := "packt_publishing.html"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		vendor := models.Vendors["Packt Publishing"]
		source := vendor.Sources["Packt Publishing"]
		deal := source.PacktPublishing(&vendor, contents)
		assert.Equal(t, "Mastering JavaScript", deal.Title)
		assert.Equal(t, "https://www.packtpub.com/books/deal-of-the-day", deal.Url)
	}
}

func TestPacktPublishingEmptyReturnsNoResults(t *testing.T) {
	vendor := models.Vendors["Packt Publishing"]
	source := vendor.Sources["Packt Publishing"]
	deal := source.PacktPublishing(&vendor, []byte{})
	assert.Equal(t, "No Deal", deal.Title)
}
