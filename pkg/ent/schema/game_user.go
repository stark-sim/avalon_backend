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
		field.Int64("user_id").StructTag(`json:"user_id"`).Annotations(entgql.Type("ID")).Comment("用户 ID"),
		field.Int64("game_id").StructTag(`json:"game_id"`).Comment("游戏 ID"),
		field.Int64("card_id").StructTag(`json:"card_id"`).Comment("卡牌 ID"),
		field.Uint8("number").StructTag(`json:"number"`).Comment("排序，号数"),
		field.Bool("exited").StructTag(`json:"exited"`).Default(false).Comment("用户是否已经退出游戏，都退出游戏时游戏关闭，为了让游戏结束时用户还可以进去查看结果"),
	}
}

// Edges of the Card.
func (GameUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("game", Game.Type).Ref("game_users").Field("game_id").Unique().Required(),
		edge.From("card", Card.Type).Ref("game_users").Field("card_id").Unique().Required(),
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
