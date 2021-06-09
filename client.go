package alibabaopen

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

var baseURL = "http://gw.open.1688.com/openapi/"

type Client struct {
	AppKey      string
	AppSecret   string
	AccessToken string
	method      string
	url         string
}

func (c *Client) SetMethod(method string) *Client {
	c.method = method
	return c
}

func (c *Client) Do(uri string, params map[string]string) ([]byte, error) {
	split := strings.Split(uri, ":")
	spacename := split[0]
	split = strings.Split(split[1], "-")
	apiname := split[0]
	version := split[1]
	urlInfo := fmt.Sprintf("param2/%s/%s/%s/", version, spacename, apiname)
	apiInfo := urlInfo + c.AppKey
	params["access_token"] = c.AccessToken

	var strs []string
	for k := range params {
		strs = append(strs, k)
	}
	sort.Strings(strs)
	var signParamsStr string
	v := url.Values{}
	for _, k := range strs {
		signParamsStr += k + params[k]
		v.Set(k, params[k])
	}
	signStr := apiInfo + signParamsStr
	codeSign := strings.ToUpper(HmacSHA1(c.AppSecret, signStr))
	if c.method == "GET" {
		v.Set("_aop_signature", codeSign)
		c.url = baseURL + apiInfo + "?" + v.Encode()
	} else {
		c.url = baseURL + apiInfo + "?_aop_signature=" + codeSign
	}
	postBody := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	resp, err := http.Post(c.url, "application/x-www-form-urlencoded", postBody)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// HmacSHA1 HmacSHA1
func HmacSHA1(key string, data string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}
