// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CardsColumns holds the columns for the "cards" table.
	CardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeEnum, Enums: []string{"Merlin", "Percival", "Galahad", "Bors", "Bedivere", "Gawain", "Kay", "Ector", "Mordred", "Morgana", "Oberon", "Agravain", "Lancelot", "Kevin", "Stuart", "Bob"}, Default: "Merlin"},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"Prophet", "Knight", "Loyal", "Usurper", "Enchantress", "Assassin", "Erlking", "Ace", "Sinner", "Minion"}},
		{Name: "tale", Type: field.TypeString, Default: ""},
		{Name: "red", Type: field.TypeBool, Default: false},
	}
	// CardsTable holds the schema information for the "cards" table.
	CardsTable = &schema.Table{
		Name:       "cards",
		Columns:    CardsColumns,
		PrimaryKey: []*schema.Column{CardsColumns[0]},
	}
	// GamesColumns holds the columns for the "games" table.
	GamesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "end_by", Type: field.TypeEnum, Enums: []string{"none", "blue", "red", "assassination"}, Default: "none"},
		{Name: "capacity", Type: field.TypeUint8, Default: 0},
		{Name: "the_assassinated_ids", Type: field.TypeJSON, Nullable: true},
		{Name: "assassin_chance", Type: field.TypeUint8, Default: 1},
		{Name: "room_id", Type: field.TypeInt64},
	}
	// GamesTable holds the schema information for the "games" table.
	GamesTable = &schema.Table{
		Name:       "games",
		Columns:    GamesColumns,
		PrimaryKey: []*schema.Column{GamesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "games_rooms_games",
				Columns:    []*schema.Column{GamesColumns[10]},
				RefColumns: []*schema.Column{RoomsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GameUsersColumns holds the columns for the "game_users" table.
	GameUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "number", Type: field.TypeUint8},
		{Name: "card_id", Type: field.TypeInt64},
		{Name: "game_id", Type: field.TypeInt64},
	}
	// GameUsersTable holds the schema information for the "game_users" table.
	GameUsersTable = &schema.Table{
		Name:       "game_users",
		Columns:    GameUsersColumns,
		PrimaryKey: []*schema.Column{GameUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "game_users_cards_game_users",
				Columns:    []*schema.Column{GameUsersColumns[8]},
				RefColumns: []*schema.Column{CardsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "game_users_games_game_users",
				Columns:    []*schema.Column{GameUsersColumns[9]},
				RefColumns: []*schema.Column{GamesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MissionsColumns holds the columns for the "missions" table.
	MissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "sequence", Type: field.TypeUint8},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"picking", "voting", "acting", "closed", "delayed"}, Default: "picking"},
		{Name: "failed", Type: field.TypeBool, Default: false},
		{Name: "capacity", Type: field.TypeUint8, Default: 0},
		{Name: "leader_id", Type: field.TypeInt64, Default: 0},
		{Name: "protected", Type: field.TypeBool, Default: false},
		{Name: "game_id", Type: field.TypeInt64},
	}
	// MissionsTable holds the schema information for the "missions" table.
	MissionsTable = &schema.Table{
		Name:       "missions",
		Columns:    MissionsColumns,
		PrimaryKey: []*schema.Column{MissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "missions_games_missions",
				Columns:    []*schema.Column{MissionsColumns[12]},
				RefColumns: []*schema.Column{GamesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// RecordsColumns holds the columns for the "records" table.
	RecordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "score", Type: field.TypeInt32, Default: 0},
		{Name: "room_id", Type: field.TypeInt64},
	}
	// RecordsTable holds the schema information for the "records" table.
	RecordsTable = &schema.Table{
		Name:       "records",
		Columns:    RecordsColumns,
		PrimaryKey: []*schema.Column{RecordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "records_rooms_records",
				Columns:    []*schema.Column{RecordsColumns[8]},
				RefColumns: []*schema.Column{RoomsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// RoomsColumns holds the columns for the "rooms" table.
	RoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Default: ""},
		{Name: "closed", Type: field.TypeBool, Default: false},
		{Name: "game_on", Type: field.TypeBool, Default: false},
	}
	// RoomsTable holds the schema information for the "rooms" table.
	RoomsTable = &schema.Table{
		Name:       "rooms",
		Columns:    RoomsColumns,
		PrimaryKey: []*schema.Column{RoomsColumns[0]},
	}
	// RoomUsersColumns holds the columns for the "room_users" table.
	RoomUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "host", Type: field.TypeBool, Default: false},
		{Name: "room_id", Type: field.TypeInt64},
	}
	// RoomUsersTable holds the schema information for the "room_users" table.
	RoomUsersTable = &schema.Table{
		Name:       "room_users",
		Columns:    RoomUsersColumns,
		PrimaryKey: []*schema.Column{RoomUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "room_users_rooms_room_users",
				Columns:    []*schema.Column{RoomUsersColumns[8]},
				RefColumns: []*schema.Column{RoomsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SquadsColumns holds the columns for the "squads" table.
	SquadsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "rat", Type: field.TypeBool, Default: false},
		{Name: "acted", Type: field.TypeBool, Default: false},
		{Name: "mission_id", Type: field.TypeInt64},
	}
	// SquadsTable holds the schema information for the "squads" table.
	SquadsTable = &schema.Table{
		Name:       "squads",
		Columns:    SquadsColumns,
		PrimaryKey: []*schema.Column{SquadsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "squads_missions_squads",
				Columns:    []*schema.Column{SquadsColumns[9]},
				RefColumns: []*schema.Column{MissionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// VotesColumns holds the columns for the "votes" table.
	VotesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "pass", Type: field.TypeBool, Default: false},
		{Name: "voted", Type: field.TypeBool, Default: false},
		{Name: "mission_id", Type: field.TypeInt64},
	}
	// VotesTable holds the schema information for the "votes" table.
	VotesTable = &schema.Table{
		Name:       "votes",
		Columns:    VotesColumns,
		PrimaryKey: []*schema.Column{VotesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "votes_missions_votes",
				Columns:    []*schema.Column{VotesColumns[9]},
				RefColumns: []*schema.Column{MissionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CardsTable,
		GamesTable,
		GameUsersTable,
		MissionsTable,
		RecordsTable,
		RoomsTable,
		RoomUsersTable,
		SquadsTable,
		VotesTable,
	}
)

func init() {
	GamesTable.ForeignKeys[0].RefTable = RoomsTable
	GameUsersTable.ForeignKeys[0].RefTable = CardsTable
	GameUsersTable.ForeignKeys[1].RefTable = GamesTable
	MissionsTable.ForeignKeys[0].RefTable = GamesTable
	RecordsTable.ForeignKeys[0].RefTable = RoomsTable
	RoomUsersTable.ForeignKeys[0].RefTable = RoomsTable
	SquadsTable.ForeignKeys[0].RefTable = MissionsTable
	VotesTable.ForeignKeys[0].RefTable = MissionsTable
}
