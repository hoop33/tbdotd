package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/stretchr/testify/assert"
)

var informIT = models.VendorWithName("InformIT")

func TestInformITParsesContents(t *testing.T) {
	filename := "informit.xml"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		deal := informIT.InformIT(contents)
		assert.Equal(t, "eBook Deal of the Day ::\n\t\t\t\tMMIX Supplement, The: Supplement to The Art of Computer Programming Volumes 1, 2, 3 by Donald E. Knuth by Martin Ruckert", deal.Title)
		assert.Equal(t, int64(1442991600), deal.Date.Unix())
	}
}

func TestInformITEmptyReturnsNoResults(t *testing.T) {
	deal := informIT.InformIT([]byte{})
	assert.Equal(t, "No Results", deal.Title)
}
