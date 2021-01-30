package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sky0621/fg/graph/model"
)

func (r *mutationResolver) CreateQuest(ctx context.Context, input model.NewQuest) (*model.Quest, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	q := &model.Quest{
		ID:        id.String(),
		Title:     input.Title,
		Text:      input.Text,
		Reward:    input.Reward,
		Incentive: input.Incentive,
	}
	if err := r.cache.Add(fmt.Sprintf("quest%s", id.String()), q, -1); err != nil {
		return nil, err
	}
	return q, nil
}

func (r *mutationResolver) UpdateQuest(ctx context.Context, id string, input model.NewQuest) (*model.Quest, error) {
	q := &model.Quest{
		ID:        id,
		Title:     input.Title,
		Text:      input.Text,
		Reward:    input.Reward,
		Incentive: input.Incentive,
	}
	if err := r.cache.Replace(fmt.Sprintf("quest%s", id), q, -1); err != nil {
		return nil, err
	}
	return q, nil
}

func (r *queryResolver) Quests(ctx context.Context) ([]*model.Quest, error) {
	results := []*model.Quest{}
	for _, item := range r.cache.Items() {
		if v, ok := item.Object.(*model.Quest); ok {
			results = append(results, v)
		}
	}
	return results, nil
}

func (r *queryResolver) Quest(ctx context.Context, id string) (*model.Quest, error) {
	item, ok := r.cache.Get(fmt.Sprintf("quest%s", id))
	if !ok {
		return nil, nil
	}
	q := item.(*model.Quest)
	return q, nil
}
