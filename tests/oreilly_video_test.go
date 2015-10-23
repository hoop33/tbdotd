package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/app/models"
	"github.com/stretchr/testify/assert"
)

func TestOreillyVideoParsesContents(t *testing.T) {
	filename := "oreilly_video.xml"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		vendor := models.Vendors["O'Reilly"]
		source := vendor.Sources["O'Reilly Video"]
		deal := source.OReillyVideo(&vendor, contents)
		assert.Equal(t, "Porting from Python 2 to Python 3 -\n$24.98 (Save 50%)\n\n\n Use code VDWK", deal.Title)
		assert.Equal(t, "http://shop.oreilly.com/product/0636920042525.do#2015-09-23feed", deal.Url)
	}
}

func TestOReillyVideoEmptyReturnsNoResults(t *testing.T) {
	vendor := models.Vendors["O'Reilly"]
	source := vendor.Sources["O'Reilly Video"]
	deal := source.OReillyVideo(&vendor, []byte{})
	assert.Equal(t, "No Deal", deal.Title)
}
