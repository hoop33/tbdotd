package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/stretchr/testify/assert"
)

func TestPacktPublishingVideoParsesContents(t *testing.T) {
	filename := "packt_publishing_video.html"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		vendor := models.Vendors["Packt Publishing"]
		source := vendor.Sources["Packt Publishing Video"]
		deal := source.PacktPublishingVideo(&vendor, contents)
		assert.Equal(t, "Learning Functional JavaScript", deal.Title)
		assert.Equal(t, "https://www.packtpub.com/videos/deal-of-the-week", deal.Url)
	}
}

func TestPacktPublishingVideoEmptyReturnsNoResults(t *testing.T) {
	vendor := models.Vendors["Packt Publishing"]
	source := vendor.Sources["Packt Publishing Video"]
	deal := source.PacktPublishingVideo(&vendor, []byte{})
	assert.Equal(t, "No Deal", deal.Title)
}
