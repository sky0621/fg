package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/sky0621/fg/graph/generated"
	"github.com/sky0621/fg/graph/model"
)

func (r *mutationResolver) CreateNegotiation(ctx context.Context, input model.InputNegotiation) (*model.Negotiation, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	n, err := r.negotiationRepository.Create(id.String(), input)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func (r *mutationResolver) UpdateNegotiation(ctx context.Context, id string, input model.InputNegotiation) (*model.Negotiation, error) {
	n, err := r.negotiationRepository.Update(id, input)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func (r *negotiationResolver) Quest(ctx context.Context, obj *model.Negotiation) (*model.Quest, error) {
	n, err := r.negotiationRepository.Get(obj.ID)
	if err != nil {
		return nil, err
	}
	if n == nil || n.Quest == nil {
		return nil, errors.New(fmt.Sprintf("not found: %s", obj.ID))
	}
	q, err := r.questRepository.Get(n.Quest.ID)
	if err != nil {
		return nil, err
	}
	return q, nil
}

func (r *queryResolver) Negotiation(ctx context.Context, id string) (*model.Negotiation, error) {
	return r.negotiationRepository.Get(id)
}

// Negotiation returns generated.NegotiationResolver implementation.
func (r *Resolver) Negotiation() generated.NegotiationResolver { return &negotiationResolver{r} }

type negotiationResolver struct{ *Resolver }
