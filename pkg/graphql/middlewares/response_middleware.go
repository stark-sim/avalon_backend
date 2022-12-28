package middlewares

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/stark-sim/avalon_backend/pkg/ent"
	"github.com/stark-sim/avalon_backend/pkg/ent/card"
	"github.com/stark-sim/avalon_backend/tools"
)

type BlueCardVaguer struct{}

func (BlueCardVaguer) ExtensionName() string {
	return "BlueCardVaguer"
}

func (BlueCardVaguer) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func (s BlueCardVaguer) InterceptField(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
	fc := graphql.GetFieldContext(ctx)
	rc := graphql.GetRootFieldContext(ctx)
	if rc != nil && rc.Object == "getVagueGameUsers" {
		if fc.Object == "Card" {
			tempCard := fc.Parent.Result.(*ent.Card)
			if tools.IsOneOf(string(tempCard.Role), string(card.RoleKnight), string(card.RoleLoyal), string(card.RoleAce)) {
				tempCard.Role = card.RoleProphet
				tempCard.Name = card.NameMerlin
				fc.Parent.Result = tempCard
			}
		}
	}
	return next(ctx)
}
