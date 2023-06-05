package caddyapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CaddyAPI struct {
	ServerURL string
}

func NewCaddyAPI(serverURL string) *CaddyAPI {
	return &CaddyAPI{
		ServerURL: serverURL,
	}
}

func (c *CaddyAPI) LoadConfig(config Config) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	resp, err := http.Post(fmt.Sprintf("%s/load", c.ServerURL), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to load config: %s", body)
	}

	return nil
}

func (c *CaddyAPI) StopServer() error {
	resp, err := http.Post(fmt.Sprintf("%s/stop", c.ServerURL), "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to stop server: %s", body)
	}

	return nil
}

func (c *CaddyAPI) GetConfig(path string) (Config, error) {
	resp, err := http.Get(fmt.Sprintf("%s/config/%s", c.ServerURL, path))
	if err != nil {
		return Config{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return Config{}, fmt.Errorf("failed to get config: %s", body)
	}

	var config Config
	err = json.NewDecoder(resp.Body).Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *CaddyAPI) PostConfig(path string, config Config) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	resp, err := http.Post(fmt.Sprintf("%s/config/%s", c.ServerURL, path), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to post config: %s", body)
	}

	return nil
}

func (c *CaddyAPI) PatchConfig(path string, config Config) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/config/%s", c.ServerURL, path), bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to patch config: %s", body)
	}

	return nil
}

func (c *CaddyAPI) PutConfig(path string, config Config) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/config/%s", c.ServerURL, path), bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to put config: %s", body)
	}

	return nil
}

func (c *CaddyAPI) DeleteConfig(path string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/config/%s", c.ServerURL, path), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to delete config: %s", body)
	}

	return nil
}
