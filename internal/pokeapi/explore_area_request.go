package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreAreaResponse(area string) (ExploreAreaResponse, error) {
	fullURL := baseURL + "/location-area/" + area

	if entry, exists := c.cache.Get(fullURL); exists {
		res := ExploreAreaResponse{}
		err := json.Unmarshal(entry, &res)
		if err != nil {
			return ExploreAreaResponse{}, err
		}
		return res, nil
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return ExploreAreaResponse{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return ExploreAreaResponse{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return ExploreAreaResponse{}, fmt.Errorf("area not found")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ExploreAreaResponse{}, err
	}
	exploreAreaResponse := ExploreAreaResponse{}
	err = json.Unmarshal(body, &exploreAreaResponse)
	if err != nil {
		return ExploreAreaResponse{}, err
	}
	c.cache.Add(fullURL, body)
	return exploreAreaResponse, nil
}
