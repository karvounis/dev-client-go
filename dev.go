package dev

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

const (
	BASE_URL = "https://dev.to/api"
)

type Client struct {
	Client  *http.Client
	BaseUrl *url.URL
	Token   string
}

type Options struct {
	Host  string
	Token string
}

func NewClient(options Options) (*Client, error) {
	host := BASE_URL
	if options.Host != "" {
		host = options.Host
	}

	u, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	if options.Token == "" {
		return nil, errors.New("invalid token")
	}

	c := &Client{
		Client:  http.DefaultClient,
		BaseUrl: u,
		Token:   options.Token,
	}

	return c, nil
}

func NewTestClient() (*Client, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	return NewClient(Options{Host: os.Getenv("DEV_HOST"), Token: os.Getenv("DEV_API_KEY")})
}

func (c *Client) NewRequest(ctx context.Context, method, path string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	url := c.BaseUrl.String() + path

	if payload != nil {
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}

		buf = bytes.NewBuffer(b)
	}

	return http.NewRequestWithContext(ctx, method, url, buf)
}

func (c *Client) SendHttpRequest(r *http.Request, v interface{}) error {
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("api-key", c.Token)

	resp, err := c.Client.Do(r)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err := extractDevError(resp)

		return err
	}

	if v == nil {
		return nil
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, v); err != nil {
		return err
	}

	return nil
}
