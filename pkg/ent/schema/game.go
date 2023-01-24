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
		field.Int64("room_id").StructTag(`json:"room_id"`).Comment("所属房间 ID"),
		field.Enum("end_by").Values("none", "blue", "red", "assassination").Default("none").StructTag(`json:"end_by"`).Comment("游戏结束方式"),
		field.Uint8("capacity").Default(0).StructTag(`json:"capacity"`).Comment("游戏人数"),
		field.Strings("the_assassinated_ids").StructTag(`json:"the_assassinated_ids"`).Comment("被刺杀者[们] ID"),
		field.Uint8("assassin_chance").StructTag(`json:"assassin_chance"`).Default(1).Comment("刺杀机会，默认为 1 次"),
	}
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("game_users", GameUser.Type),
		edge.To("missions", Mission.Type),
		edge.From("room", Room.Type).Ref("games").Unique().Field("room_id").Required(),
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
