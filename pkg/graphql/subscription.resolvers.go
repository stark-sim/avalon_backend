package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/stark-sim/avalon_backend/pkg/ent"
	"time"
)

// GetRoomUser is the resolver for the GetRoomUser field.
func (r *subscriptionResolver) GetRoomUser(ctx context.Context) (<-chan *ent.RoomUser, error) {
	ch := make(chan *ent.RoomUser)
	go func() {
		roomUsers, err := r.client.RoomUser.Query().All(ctx)
		if err != nil {
			return
		}
		for _, roomUser := range roomUsers {
			select {
			case ch <- roomUser:
				// pass one
			default:
				return
			}
			time.Sleep(2 * time.Second)
		}
	}()
	return ch, nil
}

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
