package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Mission holds the schema definition for the Mission entity.
type Mission struct {
	ent.Schema
}

// Fields of the Mission.
func (Mission) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("sequence").Max(5).Min(1).StructTag(`json:"sequence"`),
		field.Enum("status").Values("picking", "voting", "acting", "closed", "delayed").Default("picking").StructTag(`json:"status"`),
		field.Bool("failed").Default(false).StructTag(`json:"failed"`),
		field.Int64("game_id").StructTag(`json:"game_id"`),
		field.Uint8("capacity").Default(0).StructTag(`json:"capacity"`).Comment("任务人数"),
		field.Int64("leader").StructTag(`json:"leader"`),
	}
}

// Edges of the Mission.
func (Mission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("game", Game.Type).Ref("missions").Field("game_id").Unique().Required(),
		edge.To("squads", Squad.Type),
		edge.To("votes", Vote.Type),
	}
}

func (Mission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Mission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		//entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
