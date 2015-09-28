package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/hoop33/tbdotd/app/models"
)

var oreillyBusiness = models.VendorWithName("O'Reilly Business")

func TestOreillyBusinessParsesContents(t *testing.T) {
	filename := "oreilly_business.xml"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		deal := oreillyBusiness.OReillyBusiness(contents)
		assert.Equal(t, "#Video Deal/Day: Software Architecture Fundamentals People Skills -\n$39.99 (Save 50%)\n\n\n Use code MSDEAL", deal.Title)
	}
}

func TestOReillyBusinessEmptyReturnsNoResults(t *testing.T) {
	deal := oreillyBusiness.OReillyBusiness([]byte{})
	assert.Equal(t, "No Results", deal.Title)
}
