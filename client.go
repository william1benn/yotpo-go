package yotpo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
)

//client_credentials

type YotpoClient struct {
	AppKey    string
	ApiSecret string
	BaseUrl   *url.URL
}

func NewYotpoClient(yotpoAppKey, yotpoSecret string) YotpoClient {
	url, err := url.Parse("https://api.yotpo.com")
	CheckErrorFatal(err)
	return YotpoClient{
		AppKey:    yotpoAppKey,
		ApiSecret: yotpoSecret,
		BaseUrl:   url,
	}
}

func (y *YotpoClient) CreateUtoken() (YotpoTokenResponse, error) {
	yotpoTokenResponse := YotpoTokenResponse{}
	yotpoGrant := &YotpoTokenRequest{
		ClientId:     y.AppKey,
		ClientSecret: y.ApiSecret,
		GrantType:    "client_credentials",
	}
	grantTypePayload, _ := json.Marshal(yotpoGrant)
	grantResponse, err := y.PostRequest("/oauth/token", grantTypePayload, nil, nil)
	if err != nil {
		return YotpoTokenResponse{}, err
	}

	err = json.NewDecoder(grantResponse.Body).Decode(&yotpoTokenResponse)
	if err != nil {
		return YotpoTokenResponse{}, err
	}
	return yotpoTokenResponse, err
}

func (y *YotpoClient) AddUrlParams(params interface{}) string {
	values, err := query.Values(params)
	CheckErrorFatal(err)
	return values.Encode()
}

func (y *YotpoClient) SendRequests(httpMethod, endpointPath string, body io.Reader, params interface{}, utoken *string) (*http.Response, error) {
	clientHttp := &http.Client{}
	requestUrl := y.BaseUrl.JoinPath(endpointPath)

	if utoken != nil && params != nil {
		tokenParams := y.AddUrlParams(UToken{Token: *utoken})
		submitedUrlParams := y.AddUrlParams(params)
		requestUrl.RawQuery = fmt.Sprintf("%s&%s", submitedUrlParams, tokenParams)

	} else if utoken != nil && params == nil {
		requestUrl.RawQuery = y.AddUrlParams(UToken{Token: *utoken})

	} else if utoken == nil && params != nil {
		requestUrl.RawQuery = y.AddUrlParams(params)

	}
	fmt.Println(requestUrl)
	request, err := http.NewRequest(httpMethod, requestUrl.String(), body)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	sentRequest, err := clientHttp.Do(request)
	if err != nil {
		return nil, err
	}
	return sentRequest, nil
}

func (y *YotpoClient) GetRequest(endpointPath string, urlParams interface{}, uToken *string) (*http.Response, error) {
	apiResponse, err := y.SendRequests("GET", endpointPath, nil, urlParams, uToken)
	if err != nil {
		return nil, err
	}
	return apiResponse, nil
}

func (y *YotpoClient) PostRequest(endpointPath string, reqBody []byte, urlParams interface{}, uToken *string) (*http.Response, error) {
	apiResponse, err := y.SendRequests("POST", endpointPath, bytes.NewReader(reqBody), urlParams, uToken)
	if err != nil {
		return nil, err
	}
	return apiResponse, nil
}

type YotpoTokenRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}
type YotpoTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type UToken struct {
	Token string `url:"utoken"`
}
