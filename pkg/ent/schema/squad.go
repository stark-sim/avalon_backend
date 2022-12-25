package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Squad 小队，任务和出行人员中间表
// Squad holds the schema definition for the Squad entity.
type Squad struct {
	ent.Schema
}

// Fields of the Squad.
func (Squad) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("mission_id").StructTag(`json:"mission_id"`),
		field.Int64("user_id").StructTag(`json:"user_id"`).Annotations(entgql.Type("ID")),
		field.Bool("rat").Default(false).StructTag(`json:"rat"`).Comment("是否破坏任务"),
		field.Bool("acted").Default(false).StructTag(`json:"acted"`).Comment("是否已经行动"),
	}
}

// Edges of the Squad.
func (Squad) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("mission", Mission.Type).Ref("squads").Field("mission_id").Unique().Required(),
	}
}

func (Squad) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Squad) Annotations() []schema.Annotation {
	return []schema.Annotation{
		//entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
