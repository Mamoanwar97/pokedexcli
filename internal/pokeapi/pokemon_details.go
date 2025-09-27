package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetPokemonDetails -
func (c *Client) PokemonDetails(name string) (RespShallowPokemonDetails, error) {
	url := baseURL + "/pokemon/" + name

	cachedData, exists := c.cache.Get(url)

	if exists {
		return unmarshalPokemonDetails(cachedData)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemonDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemonDetails{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemonDetails{}, err
	}

	c.cache.Add(url, dat)

	return unmarshalPokemonDetails(dat)
}

func unmarshalPokemonDetails(data []byte) (RespShallowPokemonDetails, error) {
	pokemonDetailsResp := RespShallowPokemonDetails{}
	err := json.Unmarshal(data, &pokemonDetailsResp)
	if err != nil {
		return RespShallowPokemonDetails{}, err
	}
	return pokemonDetailsResp, nil
}
