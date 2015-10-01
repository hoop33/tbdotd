package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/stretchr/testify/assert"
)

var packtPublishingVideo = models.VendorWithName("Packt Publishing Video")

func TestPacktPublishingVideoParsesContents(t *testing.T) {
	filename := "packt_publishing_video.html"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		deal := packtPublishingVideo.PacktPublishingVideo(contents)
		assert.Equal(t, "Learning Functional JavaScript [Video]", deal.Title)
    assert.Equal(t, "https://www.packtpub.com/videos/deal-of-the-week", deal.Url)
	}
}

func TestPacktPublishingVideoEmptyReturnsNoResults(t *testing.T) {
	deal := packtPublishingVideo.PacktPublishingVideo([]byte{})
	assert.Equal(t, "No Results", deal.Title)
}
