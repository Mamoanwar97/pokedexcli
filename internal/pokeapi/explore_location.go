package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ExploreLocation(locationName string) (RespShallowExploreLocation, error) {
	url := baseURL + "/location-area/" + locationName

	cachedData, exists := c.cache.Get(url)

	if exists {
		return unmarshalExploredLocation(cachedData)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowExploreLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowExploreLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowExploreLocation{}, err
	}

	c.cache.Add(url, dat)

	return unmarshalExploredLocation(dat)
}

func unmarshalExploredLocation(data []byte) (RespShallowExploreLocation, error) {
	exploreLocationResp := RespShallowExploreLocation{}
	err := json.Unmarshal(data, &exploreLocationResp)
	if err != nil {
		return RespShallowExploreLocation{}, err
	}
	return exploreLocationResp, nil
}
