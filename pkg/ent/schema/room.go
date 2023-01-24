package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

// Fields of the Card.
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("").Annotations(entgql.OrderField("NAME")).StructTag(`json:"name"`).Comment("房间名"),
		field.Bool("closed").Default(false).StructTag(`json:"closed"`).Comment("房间是否已关闭"),
		field.Bool("game_on").Default(false).StructTag(`json:"game_on"`).Comment("房间是否有游戏进行中"),
	}
}

// Edges of the Card.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("room_users", RoomUser.Type),
		edge.To("games", Game.Type),
		edge.To("records", Record.Type),
	}
}

func (Room) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Room) Annotations() []schema.Annotation {
	return []schema.Annotation{
		//entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
