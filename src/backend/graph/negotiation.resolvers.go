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

func (r *mutationResolver) CreateNegotiation(ctx context.Context, input model.NewNegotiation) (*model.Negotiation, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	q := &model.Negotiation{
		ID:    id.String(),
		Note:  input.Note,
		Done:  input.Done,
		Quest: &model.Quest{ID: input.QuestID},
	}
	if err := r.cache.Add(fmt.Sprintf("negotiation%s", id.String()), q, -1); err != nil {
		return nil, err
	}
	return q, nil
}

func (r *mutationResolver) UpdateNegotiation(ctx context.Context, id string, input model.NewNegotiation) (*model.Negotiation, error) {
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

func (r *negotiationResolver) Quest(ctx context.Context, obj *model.Negotiation) (*model.Quest, error) {
	item, ok := r.cache.Get(fmt.Sprintf("quest%s", obj.Quest.ID))
	if !ok {
		return nil, nil
	}
	return item.(*model.Quest), nil
}

func (r *queryResolver) Negotiations(ctx context.Context) ([]*model.Negotiation, error) {
	results := []*model.Negotiation{}
	for _, item := range r.cache.Items() {
		if v, ok := item.Object.(*model.Negotiation); ok {
			results = append(results, v)
		}
	}
	return results, nil
}

func (r *queryResolver) Negotiation(ctx context.Context, id string) (*model.Negotiation, error) {
	item, ok := r.cache.Get(fmt.Sprintf("negotiation%s", id))
	if !ok {
		return nil, nil
	}
	q := item.(*model.Negotiation)
	return q, nil
}

// Negotiation returns generated.NegotiationResolver implementation.
func (r *Resolver) Negotiation() generated.NegotiationResolver { return &negotiationResolver{r} }

type negotiationResolver struct{ *Resolver }
