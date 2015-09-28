package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/stretchr/testify/assert"
)

var packtPublishing = models.VendorWithName("Packt Publishing")

func TestPacktPublishingParsesContents(t *testing.T) {
	filename := "packt_publishing.html"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		deal := packtPublishing.PacktPublishing(contents)
		assert.Equal(t, "Mastering JavaScript [Video]", deal.Title)
	}
}

func TestPacktPublishingEmptyReturnsNoResults(t *testing.T) {
	deal := packtPublishing.PacktPublishing([]byte{})
	assert.Equal(t, "No Results", deal.Title)
}
