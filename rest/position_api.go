/*
 * BitMEX API
 *
 * ## REST API for the BitMEX Trading Platform  [Changelog](/app/apiChangelog)    #### Getting Started   ##### Fetching Data  All REST endpoints are documented below. You can try out any query right from this interface.  Most table queries accept `count`, `start`, and `reverse` params. Set `reverse=true` to get rows newest-first.  Additional documentation regarding filters, timestamps, and authentication is available in [the main API documentation](https://www.bitmex.com/app/restAPI).  *All* table data is available via the [Websocket](/app/wsAPI). We highly recommend using the socket if you want to have the quickest possible data without being subject to ratelimits.  ##### Return Types  By default, all data is returned as JSON. Send `?_format=csv` to get CSV data or `?_format=xml` to get XML data.  ##### Trade Data Queries  *This is only a small subset of what is available, to get you started.*  Fill in the parameters and click the `Try it out!` button to try any of these queries.  * [Pricing Data](#!/Quote/Quote_get)  * [Trade Data](#!/Trade/Trade_get)  * [OrderBook Data](#!/OrderBook/OrderBook_getL2)  * [Settlement Data](#! /Settlement/Settlement_get)  * [Exchange Statistics](#!/Stats/Stats_history)  Every function of the BitMEX.com platform is exposed here and documented. Many more functions are available.  -  ## All API Endpoints Click to expand a section.
 *
 * OpenAPI spec version: 1.2.0
 * Contact: support@bitmex.com
 * Generated by: https://github.com/sunrisedo/bitmex.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not  use this file except in compliance with the License.
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
	"encoding/json"
	"errors"
	"fmt"
	// "log"
	"net/url"
	"strconv"
)

type PositionApi struct {
	Configuration Configuration
}

func NewPositionApi() *PositionApi {
	configuration := NewConfiguration()
	return &PositionApi{Configuration: *configuration}
}

func NewPositionApiWithConfig(config *Configuration) *PositionApi {
	return &PositionApi{Configuration: *config}
}

func NewPositionApiWithBasePath(basePath string) *PositionApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &PositionApi{Configuration: *configuration}
}

/**
 * Get your positions.
 * See &lt;a href&#x3D;\&quot;http://www.onixs.biz/fix-dictionary/5.0.SP2/msgType_AP_6580.html\&quot;&gt;the FIX Spec&lt;/a&gt; for explanations of these fields.
 *
 * @param filter Table filter. For example, send {\&quot;symbol\&quot;: \&quot;XBT24H\&quot;}.
 * @param columns Which columns to fetch. For example, send [\&quot;columnName\&quot;].
 * @param count Number of rows to fetch.
 * @return []Position
 */
