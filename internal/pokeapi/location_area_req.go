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

	data, ok := c.cache.Get(fullURL)
	if ok {
		locationAreaResp := LocationAreaResp{}
		err := json.Unmarshal(data, &locationAreaResp)
		if err != nil {
			return LocationAreaResp{}, err
		}

		fmt.Print("+++++++++++ cache hit! ++++++++++++++\n")
		return locationAreaResp, nil
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

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreaResp := LocationAreaResp{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	fmt.Print("*********** cache miss! ***********\n")
	c.cache.Add(fullURL, data)
	return locationAreaResp, nil
}

func (c *Client) GetLocationArea(locationId string) (LocationArea, error) {
	endpoint := "/location-area/" + locationId
	fullURL := baseURL + endpoint

	data, ok := c.cache.Get(fullURL)
	if ok {
		locationArea := LocationArea{}
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		fmt.Print("+++++++++++ cache hit! ++++++++++++++\n")
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	fmt.Print("*********** cache miss! ***********\n")
	c.cache.Add(fullURL, data)
	return locationArea, nil
}
