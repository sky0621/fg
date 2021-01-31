package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/sky0621/fg/graph/generated"
	"github.com/sky0621/fg/graph/model"
)

func (r *mutationResolver) CreateQuest(ctx context.Context, input model.InputQuest) (*model.Quest, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	q, err := r.questRepository.Create(id.String(), input)
	if err != nil {
		return nil, err
	}
	return q, nil
}

func (r *mutationResolver) UpdateQuest(ctx context.Context, id string, input model.InputQuest) (*model.Quest, error) {
	q, err := r.questRepository.Update(id, input)
	if err != nil {
		return nil, err
	}
	return q, nil
}

func (r *queryResolver) Quests(ctx context.Context) ([]*model.Quest, error) {
	return r.questRepository.Find()
}

func (r *queryResolver) Quest(ctx context.Context, id string) (*model.Quest, error) {
	return r.questRepository.Get(id)
}

// Questに紐づくNegotiationを返す
func (r *questResolver) Negotiation(ctx context.Context, obj *model.Quest) (*model.Negotiation, error) {
	if obj == nil || obj.Negotiation == nil {
		return nil, nil
	}
	return r.negotiationRepository.Get(obj.Negotiation.ID)
}

// Questに紐づくContractを返す
func (r *questResolver) Contract(ctx context.Context, obj *model.Quest) (*model.Contract, error) {
	if obj == nil || obj.Contract == nil {
		return nil, nil
	}
	// FIXME:
	return nil, nil
}

// Quest returns generated.QuestResolver implementation.
func (r *Resolver) Quest() generated.QuestResolver { return &questResolver{r} }

type questResolver struct{ *Resolver }
