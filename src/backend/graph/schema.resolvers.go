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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Negotiation returns generated.NegotiationResolver implementation.
func (r *Resolver) Negotiation() generated.NegotiationResolver { return &negotiationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type negotiationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
