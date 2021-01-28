package graph

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	cache *cache.Cache
}

func NewResolver() *Resolver {
	return &Resolver{cache: cache.New(60*time.Second, 60*time.Second)}
}
