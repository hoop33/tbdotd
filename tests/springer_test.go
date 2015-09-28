package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/hoop33/tbdotd/app/models"
)

var springer = models.VendorWithName("Springer")

func TestSpringerParsesContents(t *testing.T) {
	filename := "apress.html"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		deal := springer.Springer(contents)
		assert.Equal(t, "Concurrent Programming: Algorithms, Principles, and Foundations", deal.Title)
	}
}

func TestSpringerEmptyReturnsNoResults(t *testing.T) {
	deal := springer.Springer([]byte{})
	assert.Equal(t, "No Results", deal.Title)
}
