package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/hoop33/tbdotd/app/models"
)

var packtPublishingVideo = models.VendorWithName("Packt Publishing Video")

func TestPacktPublishingVideoParsesContents(t *testing.T) {
	filename := "packt_publishing_video.html"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		deal := packtPublishingVideo.PacktPublishingVideo(contents)
		assert.Equal(t, "\t\t\t\t\tLearning Functional JavaScript [Video]", deal.Title)
	}
}

func TestPacktPublishingVideoEmptyReturnsNoResults(t *testing.T) {
	deal := packtPublishingVideo.PacktPublishingVideo([]byte{})
	assert.Equal(t, "No Results", deal.Title)
}
