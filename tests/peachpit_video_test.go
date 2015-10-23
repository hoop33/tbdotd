package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/stretchr/testify/assert"
)

func TestPeachpitVideoParsesContents(t *testing.T) {
	filename := "peachpit_video.xml"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		vendor := models.Vendors["Peachpit"]
		source := vendor.Sources["Peachpit Video"]
		deal := source.PeachpitVideo(&vendor, contents)
		assert.Equal(t, "Workflow for Wedding Photographers: Learn by Video: Edit, design, and deliver everything from proofs to album layout in a single day by Khara Plicanic", deal.Title)
		assert.Equal(t, "http://www.peachpit.com/deals/video/", deal.Url)
		assert.Equal(t, int64(1442815200), deal.ExpirationDate.Unix())
	}
}

func TestPeachpitVideoEmptyReturnsNoResults(t *testing.T) {
	vendor := models.Vendors["Peachpit"]
	source := vendor.Sources["Peachpit Video"]
	deal := source.PeachpitVideo(&vendor, []byte{})
	assert.Equal(t, "No Deal", deal.Title)
}
