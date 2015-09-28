package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/hoop33/tbdotd/app/models"
)

var apress = models.VendorWithName("Apress")

func TestApressParsesContents(t *testing.T) {
	filename := "apress.html"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		deal := apress.Apress(contents)
		assert.Equal(t, "Android Best Practices", deal.Title)
	}
}

func TestApressEmptyReturnsNoResults(t *testing.T) {
	deal := apress.Apress([]byte{})
	assert.Equal(t, "No Results", deal.Title)
}
