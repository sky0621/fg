package repository

import (
	"fmt"

	"github.com/patrickmn/go-cache"
	"github.com/sky0621/fg/graph/model"
)

type NegotiationRepository interface {
	Get(id string) (*model.Negotiation, error)
	Create(id string, input model.InputNegotiation) (*model.Negotiation, error)
	Update(id string, input model.InputNegotiation) (*model.Negotiation, error)
}

func NewNegotiationRepository(cache *cache.Cache) NegotiationRepository {
	return &negotiationRepository{cache: cache}
}

type negotiationRepository struct {
	cache *cache.Cache
}

func (r *negotiationRepository) Get(id string) (*model.Negotiation, error) {
	item, ok := r.cache.Get(fmt.Sprintf("negotiation%s", id))
	if !ok {
		return nil, nil
	}
	q := item.(*model.Negotiation)
	return q, nil
}

func (r *negotiationRepository) Create(id string, input model.InputNegotiation) (*model.Negotiation, error) {
	q := &model.Negotiation{
		ID:    id,
		Note:  input.Note,
		Done:  false,
		Quest: &model.Quest{ID: input.QuestID},
	}
	if err := r.cache.Add(fmt.Sprintf("negotiation%s", id), q, -1); err != nil {
		return nil, err
	}
	return q, nil
}

func (r *negotiationRepository) Update(id string, input model.InputNegotiation) (*model.Negotiation, error) {
	q := &model.Negotiation{
		ID:    id,
		Note:  input.Note,
		Done:  input.Done,
		Quest: &model.Quest{ID: input.QuestID},
	}
	if err := r.cache.Replace(fmt.Sprintf("negotiation%s", id), q, -1); err != nil {
		return nil, err
	}
	return q, nil
}
