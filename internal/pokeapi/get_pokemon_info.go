package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {

	endpoint := "/pokemon/" + name
	fullURL := baseURL + endpoint
	if entry, exists := c.cache.Get(fullURL); exists {
		res := Pokemon{}
		err := json.Unmarshal(entry, &res)
		if err != nil {
			return Pokemon{}, err
		}
		return res, nil
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("pokemon not found")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	pokemonResponse := Pokemon{}
	err = json.Unmarshal(body, &pokemonResponse)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(fullURL, body)
	return pokemonResponse, nil
}
