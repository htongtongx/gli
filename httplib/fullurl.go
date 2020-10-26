package httplib

import "net/url"

type FullURL struct {
	URL    string
	Params map[string]string
}

func NewFullURL(url string, params map[string]string) (fullURL *FullURL) {
	if params == nil {
		params = make(map[string]string)
	}
	fullURL = &FullURL{URL: url, Params: params}
	return
}

func (u *FullURL) String() (urlStr string) {
	params := url.Values{}
	urlEntity, err := url.Parse(u.URL)
	if err != nil {
		return err.Error()
	}
	for k, v := range u.Params {
		params.Set(k, v)
	}
	urlEntity.RawQuery = params.Encode()
	urlStr = urlEntity.String()
	return urlStr
}

func (u *FullURL) Add(key, value string) {
	u.Params[key] = value
}
