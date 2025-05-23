package pokeapi

import (
	"net/http"
	"time"

	"github.com/rpowelson12/pokedex/internal/pokecache"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

//print pokemon info as a test

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
