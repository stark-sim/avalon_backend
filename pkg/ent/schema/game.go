package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Game holds the schema definition for the Game entity.
type Game struct {
	ent.Schema
}

// Fields of the Game.
func (Game) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("room_id").StructTag(`json:"room_id"`),
		field.Bool("closed").Default(false).StructTag(`json:"closed"`),
		field.Uint8("capacity").Default(0).StructTag(`json:"capacity"`).Comment("游戏人数"),
	}
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("game_users", GameUser.Type),
		edge.To("game_missions", Mission.Type),
		edge.From("room", Room.Type).Ref("room_games").Unique().Required(),
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
