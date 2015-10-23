package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/stretchr/testify/assert"
)

func TestOreillyParsesContents(t *testing.T) {
	filename := "oreilly.xml"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		vendor := models.Vendors["O'Reilly"]
		source := vendor.Sources["O'Reilly"]
		deal := source.OReilly(&vendor, contents)
		assert.Equal(t, "Social Sensing -\n$49.98 (Save 50%)\n\n\n Use code DEAL", deal.Title)
		assert.Equal(t, "http://shop.oreilly.com/product/9780128008676.do#2015-09-25feed", deal.Url)
	}
}

func TestOReillyEmptyReturnsNoResults(t *testing.T) {
	vendor := models.Vendors["O'Reilly"]
	source := vendor.Sources["O'Reilly"]
	deal := source.OReilly(&vendor, []byte{})
	assert.Equal(t, "No Deal", deal.Title)
}
