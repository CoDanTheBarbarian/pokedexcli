package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasRespose, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasRespose{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasRespose{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return LocationAreasRespose{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, res.Body)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasRespose{}, err
	}
	locationAreasRes := LocationAreasRespose{}
	err = json.Unmarshal(body, &locationAreasRes)
	if err != nil {
		return LocationAreasRespose{}, err
	}
	return locationAreasRes, nil
}
