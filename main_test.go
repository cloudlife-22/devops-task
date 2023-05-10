package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloServer(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(HelloServer))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	expected := "Hello World, !"
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if string(b) != expected {
		t.Errorf("expected %s but we got %s", expected, string(b))
	}
}

func TestHelloServerPath(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(HelloServer))
	resp, err := http.Get(server.URL + "/World")
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	expected := "Hello World, World!"
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if string(b) != expected {
		t.Errorf("expected %s but we got %s", expected, string(b))
	}
}
