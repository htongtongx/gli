package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"

	JSON = "json"
	FORM = "form"
)

var (
	ErrMethodNotSupported = errors.New("method is not supported")
	ErrMIMENotSupported   = errors.New("mime is not supported")
)

// make query string from params
func MakeQueryStrFrom(params interface{}) (result string) {
	if params == nil {
		return
	}
	value := reflect.ValueOf(params)

	switch value.Kind() {
	case reflect.Struct:
		var formName string
		for i := 0; i < value.NumField(); i++ {
			if formName = value.Type().Field(i).Tag.Get("form"); formName == "" {
				// don't tag the form name, use camel name
				formName = GetCamelNameFrom(value.Type().Field(i).Name)
			}
			result += "&" + formName + "=" + fmt.Sprintf("%v", value.Field(i).Interface())
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			result += "&" + fmt.Sprintf("%v", key.Interface()) + "=" + fmt.Sprintf("%v", value.MapIndex(key).Interface())
		}
	default:
		return
	}

	if result != "" {
		result = result[1:]
	}
	return
}

func InvokeHandler(router *gin.Engine, req *http.Request) (bodyByte []byte, err error) {

	// initialize response record
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// extract the response from the response record
	result := w.Result()
	defer result.Body.Close()

	// extract response body
	bodyByte, err = ioutil.ReadAll(result.Body)

	return
}

func MakeRequest(method, mime, api string, param interface{}) (request *http.Request, err error) {
	method = strings.ToUpper(method)
	mime = strings.ToLower(mime)

	switch mime {
	case JSON:
		var (
			contentBuffer *bytes.Buffer
			jsonBytes     []byte
		)
		jsonBytes, err = json.Marshal(param)
		if err != nil {
			return
		}
		contentBuffer = bytes.NewBuffer(jsonBytes)
		request, err = http.NewRequest(string(method), api, contentBuffer)
		if err != nil {
			return
		}
		request.Header.Set("Content-Type", "application/json;charset=utf-8")
	case FORM:
		queryStr := MakeQueryStrFrom(param)
		var buffer io.Reader

		if (method == DELETE || method == GET) && queryStr != "" {
			api += "?" + queryStr
		} else {
			buffer = bytes.NewReader([]byte(queryStr))
		}

		request, err = http.NewRequest(string(method), api, buffer)
		if err != nil {
			return
		}
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	default:
		err = ErrMIMENotSupported
		return
	}
	return
}

func GetCamelNameFrom(name string) string {
	result := ""
	i := 0
	j := 0
	r := []rune(name)
	for m, v := range r {
		// if the char is the capital
		if v >= 'A' && v < 'a' {
			// if the prior is the lower-case || if the prior is the capital and the latter is the lower-case
			if (m != 0 && r[m-1] >= 'a') || ((m != 0 && r[m-1] >= 'A' && r[m-1] < 'a') && (m != len(r)-1 && r[m+1] >= 'a')) {
				i = j
				j = m
				result += name[i:j] + "_"
			}
		}
	}

	result += name[j:]
	return strings.ToLower(result)
}

func CheckBase(t *testing.T, r *gin.Engine, method, url string, postData interface{}, result interface{}, header ...map[string]string) []byte {
	return checkBase(t, r, JSON, method, url, postData, result, header...)
}

func checkBase(t *testing.T, r *gin.Engine, mime, method, url string, postData interface{}, result interface{}, header ...map[string]string) []byte {
	req, err := MakeRequest(method, mime, url, postData)
	if err != nil {
		assert.Equal(t, err, nil)
		return nil
	}
	if len(header) > 0 {
		for k, v := range header[0] {
			req.Header.Set(k, v)
		}
	}

	bodyByte, err := InvokeHandler(r, req)
	assert.Equal(t, err, nil)
	if result != nil {
		err = json.Unmarshal(bodyByte, result)
		assert.Equal(t, err, nil)
	}
	return bodyByte
}

func CheckBaseForm(t *testing.T, r *gin.Engine, method, url string, postData interface{}, result interface{}, header ...map[string]string) []byte {
	return checkBase(t, r, FORM, method, url, postData, result, header...)
}
