package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Giira/pokedexcli/internal/pokecache"
)

const (
	apiUrlBase = "https://pokeapi.co/api/v2"
)

type Client struct {
	httpclient http.Client
}

func NewClient() Client {
	return Client{
		httpclient: http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *Client) GetLocAreas(sectionUrl *string, cache *pokecache.Cache) (LocAreas, error) {
	url := apiUrlBase + "/location-area"
	if sectionUrl != nil {
		url = *sectionUrl
	}

	body, ok := cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocAreas{}, err
		}

		res, err := c.httpclient.Do(req)
		if err != nil {
			return LocAreas{}, err
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return LocAreas{}, err
		}

		cache.Add(url, body)
	}

	locs := LocAreas{}
	err := json.Unmarshal(body, &locs)
	if err != nil {
		return LocAreas{}, err
	}

	return locs, nil
}

func (c *Client) GetAreaExplore(area *string, cache *pokecache.Cache) (AreaExplore, error) {
	url := apiUrlBase + "/location-area/" + *area

	body, ok := cache.Get((url))
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return AreaExplore{}, err
		}

		res, err := c.httpclient.Do(req)
		if err != nil {
			return AreaExplore{}, err
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return AreaExplore{}, err
		}

		cache.Add(url, body)
	}

	exps := AreaExplore{}
	err := json.Unmarshal(body, &exps)
	if err != nil {
		return AreaExplore{}, err
	}

	return exps, nil
}
