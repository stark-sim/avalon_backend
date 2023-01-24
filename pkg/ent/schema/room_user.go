package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RoomUser holds the schema definition for the RoomUser entity.
type RoomUser struct {
	ent.Schema
}

// Fields of the Card.
func (RoomUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id"`).Annotations(entgql.Type("ID")).Comment("用户 ID"),
		field.Int64("room_id").StructTag(`json:"room_id"`).Comment("房间 ID"),
		field.Bool("host").StructTag(`json:"host"`).Default(false).Comment("是否为房主"),
	}
}

// Edges of the Card.
func (RoomUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).Ref("room_users").Field("room_id").Unique().Required(),
	}
}

func (RoomUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (RoomUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		//entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
