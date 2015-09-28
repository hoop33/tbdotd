package tests

import (
	"io/ioutil"
	"testing"

	"github.com/hoop33/tbdotd/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/hoop33/tbdotd/app/models"
)

var oreillyVideo = models.VendorWithName("O'Reilly Video")

func TestOreillyVideoParsesContents(t *testing.T) {
	filename := "oreilly_video.xml"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fail()
	} else {
		deal := oreillyVideo.OReillyVideo(contents)
		assert.Equal(t, "#Video Deal/Week: Porting from Python 2 to Python 3 -\n$24.98 (Save 50%)\n\n\n Use code VDWK", deal.Title)
	}
}

func TestOReillyVideoEmptyReturnsNoResults(t *testing.T) {
	deal := oreillyVideo.OReillyVideo([]byte{})
	assert.Equal(t, "No Results", deal.Title)
}
