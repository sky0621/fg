package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sky0621/fg/graph/generated"
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
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateNegotiation(ctx context.Context, input model.NewNegotiation) (*model.Negotiation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateNegotiation(ctx context.Context, id string, input model.NewNegotiation) (*model.Negotiation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Quests(ctx context.Context) ([]*model.Quest, error) {
	results := []*model.Quest{}
	for _, item := range r.cache.Items() {
		results = append(results, item.Object.(*model.Quest))
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

func (r *queryResolver) Negotiations(ctx context.Context) ([]*model.Negotiation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Negotiation(ctx context.Context, id string) (*model.Negotiation, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
