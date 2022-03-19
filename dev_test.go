package dev

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestNewClient(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Errorf("Error loading env file: %s", err.Error())
	}

	token := os.Getenv("DEV_API_KEY")
	host := os.Getenv("DEV_HOST")

	t.Run("invalid token", func(t *testing.T) {
		client, err := NewClient(Options{Token: ""})

		if client != nil {
			t.Errorf("expected client to be nil, got %+v", client)
		}
		if err == nil {
			t.Errorf("expected {NewClient} to return invalid token error")
		}
	})

	t.Run("valid api", func(t *testing.T) {
		// t.Skip()
		client, err := NewClient(Options{Token: token})

		if client == nil {
			t.Errorf("expected client to not be nil")
		}
		if err != nil {
			t.Errorf("expected {NewClient} to run without error, got:\n %v", err)
		}
	})

	t.Run("base-url", func(t *testing.T) {
		// t.Skip()
		c, _ := NewClient(Options{Token: token})
		baseURL := c.BaseUrl
		if baseURL == nil {
			t.Errorf("expected baseUrl to be defined")
		}
		if baseURL.String() != BASE_URL {
			t.Errorf("expected baseUrl to be equal to `%s`", BASE_URL)
		}
	})

	t.Run("user-defined url", func(t *testing.T) {
		// t.Skip()
		c, _ := NewClient(Options{Token: token, Host: host})
		baseURL := c.BaseUrl
		if baseURL == nil {
			t.Errorf("expected baseUrl to be defined")
		}
		if baseURL.String() != host {
			t.Errorf("expected baseUrl to be equal to `%s`", host)
		}
	})
}
