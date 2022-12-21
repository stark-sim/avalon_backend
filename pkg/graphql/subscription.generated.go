// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stark-sim/avalon_backend/pkg/ent"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type SubscriptionResolver interface {
	GetRoomUser(ctx context.Context) (<-chan *ent.RoomUser, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Subscription_GetRoomUser(ctx context.Context, field graphql.CollectedField) (ret func(ctx context.Context) graphql.Marshaler) {
	fc, err := ec.fieldContext_Subscription_GetRoomUser(ctx, field)
	if err != nil {
		return nil
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = nil
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Subscription().GetRoomUser(rctx)
	})
	if err != nil {
		ec.Error(ctx, err)
		return nil
	}
	if resTmp == nil {
		return nil
	}
	return func(ctx context.Context) graphql.Marshaler {
		select {
		case res, ok := <-resTmp.(<-chan *ent.RoomUser):
			if !ok {
				return nil
			}
			return graphql.WriterFunc(func(w io.Writer) {
				w.Write([]byte{'{'})
				graphql.MarshalString(field.Alias).MarshalGQL(w)
				w.Write([]byte{':'})
				ec.marshalORoomUser2ᚖgithubᚗcomᚋstarkᚑsimᚋavalon_backendᚋpkgᚋentᚐRoomUser(ctx, field.Selections, res).MarshalGQL(w)
				w.Write([]byte{'}'})
			})
		case <-ctx.Done():
			return nil
		}
	}
}

func (ec *executionContext) fieldContext_Subscription_GetRoomUser(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Subscription",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_RoomUser_id(ctx, field)
			case "createdBy":
				return ec.fieldContext_RoomUser_createdBy(ctx, field)
			case "updatedBy":
				return ec.fieldContext_RoomUser_updatedBy(ctx, field)
			case "createdAt":
				return ec.fieldContext_RoomUser_createdAt(ctx, field)
			case "updatedAt":
				return ec.fieldContext_RoomUser_updatedAt(ctx, field)
			case "deletedAt":
				return ec.fieldContext_RoomUser_deletedAt(ctx, field)
			case "userID":
				return ec.fieldContext_RoomUser_userID(ctx, field)
			case "roomID":
				return ec.fieldContext_RoomUser_roomID(ctx, field)
			case "room":
				return ec.fieldContext_RoomUser_room(ctx, field)
			case "user":
				return ec.fieldContext_RoomUser_user(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type RoomUser", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var subscriptionImplementors = []string{"Subscription"}

func (ec *executionContext) _Subscription(ctx context.Context, sel ast.SelectionSet) func(ctx context.Context) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, subscriptionImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Subscription",
	})
	if len(fields) != 1 {
		ec.Errorf(ctx, "must subscribe to exactly one stream")
		return nil
	}

	switch fields[0].Name {
	case "GetRoomUser":
		return ec._Subscription_GetRoomUser(ctx, fields[0])
	default:
		panic("unknown field " + strconv.Quote(fields[0].Name))
	}
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

// endregion ***************************** type.gotpl *****************************