package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
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

func (c *Client) GetLocAreas(sectionUrl *string) (LocAreas, error) {
	url := apiUrlBase + "/location-area"
	if sectionUrl != nil {
		url = *sectionUrl
	}

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

	locs := LocAreas{}
	err = json.Unmarshal(body, &locs)
	if err != nil {
		return LocAreas{}, err
	}

	return locs, nil
}
