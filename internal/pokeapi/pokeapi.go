package pokeapi

import (
	"net/http"
	"time"

	"github.com/jampetz/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: 30 * time.Second,
		},
	}
}
