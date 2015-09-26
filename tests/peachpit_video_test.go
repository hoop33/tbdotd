package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/stretchr/testify/assert"
)

var peachpitVideo = models.VendorWithName("PeachpitVideo")

func TestPeachpitVideoParsesContents(t *testing.T) {
	filename := "peachpit_video.xml"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		deal := peachpitVideo.PeachpitVideo(contents)
		assert.Equal(t, "Video Deal of the Week ::\n          Workflow for Wedding Photographers: Learn by Video: Edit, design, and deliver everything from proofs to album layout in a single day by Khara Plicanic", deal.Title)
	}
}

func TestPeachpitVideoEmptyReturnsNoResults(t *testing.T) {
	deal := peachpitVideo.PeachpitVideo([]byte{})
	assert.Equal(t, "No Results", deal.Title)
}
