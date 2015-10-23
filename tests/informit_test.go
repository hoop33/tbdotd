package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/stretchr/testify/assert"
)

func TestInformITParsesContents(t *testing.T) {
	filename := "informit.xml"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		var vendor = models.Vendors["InformIT"]
		var source = vendor.Sources["InformIT"]
		deal := source.InformIT(&vendor, contents)
		assert.Equal(t, "MMIX Supplement, The: Supplement to The Art of Computer Programming Volumes 1, 2, 3 by Donald E. Knuth by Martin Ruckert", deal.Title)
		assert.Equal(t, "http://www.informit.com/deals/", deal.Url)
		assert.Equal(t, int64(1443078000), deal.ExpirationDate.Unix())
	}
}

func TestInformITParsesContentsWhenXMLIsMalformed(t *testing.T) {
	filename := "informit_bad.xml"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		var vendor = models.Vendors["InformIT"]
		var source = vendor.Sources["InformIT"]
		deal := source.InformIT(&vendor, contents)
		assert.Equal(t, "End-to-End QoS Network Design: Quality of Service for Rich-Media & Cloud Networks by Christina Hattingh, Robert Barton, Tim Szigeti, Kenneth Briley, Jr.", deal.Title)
		assert.Equal(t, "http://www.informit.com/deals/", deal.Url)
		assert.Equal(t, int64(1444978800), deal.ExpirationDate.Unix())
	}
}

func TestInformITEmptyReturnsNoResults(t *testing.T) {
	var vendor = models.Vendors["InformIT"]
	var source = vendor.Sources["InformIT"]
	deal := source.InformIT(&vendor, []byte{})
	assert.Equal(t, "No Deal", deal.Title)
}
