package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/stretchr/testify/assert"
)

func TestPeachpitParsesContents(t *testing.T) {
	filename := "peachpit.xml"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		vendor := models.Vendors["Peachpit"]
		source := vendor.Sources["Peachpit"]
		deal := source.Peachpit(&vendor, contents)
		assert.Equal(t, "Presentation Zen Design: Simple Design Principles and Techniques to Enhance Your Presentations by Garr Reynolds", deal.Title)
		assert.Equal(t, "http://www.peachpit.com/deals/", deal.Url)
		assert.Equal(t, int64(1442815200), deal.ExpirationDate.Unix())
	}
}

func TestPeachpitEmptyReturnsNoResults(t *testing.T) {
	vendor := models.Vendors["Peachpit"]
	source := vendor.Sources["Peachpit"]
	deal := source.Peachpit(&vendor, []byte{})
	assert.Equal(t, "No Deal", deal.Title)
}
