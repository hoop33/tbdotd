package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/stretchr/testify/assert"
)

var peachpit = models.VendorWithName("Peachpit")

func TestPeachpitParsesContents(t *testing.T) {
	filename := "peachpit.xml"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		deal := peachpit.Peachpit(contents)
		assert.Equal(t, "eBook Deal of the Week ::\n\t\t\t\tPresentation Zen Design: Simple Design Principles and Techniques to Enhance Your Presentations by Garr Reynolds", deal.Title)
	}
}

func TestPeachpitEmptyReturnsNoResults(t *testing.T) {
	deal := peachpit.Peachpit([]byte{})
	assert.Equal(t, "No Results", deal.Title)
}
