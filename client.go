package izi_client

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"net/url"
	"time"
)

// Config configure the Client
type Config struct {

	// Host is the host
	Host string

	// APIKey is optional
	APIKey string
}

type Client struct {
	config     Config
	httpClient *fasthttp.Client

	city   *cityClient
	object *objectClient
}

func NewClient(cfg Config) (*Client, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}
	client := &fasthttp.Client{
		ReadTimeout:  60 * time.Millisecond,
		WriteTimeout: 60 * time.Millisecond,
	}

	c := &Client{
		config:     cfg,
		httpClient: client,
	}

	c.city = &cityClient{c}
	c.object = &objectClient{c}
	return c, nil
}

func NewCustomClient(cfg Config, c *fasthttp.Client) (*Client, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}
	cli := &Client{
		config:     cfg,
		httpClient: c,
	}
	return cli, nil
}

func validateConfig(c Config) error {
	if c.APIKey == "" {
		return fmt.Errorf("API key can not be empty")
	}
	if c.Host == "" {
		c.Host = defaultApiHost
	}
	if _, err := url.Parse(c.Host); err != nil {
		return err
	}
	return nil
}

func (c *Client) Object() *objectClient {
	return c.object
}

func (c *Client) City() *cityClient {
	return c.city
}

func (c *Client) executeRequest(req internalRequest) error {
	internalError := &Error{
		Endpoint:           req.endpoint,
		Method:             req.method,
		Function:           req.functionName,
		APIName:            req.apiName,
		RequestToString:    "empty request",
		ResponseToString:   "empty response",
		StatusCodeExpected: req.acceptedStatusCodes,
	}

	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)
	err := c.sendRequest(&req, internalError, response)
	if err != nil {
		return err
	}
	internalError.StatusCode = response.StatusCode()

	err = c.handleStatusCode(&req, response, internalError)
	if err != nil {
		return err
	}

	err = c.handleResponse(&req, response, internalError)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) sendRequest(req *internalRequest, internalError *Error, response *fasthttp.Response) error {
	var (
		request *fasthttp.Request

		err error
	)

	// Setup URL
	requestURL, err := url.Parse(c.config.Host + req.endpoint)
	if err != nil {
		return errors.Wrap(err, "unable to parse url")
	}

	// Build query parameters
	if req.withQueryParams != nil {
		query := requestURL.Query()
		for key, value := range req.withQueryParams {
			query.Set(key, value)
		}

		requestURL.RawQuery = query.Encode()
	}

	request = fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	request.SetRequestURI(requestURL.String())
	request.Header.SetMethod(req.method)

	if req.withRequest != nil {

		// A json request is mandatory, so the request interface{} need to be passed as a raw json body.
		rawJSONRequest := req.withRequest
		var data []byte
		var err error
		if raw, ok := rawJSONRequest.(json.Marshaler); ok {
			data, err = raw.MarshalJSON()
		} else {
			data, err = json.Marshal(rawJSONRequest)
		}
		internalError.RequestToString = string(data)
		if err != nil {
			return internalError.WithErrCode(ErrCodeMarshalRequest, err)
		}
		request.SetBody(data)
	}

	// adding request headers
	request.Header.Set("Content-Type", "application/json")
	versionHeader := fmt.Sprintf("application/izi-api-v%s+json", apiVersion)
	request.Header.Set("Accept", versionHeader)
	request.Header.Set(apiKeyHeaderName, c.config.APIKey)

	// request is sent
	err = c.httpClient.Do(request, response)

	// request execution fail
	if err != nil {
		return internalError.WithErrCode(ErrCodeRequestExecution, err)
	}

	return nil
}

func (c *Client) handleStatusCode(req *internalRequest, response *fasthttp.Response, internalError *Error) error {
	if req.acceptedStatusCodes != nil {

		// A successful status code is required so check if the response status code is in the
		// expected status code list.
		for _, acceptedCode := range req.acceptedStatusCodes {
			if response.StatusCode() == acceptedCode {
				return nil
			}
		}
		// At this point the response status code is a failure.
		rawBody := response.Body()

		internalError.ErrorBody(rawBody)

		return internalError.WithErrCode(ErrCodeResponseStatusCode)
	}

	return nil
}

func (c *Client) handleResponse(req *internalRequest, response *fasthttp.Response, internalError *Error) (err error) {
	if req.withResponse != nil {

		// A json response is mandatory, so the response interface{} need to be unmarshal from the response payload.
		rawBody := response.Body()
		internalError.ResponseToString = string(rawBody)

		var err error
		if resp, ok := req.withResponse.(json.Unmarshaler); ok {
			err = resp.UnmarshalJSON(rawBody)
			req.withResponse = resp
		} else {
			err = json.Unmarshal(rawBody, req.withResponse)
		}
		if err != nil {
			return internalError.WithErrCode(ErrCodeResponseUnmarshalBody, err)
		}
	}
	return nil
}
