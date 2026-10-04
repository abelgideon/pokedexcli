package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocations(url string) (MapList, error) {
	var maps MapList

	data, found := c.cache.Get(url)
	if found {
		jsonErr := json.Unmarshal(data, &maps)
		if jsonErr != nil {
			return MapList{}, fmt.Errorf("failed to convert json: %v", jsonErr)
		}
		return maps, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return MapList{}, fmt.Errorf("failed to create request: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return MapList{}, fmt.Errorf("failed to make request: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return MapList{}, fmt.Errorf("unexpected HTTP status: %s", res.Status)
	}

	content, err := io.ReadAll(res.Body)
	if err != nil {
		return MapList{}, fmt.Errorf("failed to read maps: %v", err)
	}

	jsonErr := json.Unmarshal(content, &maps)
	if jsonErr != nil {
		return MapList{}, fmt.Errorf("failed to convert json: %v", jsonErr)
	}

	c.cache.Add(url, content)

	return maps, nil
}
