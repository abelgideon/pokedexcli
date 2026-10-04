package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(locationName string) (LocationArea, error) {
	var poks LocationArea
	url := baseURL + "/location-area/" + locationName

	data, found := c.cache.Get(url)
	if found {
		jsonErr := json.Unmarshal(data, &poks)
		if jsonErr != nil {
			return LocationArea{}, fmt.Errorf("failed to convert json: %v", jsonErr)
		}
		return poks, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return LocationArea{}, fmt.Errorf("failed to create request: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, fmt.Errorf("failed to make request: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return LocationArea{}, fmt.Errorf("unexpected HTTP status: %s", res.Status)
	}

	content, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, fmt.Errorf("failed to read pokemon: %v", err)
	}

	jsonErr := json.Unmarshal(content, &poks)
	if jsonErr != nil {
		return LocationArea{}, fmt.Errorf("failed to convert json: %v", jsonErr)
	}

	c.cache.Add(url, content)

	return poks, nil
}
