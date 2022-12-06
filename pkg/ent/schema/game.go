package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// Game holds the schema definition for the Game entity.
type Game struct {
	ent.Schema
}

// Fields of the Card.
func (Game) Fields() []ent.Field {
	return nil
}

// Edges of the Card.
func (Game) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("game_users", GameUser.Type),
	}
}

func (Game) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Game) Annotations() []schema.Annotation {
	return []schema.Annotation{
		//entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