func (a PositionApi) PositionGet(filter string, columns string, count float32) ([]Position, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/position"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	queryParams.Add("filter", a.Configuration.APIClient.ParameterToString(filter, ""))
	queryParams.Add("columns", a.Configuration.APIClient.ParameterToString(columns, ""))
	queryParams.Add("count", a.Configuration.APIClient.ParameterToString(fmt.Sprintf("%.0f", count), ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		"text/xml",
		"application/javascript",
		"text/javascript",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	SetApiHeader(headerParams, &a.Configuration, httpMethod, path, formParams, queryParams)

	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return nil, NewAPIResponse(httpResponse.RawResponse), err
	}
	if httpResponse.StatusCode() != 200 {
		rep := new(ModelError)
		json.Unmarshal(httpResponse.Body(), &rep)
		return nil, NewAPIResponse(httpResponse.RawResponse), errors.New(fmt.Sprintf("%s,%s", rep.Error_.Name, rep.Error_.Message))
	}
	rep := new([]Position)
	// log.Println("PositionGet Json:", string(httpResponse.Body()))
	err = json.Unmarshal(httpResponse.Body(), &rep)
	return *rep, NewAPIResponse(httpResponse.RawResponse), err
	// err = json.Unmarshal(httpResponse.Body(), &successPayload)
	// return *successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Enable isolated margin or cross margin per-position.
 * On Speculative (DPE-Enabled) contracts, users can switch isolate margin per-position. This function allows switching margin isolation (aka fixed margin) on and off.
 *
 * @param symbol Position symbol to isolate.
 * @param enabled True for isolated margin, false for cross margin.
 * @return *Position
 */
func (a PositionApi) PositionIsolateMargin(symbol string, enabled bool) (*Position, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/position/isolate"

	// verify the required parameter 'symbol' is set
	if &symbol == nil {
		return new(Position), nil, errors.New("Missing required parameter 'symbol' when calling PositionApi->PositionIsolateMargin")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		"text/xml",
		"application/javascript",
		"text/javascript",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}

	formParams["symbol"] = symbol
	formParams["enabled"] = strconv.FormatBool(enabled)
	var successPayload = new(Position)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Transfer equity in or out of a position.
 * When margin is isolated on a position, use this function to add or remove margin from the position. Note that you cannot remove margin below the initial margin threshold.
 *
 * @param symbol Symbol of position to isolate.
 * @param amount Amount to transfer, in Satoshis. May be negative.
 * @return *Position
 */
func (a PositionApi) PositionTransferIsolatedMargin(symbol string, amount float32) (*Position, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/position/transferMargin"

	// verify the required parameter 'symbol' is set
	if &symbol == nil {
		return new(Position), nil, errors.New("Missing required parameter 'symbol' when calling PositionApi->PositionTransferIsolatedMargin")
	}
	// verify the required parameter 'amount' is set
	if &amount == nil {
		return new(Position), nil, errors.New("Missing required parameter 'amount' when calling PositionApi->PositionTransferIsolatedMargin")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		"text/xml",
		"application/javascript",
		"text/javascript",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}

	formParams["symbol"] = symbol
	formParams["amount"] = fmt.Sprintf("%.0f", amount)
	var successPayload = new(Position)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Choose leverage for a position.
 * On Speculative (DPE-Enabled) contracts, users can choose an isolated leverage. This will automatically enable isolated margin.
 *
 * @param symbol Symbol of position to adjust.
 * @param leverage Leverage value. Send a number between 0.01 and 100 to enable isolated margin with a fixed leverage. Send 0 to enable cross margin.
 * @return *Position
 */
func (a PositionApi) PositionUpdateLeverage(symbol string, leverage float64) (*Position, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/position/leverage"

	// verify the required parameter 'symbol' is set
	if &symbol == nil {
		return new(Position), nil, errors.New("Missing required parameter 'symbol' when calling PositionApi->PositionUpdateLeverage")
	}
	// verify the required parameter 'leverage' is set
	if &leverage == nil {
		return new(Position), nil, errors.New("Missing required parameter 'leverage' when calling PositionApi->PositionUpdateLeverage")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		"text/xml",
		"application/javascript",
		"text/javascript",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}

	formParams["symbol"] = symbol
	formParams["leverage"] = fmt.Sprintf("%.2f", leverage)
	var successPayload = new(Position)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Update your risk limit.
 * Risk Limits limit the size of positions you can trade at various margin levels. Larger positions require more margin. Please see the Risk Limit documentation for more details.
 *
 * @param symbol Symbol of position to isolate.
 * @param riskLimit New Risk Limit, in Satoshis.
 * @return *Position
 */
func (a PositionApi) PositionUpdateRiskLimit(symbol string, riskLimit float32) (*Position, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/position/riskLimit"

	// verify the required parameter 'symbol' is set
	if &symbol == nil {
		return new(Position), nil, errors.New("Missing required parameter 'symbol' when calling PositionApi->PositionUpdateRiskLimit")
	}
	// verify the required parameter 'riskLimit' is set
	if &riskLimit == nil {
		return new(Position), nil, errors.New("Missing required parameter 'riskLimit' when calling PositionApi->PositionUpdateRiskLimit")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		"text/xml",
		"application/javascript",
		"text/javascript",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}

	formParams["symbol"] = symbol
	formParams["riskLimit"] = fmt.Sprintf("%.0f", riskLimit)
	var successPayload = new(Position)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}
