/*
 * BitMEX API
 *
 * ## REST API for the BitMEX Trading Platform  [Changelog](/app/apiChangelog)    #### Getting Started   ##### Fetching Data  All REST endpoints are documented below. You can try out any query right from this interface.  Most table queries accept `count`, `start`, and `reverse` params. Set `reverse=true` to get rows newest-first.  Additional documentation regarding filters, timestamps, and authentication is available in [the main API documentation](https://www.bitmex.com/app/restAPI).  *All* table data is available via the [Websocket](/app/wsAPI). We highly recommend using the socket if you want to have the quickest possible data without being subject to ratelimits.  ##### Return Types  By default, all data is returned as JSON. Send `?_format=csv` to get CSV data or `?_format=xml` to get XML data.  ##### Trade Data Queries  *This is only a small subset of what is available, to get you started.*  Fill in the parameters and click the `Try it out!` button to try any of these queries.  * [Pricing Data](#!/Quote/Quote_get)  * [Trade Data](#!/Trade/Trade_get)  * [OrderBook Data](#!/OrderBook/OrderBook_getL2)  * [Settlement Data](#!/Settlement/Settlement_get)  * [Exchange Statistics](#!/Stats/Stats_history)  Every function of the BitMEX.com platform is exposed here and documented. Many more functions are available.  -  ## All API Endpoints  Click to expand a section.
 *
 * OpenAPI spec version: 1.2.0
 * Contact: support@bitmex.com
 * Generated by: https://github.com/sunrisedo/bitmex.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rest

import (
	"encoding/base64"
	"net/http"

	"github.com/go-resty/resty"
	"golang.org/x/net/proxy"
)

type Configuration struct {
	UserName      string            `json:"userName,omitempty"`
	Password      string            `json:"password,omitempty"`
	APIKeyPrefix  map[string]string `json:"APIKeyPrefix,omitempty"`
	APIKey        map[string]string `json:"APIKey,omitempty"`
	debug         bool              `json:"debug,omitempty"`
	DebugFile     string            `json:"debugFile,omitempty"`
	OAuthToken    string            `json:"oAuthToken,omitempty"`
	Timeout       int               `json:"timeout,omitempty"`
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	AccessToken   string            `json:"accessToken,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	APIClient     APIClient         `json:"APIClient,omitempty"`

	//add by qct
	ExpireTime int64
	ApiKey     string
	SecretKey  string
}

func NewConfiguration() *Configuration {
	return &Configuration{
		BasePath:      "https://www.bitmex.com/api/v1",
		UserName:      "",
		debug:         false,
		DefaultHeader: make(map[string]string),
		APIKey:        make(map[string]string),
		APIKeyPrefix:  make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}
}

func NewConfigurationWithKey(apiKey, secretKey string) *Configuration {
	return &Configuration{
		BasePath:      "https://www.bitmex.com/api/v1",
		UserName:      "",
		debug:         false,
		DefaultHeader: make(map[string]string),
		APIKey:        make(map[string]string),
		APIKeyPrefix:  make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
		ApiKey:        apiKey,
		SecretKey:     secretKey,
		ExpireTime:    5,
	}
}

func NewConfigurationWithKeyExpires(apiKey, secretKey string, expires int64) *Configuration {
	return &Configuration{
		BasePath:      "https://www.bitmex.com/api/v1",
		UserName:      "",
		debug:         false,
		DefaultHeader: make(map[string]string),
		APIKey:        make(map[string]string),
		APIKeyPrefix:  make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
		ApiKey:        apiKey,
		SecretKey:     secretKey,
		ExpireTime:    expires,
	}
}

func (c *Configuration) GetBasicAuthEncodedString() string {
	return base64.StdEncoding.EncodeToString([]byte(c.UserName + ":" + c.Password))
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

func (c *Configuration) GetAPIKeyWithPrefix(APIKeyIdentifier string) string {
	if c.APIKeyPrefix[APIKeyIdentifier] != "" {
		return c.APIKeyPrefix[APIKeyIdentifier] + " " + c.APIKey[APIKeyIdentifier]
	}

	return c.APIKey[APIKeyIdentifier]
}

func (c *Configuration) SetDebug(enable bool) {
	c.debug = enable
}

func (c *Configuration) GetDebug() bool {
	return c.debug
}

func (c *Configuration) Proxy(dialer proxy.Dialer) *Configuration {
	if dialer != nil {
		resty.SetTransport(&http.Transport{Dial: dialer.Dial})
	}
	return c
}

func (c *Configuration) SetTest() *Configuration {
	c.BasePath = "https://testnet.bitmex.com/api/v1"
	return c
}
