package caddyapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/load" {
			t.Errorf("expected /load but got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST but got %s", r.Method)
		}
	}))
	defer server.Close()

	caddy := NewCaddyAPI(server.URL)
	err := caddy.LoadConfig(Config{})
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
}

func TestStopServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/stop" {
			t.Errorf("expected /stop but got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST but got %s", r.Method)
		}
	}))
	defer server.Close()

	caddy := NewCaddyAPI(server.URL)
	err := caddy.StopServer()
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
}

func TestGetConfig(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/config/test" {
			t.Errorf("expected /config/test but got %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET but got %s", r.Method)
		}
		fmt.Fprintln(w, "{}")
	}))
	defer server.Close()

	caddy := NewCaddyAPI(server.URL)
	_, err := caddy.GetConfig("test")
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
}

// Similarly, you can write tests for PostConfig, PatchConfig, PutConfig, and DeleteConfig
