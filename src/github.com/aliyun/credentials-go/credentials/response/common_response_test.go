package response

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseFromHTTPResponse(t *testing.T) {
	r := &CommonResponse{}
	res := &http.Response{
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
		StatusCode: 200,
		Header:     make(map[string][]string),
	}
	res.Header.Add("Server", "GitHub.com")
	r.ParseFromHTTPResponse(res)
	assert.Equal(t, []byte{}, r.GetHTTPContentBytes())
	assert.Equal(t, "", r.GetHTTPContentString())
	assert.Equal(t, "GitHub.com", r.GetHTTPHeaders()["Server"][0])
	assert.Equal(t, 200, r.GetHTTPStatus())
}

func TestHookReadAll(t *testing.T) {
	fn := func(body io.Reader) (byt []byte, err error) {
		return nil, errors.New("hookReadAll")
	}
	result := hookReadAll(fn)
	byt, err := result(nil)
	assert.Nil(t, byt)
	assert.Equal(t, "hookReadAll", err.Error())

	originHookReadAll := hookReadAll
	hookReadAll = func(old func(body io.Reader) (byt []byte, err error)) func(body io.Reader) (byt []byte, err error) {
		return fn
	}
	defer func() {
		hookReadAll = originHookReadAll
	}()
	commonResponse := &CommonResponse{}
	httpResponse := &http.Response{
		Body: ioutil.NopCloser(strings.NewReader("creadential")),
	}
	err = commonResponse.ParseFromHTTPResponse(httpResponse)
	assert.Equal(t, "hookReadAll", err.Error())
}
