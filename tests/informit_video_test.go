package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/stretchr/testify/assert"
)

func TestInformITVideoParsesContents(t *testing.T) {
	filename := "informit_video.xml"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		vendor := models.Vendors["InformIT"]
		source := vendor.Sources["InformIT Video"]
		deal := source.InformITVideo(&vendor, contents)
		assert.Equal(t, "Ruby on Rails Tutorial LiveLessons, The: Learn Web Development With Rails by Michael Hartl", deal.Title)
		assert.Equal(t, "http://www.informit.com/deals/video/", deal.Url)
		assert.Equal(t, int64(1443333600), deal.ExpirationDate.Unix())
	}
}

func TestInformITVideoEmptyReturnsNoResults(t *testing.T) {
	vendor := models.Vendors["InformIT"]
	source := vendor.Sources["InformIT Video"]
	deal := source.InformITVideo(&vendor, []byte{})
	assert.Equal(t, "No Deal", deal.Title)
}
