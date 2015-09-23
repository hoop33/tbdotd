package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/stretchr/testify/assert"
)

var vendor = models.VendorWithName("InformIT")

func TestInformITParsesContents(t *testing.T) {
	filename := "informit.xml"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		deal := vendor.InformIT(contents)
		assert.Equal(t, deal.Title, "eBook Deal of the Day ::\n\t\t\t\tMMIX Supplement, The: Supplement to The Art of Computer Programming Volumes 1, 2, 3 by Donald E. Knuth by Martin Ruckert")
	}
}

func TestInformITEmptyReturnsNoResults(t *testing.T) {
	deal := vendor.InformIT([]byte{})
	assert.Equal(t, deal.Title, "No Results")
}
