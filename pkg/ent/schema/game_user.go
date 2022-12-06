package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GameUser holds the schema definition for the GameUser entity.
type GameUser struct {
	ent.Schema
}

// Fields of the Card.
func (GameUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id"`),
		field.Int64("game_id").StructTag(`json:"game_id"`),
	}
}

// Edges of the Card.
func (GameUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("game", Game.Type).Ref("game_users").Field("game_id").Unique().Required(),
	}
}

func (GameUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (GameUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		//entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
