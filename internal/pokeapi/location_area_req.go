package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreaResp := LocationAreaResp{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	return locationAreaResp, nil
}
