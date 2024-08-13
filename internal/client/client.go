package client

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
	"go-recipe/internal/config"
)

type Client struct {
	baseURL string
	token   string
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		baseURL: cfg.APIBaseURL,
		token:   cfg.Token,
	}
}

func (c *Client) DoRequest(method, path string, requestBody []byte) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(c.baseURL + path)
	req.Header.Set("Authorization", c.token)
	req.Header.SetMethod(method)
	if requestBody != nil {
		req.SetBody(requestBody)
		req.Header.SetContentType("application/json")
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := fasthttp.Do(req, resp)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if statusCode := resp.StatusCode(); statusCode >= 400 {
		return nil, fmt.Errorf("received error response: %s", resp.Body())
	}

	log.Info().Msgf("response body: %s", resp.String())

	return resp.Body(), nil
}

func (c *Client) CreateUser(authorData interface{}) (string, error) {
	log.Info().Msgf("create user from CreateUser: %v", authorData)
	requestBody, err := json.Marshal(authorData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	log.Info().Msgf("request body: %s", requestBody)
	resp, err := c.DoRequest(fasthttp.MethodPost, "/v2/sign-up", requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %w", err)
	}

	var responseBody map[string]string
	if err = json.Unmarshal(resp, &responseBody); err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	log.Info().Msgf("response body: %v", responseBody)

	return responseBody["user_id"], nil
}
