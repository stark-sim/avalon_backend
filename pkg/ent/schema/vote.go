package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Vote holds the schema definition for the Vote entity.
type Vote struct {
	ent.Schema
}

// Fields of the Vote.
func (Vote) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("mission_id").StructTag(`json:"mission_id"`).Comment("所属任务 ID"),
		field.Int64("user_id").StructTag(`json:"user_id"`).Annotations(entgql.Type("ID")).Comment("用户 ID"),
		field.Bool("pass").Default(false).StructTag(`json:"pass"`).Comment("是否赞同目前小队出征"),
		field.Bool("voted").Default(false).StructTag(`json:"voted"`).Comment("是否已投票"),
	}
}

// Edges of the Vote.
func (Vote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("mission", Mission.Type).Ref("votes").Field("mission_id").Unique().Required(),
	}
}

func (Vote) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Vote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		//entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
