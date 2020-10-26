package httplib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullURL(t *testing.T) {
	basURL := "https://pan.baidu.com/rest/2.0/xpan/nas"
	fullURL := NewFullURL(basURL, nil)
	fullURL.Add("access_token", "123")
	assert.Equal(t, fullURL.String(), "https://pan.baidu.com/rest/2.0/xpan/nas?access_token=123")

	basURL = "https://pan.baidu.com/rest/2.0/xpan/nas"
	fullURL = NewFullURL(basURL, map[string]string{"access_token": "123"})
	assert.Equal(t, fullURL.String(), "https://pan.baidu.com/rest/2.0/xpan/nas?access_token=123")

	basURL = "https://pan.baidu.com/rest/2.0/xpan/nas?bbc=123"
	fullURL = NewFullURL(basURL, map[string]string{"access_token": "123"})
	assert.Equal(t, fullURL.String(), "https://pan.baidu.com/rest/2.0/xpan/nas?access_token=123&bbc=123")

	basURL = "https://pan.baidu.com/rest/2.0/xpan/nas"
	fullURL = NewFullURL(basURL, map[string]string{"access_token": "123"})
	fullURL.Add("bbc", "123")
	assert.Equal(t, fullURL.String(), "https://pan.baidu.com/rest/2.0/xpan/nas?access_token=123&bbc=123")
}
