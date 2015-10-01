package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/stretchr/testify/assert"
)

var oreilly = models.VendorWithName("O'Reilly")

func TestOreillyParsesContents(t *testing.T) {
	filename := "oreilly.xml"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		deal := oreilly.OReilly(contents)
		assert.Equal(t, "#Ebook Deal/Day: Social Sensing -\n$49.98 (Save 50%)\n\n\n Use code DEAL", deal.Title)
	}
}

func TestOReillyEmptyReturnsNoResults(t *testing.T) {
	deal := oreilly.OReilly([]byte{})
	assert.Equal(t, "No Results", deal.Title)
}
