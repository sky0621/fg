package graph

import (
	"time"

	"github.com/sky0621/fg/graph/repository"

	"github.com/patrickmn/go-cache"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	questRepository       repository.QuestRepository
	negotiationRepository repository.NegotiationRepository
}

func NewResolver() *Resolver {
	c := cache.New(60*time.Second, 60*time.Second)
	return &Resolver{
		questRepository:       repository.NewQuestRepository(c),
		negotiationRepository: repository.NewNegotiationRepository(c),
	}
}
