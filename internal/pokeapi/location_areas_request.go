package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	if entry, exists := c.cache.Get(fullURL); exists {
		res := LocationAreasResponse{}
		err := json.Unmarshal(entry, &res)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return res, nil
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return LocationAreasResponse{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, res.Body)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	locationAreasRes := LocationAreasResponse{}
	err = json.Unmarshal(body, &locationAreasRes)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	c.cache.Add(fullURL, body)
	return locationAreasRes, nil
}
