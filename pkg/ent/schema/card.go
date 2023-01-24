package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("name").Values("Merlin", "Percival", "Galahad", "Bors", "Bedivere", "Gawain", "Kay", "Ector", "Mordred", "Morgana", "Oberon", "Agravain", "Lancelot", "Kevin", "Stuart", "Bob").Default("Merlin").StructTag(`json:"name"`).Annotations(entgql.OrderField("NAME")),
		field.Enum("role").Values("Prophet", "Knight", "Loyal", "Usurper", "Enchantress", "Assassin", "Erlking", "Ace", "Sinner", "Minion").StructTag(`json:"role"`),
		field.String("tale").Default("").StructTag(`json:"tale"`),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("game_users", GameUser.Type),
	}
}

func (Card) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Card) Annotations() []schema.Annotation {
	return []schema.Annotation{
		//entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
