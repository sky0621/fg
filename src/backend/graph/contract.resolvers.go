package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/fg/graph/generated"
	"github.com/sky0621/fg/graph/model"
)

func (r *contractResolver) Quest(ctx context.Context, obj *model.Contract) (*model.Quest, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateContract(ctx context.Context, input model.InputContract) (*model.Contract, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DiscardContract(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Contract(ctx context.Context, id string) (*model.Contract, error) {
	panic(fmt.Errorf("not implemented"))
}

// Contract returns generated.ContractResolver implementation.
func (r *Resolver) Contract() generated.ContractResolver { return &contractResolver{r} }

type contractResolver struct{ *Resolver }
