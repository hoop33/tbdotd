package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/hoop33/tbdotd/app/models"
)

var informITVideo = models.VendorWithName("InformIT Video")

func TestInformITVideoParsesContents(t *testing.T) {
	filename := "informit_video.xml"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		deal := informITVideo.InformITVideo(contents)
		assert.Equal(t, "Video Deal of the Week ::\n          Ruby on Rails Tutorial LiveLessons, The: Learn Web Development With Rails by Michael Hartl", deal.Title)
		assert.Equal(t, int64(1442728800), deal.Date.Unix())
	}
}

func TestInformITVideoEmptyReturnsNoResults(t *testing.T) {
	deal := informITVideo.InformITVideo([]byte{})
	assert.Equal(t, "No Results", deal.Title)
}
