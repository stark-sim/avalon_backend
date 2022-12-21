package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/stark-sim/avalon_backend/pkg/ent/roomuser"
	"github.com/stark-sim/avalon_backend/tools"
	"time"

	"github.com/stark-sim/avalon_backend/pkg/ent"
	"github.com/stark-sim/avalon_backend/pkg/graphql/model"
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

// GetRoomUsers is the resolver for the getRoomUsers field.
func (r *subscriptionResolver) GetRoomUsers(ctx context.Context, req *model.RoomRequest) (<-chan []*ent.RoomUser, error) {
	roomID := tools.StringToInt64(req.ID)
	ch := make(chan []*ent.RoomUser)
	go func() {
		for {
			roomUsers, err := r.client.RoomUser.
				Query().
				Where(roomuser.RoomID(roomID), roomuser.DeletedAt(tools.ZeroTime)).
				Order(ent.Asc(roomuser.FieldCreatedAt)).
				All(ctx)
			if err != nil {
				return
			}
			select {
			case ch <- roomUsers:
				time.Sleep(2 * time.Second)
			default:
				return
			}
		}
	}()
	return ch, nil
}

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
