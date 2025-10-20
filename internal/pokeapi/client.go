package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/kwschnizel/pokedexgo/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: *pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetBody(url string) ([]byte, error) {
	body, ok := c.cache.Get(url)
	if !ok {
		res, err := c.httpClient.Get(url)
		if err != nil {
			return []byte{}, fmt.Errorf("error from Client.GetBody: error from httpClien.Get: %w", err)
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return []byte{}, fmt.Errorf("error in Client.GetBody: error from io.ReadAll: %w", err)
		}

		if res.StatusCode > 299 {
			return []byte{}, fmt.Errorf("error from Client.GetBody: Response failed with status code %v", res.StatusCode)
		}

		c.cache.Add(url, body)
	}

	return body, nil
}
