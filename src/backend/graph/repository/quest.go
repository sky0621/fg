package repository

import (
	"errors"
	"fmt"

	"github.com/patrickmn/go-cache"
	"github.com/sky0621/fg/graph/model"
)

type QuestRepository interface {
	Find() ([]*model.Quest, error)
	Get(id string) (*model.Quest, error)
	Create(id string, input model.InputQuest) (*model.Quest, error)
	Update(id string, input model.InputQuest) (*model.Quest, error)
}

func NewQuestRepository(cache *cache.Cache) QuestRepository {
	return &questRepository{cache: cache}
}

type questRepository struct {
	cache *cache.Cache
}

func (r *questRepository) Find() ([]*model.Quest, error) {
	results := []*model.Quest{}
	for _, item := range r.cache.Items() {
		if v, ok := item.Object.(*model.Quest); ok {
			results = append(results, v)
		}
	}
	return results, nil
}

func (r *questRepository) Get(id string) (*model.Quest, error) {
	item, ok := r.cache.Get(fmt.Sprintf("quest%s", id))
	if !ok {
		return nil, nil
	}
	q := item.(*model.Quest)
	return q, nil
}

func (r *questRepository) Create(id string, input model.InputQuest) (*model.Quest, error) {
	q := &model.Quest{
		ID:        id,
		Title:     input.Title,
		Text:      input.Text,
		Reward:    input.Reward,
		Incentive: input.Incentive,
	}
	if err := r.cache.Add(fmt.Sprintf("quest%s", id), q, -1); err != nil {
		return nil, err
	}
	return q, nil
}

func (r *questRepository) Update(id string, input model.InputQuest) (*model.Quest, error) {
	m, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	if m == nil {
		return nil, errors.New(fmt.Sprintf("not found: %s", id))
	}

	q := &model.Quest{
		ID:          id,
		Title:       input.Title,
		Text:        input.Text,
		Reward:      input.Reward,
		Incentive:   input.Incentive,
		Negotiation: m.Negotiation,
		Contract:    m.Contract,
	}
	if err := r.cache.Add(fmt.Sprintf("quest%s", id), q, -1); err != nil {
		return nil, err
	}
	return q, nil
}
