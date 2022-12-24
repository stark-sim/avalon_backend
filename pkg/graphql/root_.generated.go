// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/stark-sim/avalon_backend/pkg/ent"
	"github.com/stark-sim/avalon_backend/pkg/graphql/model"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Card() CardResolver
	Entity() EntityResolver
	Game() GameResolver
	GameUser() GameUserResolver
	Mission() MissionResolver
	Mutation() MutationResolver
	Query() QueryResolver
	Record() RecordResolver
	Room() RoomResolver
	RoomUser() RoomUserResolver
	Squad() SquadResolver
	Subscription() SubscriptionResolver
	Vote() VoteResolver
	CardWhereInput() CardWhereInputResolver
	CreateCardInput() CreateCardInputResolver
	CreateGameInput() CreateGameInputResolver
	CreateGameUserInput() CreateGameUserInputResolver
	CreateMissionInput() CreateMissionInputResolver
	CreateRecordInput() CreateRecordInputResolver
	CreateRoomInput() CreateRoomInputResolver
	CreateRoomUserInput() CreateRoomUserInputResolver
	CreateSquadInput() CreateSquadInputResolver
	CreateVoteInput() CreateVoteInputResolver
	GameUserWhereInput() GameUserWhereInputResolver
	GameWhereInput() GameWhereInputResolver
	MissionWhereInput() MissionWhereInputResolver
	RecordWhereInput() RecordWhereInputResolver
	RoomUserWhereInput() RoomUserWhereInputResolver
	RoomWhereInput() RoomWhereInputResolver
	SquadWhereInput() SquadWhereInputResolver
	UpdateCardInput() UpdateCardInputResolver
	UpdateGameInput() UpdateGameInputResolver
	UpdateGameUserInput() UpdateGameUserInputResolver
	UpdateMissionInput() UpdateMissionInputResolver
	UpdateRecordInput() UpdateRecordInputResolver
	UpdateRoomInput() UpdateRoomInputResolver
	UpdateRoomUserInput() UpdateRoomUserInputResolver
	UpdateSquadInput() UpdateSquadInputResolver
	UpdateVoteInput() UpdateVoteInputResolver
	VoteWhereInput() VoteWhereInputResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Card struct {
		CreatedAt func(childComplexity int) int
		CreatedBy func(childComplexity int) int
		DeletedAt func(childComplexity int) int
		GameUsers func(childComplexity int) int
		ID        func(childComplexity int) int
		Name      func(childComplexity int) int
		Role      func(childComplexity int) int
		Tale      func(childComplexity int) int
		UpdatedAt func(childComplexity int) int
		UpdatedBy func(childComplexity int) int
	}

	Entity struct {
		FindUserByID func(childComplexity int, id string) int
	}

	Game struct {
		Capacity  func(childComplexity int) int
		CreatedAt func(childComplexity int) int
		CreatedBy func(childComplexity int) int
		DeletedAt func(childComplexity int) int
		EndBy     func(childComplexity int) int
		GameUsers func(childComplexity int) int
		ID        func(childComplexity int) int
		Missions  func(childComplexity int) int
		Room      func(childComplexity int) int
		RoomID    func(childComplexity int) int
		UpdatedAt func(childComplexity int) int
		UpdatedBy func(childComplexity int) int
	}

	GameUser struct {
		Card      func(childComplexity int) int
		CardID    func(childComplexity int) int
		CreatedAt func(childComplexity int) int
		CreatedBy func(childComplexity int) int
		DeletedAt func(childComplexity int) int
		Game      func(childComplexity int) int
		GameID    func(childComplexity int) int
		ID        func(childComplexity int) int
		Number    func(childComplexity int) int
		UpdatedAt func(childComplexity int) int
		UpdatedBy func(childComplexity int) int
		User      func(childComplexity int) int
		UserID    func(childComplexity int) int
	}

	Mission struct {
		Capacity  func(childComplexity int) int
		CreatedAt func(childComplexity int) int
		CreatedBy func(childComplexity int) int
		DeletedAt func(childComplexity int) int
		Failed    func(childComplexity int) int
		Game      func(childComplexity int) int
		GameID    func(childComplexity int) int
		ID        func(childComplexity int) int
		Leader    func(childComplexity int) int
		Sequence  func(childComplexity int) int
		Squads    func(childComplexity int) int
		Status    func(childComplexity int) int
		UpdatedAt func(childComplexity int) int
		UpdatedBy func(childComplexity int) int
		Votes     func(childComplexity int) int
	}

	Mutation struct {
		CloseRoom           func(childComplexity int, req model.RoomRequest) int
		CreateCard          func(childComplexity int, req ent.CreateCardInput) int
		CreateGame          func(childComplexity int, req model.RoomRequest) int
		CreateRoom          func(childComplexity int, req ent.CreateRoomInput) int
		JoinRoom            func(childComplexity int, req ent.CreateRoomUserInput) int
		JoinRoomByShortCode func(childComplexity int, req model.JoinRoomInput) int
		LeaveRoom           func(childComplexity int, req ent.CreateRoomUserInput) int
	}

	PageInfo struct {
		EndCursor       func(childComplexity int) int
		HasNextPage     func(childComplexity int) int
		HasPreviousPage func(childComplexity int) int
		StartCursor     func(childComplexity int) int
	}

	Query struct {
		Cards              func(childComplexity int) int
		GameUsers          func(childComplexity int) int
		Games              func(childComplexity int) int
		Missions           func(childComplexity int) int
		Node               func(childComplexity int, id string) int
		Nodes              func(childComplexity int, ids []string) int
		Records            func(childComplexity int) int
		RoomUsers          func(childComplexity int) int
		Rooms              func(childComplexity int) int
		Squads             func(childComplexity int) int
		Votes              func(childComplexity int) int
		__resolve__service func(childComplexity int) int
		__resolve_entities func(childComplexity int, representations []map[string]interface{}) int
	}

	Record struct {
		CreatedAt func(childComplexity int) int
		CreatedBy func(childComplexity int) int
		DeletedAt func(childComplexity int) int
		ID        func(childComplexity int) int
		Room      func(childComplexity int) int
		RoomID    func(childComplexity int) int
		Score     func(childComplexity int) int
		UpdatedAt func(childComplexity int) int
		UpdatedBy func(childComplexity int) int
		User      func(childComplexity int) int
		UserID    func(childComplexity int) int
	}

	Room struct {
		Closed    func(childComplexity int) int
		CreatedAt func(childComplexity int) int
		CreatedBy func(childComplexity int) int
		DeletedAt func(childComplexity int) int
		GameOn    func(childComplexity int) int
		Games     func(childComplexity int) int
		ID        func(childComplexity int) int
		Name      func(childComplexity int) int
		Records   func(childComplexity int) int
		RoomUsers func(childComplexity int) int
		UpdatedAt func(childComplexity int) int
		UpdatedBy func(childComplexity int) int
	}

	RoomUser struct {
		CreatedAt func(childComplexity int) int
		CreatedBy func(childComplexity int) int
		DeletedAt func(childComplexity int) int
		ID        func(childComplexity int) int
		Room      func(childComplexity int) int
		RoomID    func(childComplexity int) int
		UpdatedAt func(childComplexity int) int
		UpdatedBy func(childComplexity int) int
		User      func(childComplexity int) int
		UserID    func(childComplexity int) int
	}

	Squad struct {
		CreatedAt func(childComplexity int) int
		CreatedBy func(childComplexity int) int
		DeletedAt func(childComplexity int) int
		ID        func(childComplexity int) int
		Mission   func(childComplexity int) int
		MissionID func(childComplexity int) int
		Rat       func(childComplexity int) int
		UpdatedAt func(childComplexity int) int
		UpdatedBy func(childComplexity int) int
		User      func(childComplexity int) int
		UserID    func(childComplexity int) int
	}

	Subscription struct {
		GetMissionsByGame  func(childComplexity int, req model.GameRequest) int
		GetRoomOngoingGame func(childComplexity int, req model.RoomRequest) int
		GetRoomUser        func(childComplexity int) int
		GetRoomUsers       func(childComplexity int, req *model.RoomRequest) int
	}

	User struct {
		ID    func(childComplexity int) int
		Name  func(childComplexity int) int
		Phone func(childComplexity int) int
	}

	Vote struct {
		CreatedAt func(childComplexity int) int
		CreatedBy func(childComplexity int) int
		DeletedAt func(childComplexity int) int
		ID        func(childComplexity int) int
		Mission   func(childComplexity int) int
		MissionID func(childComplexity int) int
		Pass      func(childComplexity int) int
		UpdatedAt func(childComplexity int) int
		UpdatedBy func(childComplexity int) int
		User      func(childComplexity int) int
		UserID    func(childComplexity int) int
	}

	_Service struct {
		SDL func(childComplexity int) int
	}
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "Card.createdAt":
		if e.complexity.Card.CreatedAt == nil {
			break
		}

		return e.complexity.Card.CreatedAt(childComplexity), true

	case "Card.createdBy":
		if e.complexity.Card.CreatedBy == nil {
			break
		}

		return e.complexity.Card.CreatedBy(childComplexity), true

	case "Card.deletedAt":
		if e.complexity.Card.DeletedAt == nil {
			break
		}

		return e.complexity.Card.DeletedAt(childComplexity), true

	case "Card.gameUsers":
		if e.complexity.Card.GameUsers == nil {
			break
		}

		return e.complexity.Card.GameUsers(childComplexity), true

	case "Card.id":
		if e.complexity.Card.ID == nil {
			break
		}

		return e.complexity.Card.ID(childComplexity), true

	case "Card.name":
		if e.complexity.Card.Name == nil {
			break
		}

		return e.complexity.Card.Name(childComplexity), true

	case "Card.role":
		if e.complexity.Card.Role == nil {
			break
		}

		return e.complexity.Card.Role(childComplexity), true

	case "Card.tale":
		if e.complexity.Card.Tale == nil {
			break
		}

		return e.complexity.Card.Tale(childComplexity), true

	case "Card.updatedAt":
		if e.complexity.Card.UpdatedAt == nil {
			break
		}

		return e.complexity.Card.UpdatedAt(childComplexity), true

	case "Card.updatedBy":
		if e.complexity.Card.UpdatedBy == nil {
			break
		}

		return e.complexity.Card.UpdatedBy(childComplexity), true

	case "Entity.findUserByID":
		if e.complexity.Entity.FindUserByID == nil {
			break
		}

		args, err := ec.field_Entity_findUserByID_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Entity.FindUserByID(childComplexity, args["id"].(string)), true

	case "Game.capacity":
		if e.complexity.Game.Capacity == nil {
			break
		}

		return e.complexity.Game.Capacity(childComplexity), true

	case "Game.createdAt":
		if e.complexity.Game.CreatedAt == nil {
			break
		}

		return e.complexity.Game.CreatedAt(childComplexity), true

	case "Game.createdBy":
		if e.complexity.Game.CreatedBy == nil {
			break
		}

		return e.complexity.Game.CreatedBy(childComplexity), true

	case "Game.deletedAt":
		if e.complexity.Game.DeletedAt == nil {
			break
		}

		return e.complexity.Game.DeletedAt(childComplexity), true

	case "Game.endBy":
		if e.complexity.Game.EndBy == nil {
			break
		}

		return e.complexity.Game.EndBy(childComplexity), true

	case "Game.gameUsers":
		if e.complexity.Game.GameUsers == nil {
			break
		}

		return e.complexity.Game.GameUsers(childComplexity), true

	case "Game.id":
		if e.complexity.Game.ID == nil {
			break
		}

		return e.complexity.Game.ID(childComplexity), true

	case "Game.missions":
		if e.complexity.Game.Missions == nil {
			break
		}

		return e.complexity.Game.Missions(childComplexity), true

	case "Game.room":
		if e.complexity.Game.Room == nil {
			break
		}

		return e.complexity.Game.Room(childComplexity), true

	case "Game.roomID":
		if e.complexity.Game.RoomID == nil {
			break
		}

		return e.complexity.Game.RoomID(childComplexity), true

	case "Game.updatedAt":
		if e.complexity.Game.UpdatedAt == nil {
			break
		}

		return e.complexity.Game.UpdatedAt(childComplexity), true

	case "Game.updatedBy":
		if e.complexity.Game.UpdatedBy == nil {
			break
		}

		return e.complexity.Game.UpdatedBy(childComplexity), true

	case "GameUser.card":
		if e.complexity.GameUser.Card == nil {
			break
		}

		return e.complexity.GameUser.Card(childComplexity), true

	case "GameUser.cardID":
		if e.complexity.GameUser.CardID == nil {
			break
		}

		return e.complexity.GameUser.CardID(childComplexity), true

	case "GameUser.createdAt":
		if e.complexity.GameUser.CreatedAt == nil {
			break
		}

		return e.complexity.GameUser.CreatedAt(childComplexity), true

	case "GameUser.createdBy":
		if e.complexity.GameUser.CreatedBy == nil {
			break
		}

		return e.complexity.GameUser.CreatedBy(childComplexity), true

	case "GameUser.deletedAt":
		if e.complexity.GameUser.DeletedAt == nil {
			break
		}

		return e.complexity.GameUser.DeletedAt(childComplexity), true

	case "GameUser.game":
		if e.complexity.GameUser.Game == nil {
			break
		}

		return e.complexity.GameUser.Game(childComplexity), true

	case "GameUser.gameID":
		if e.complexity.GameUser.GameID == nil {
			break
		}

		return e.complexity.GameUser.GameID(childComplexity), true

	case "GameUser.id":
		if e.complexity.GameUser.ID == nil {
			break
		}

		return e.complexity.GameUser.ID(childComplexity), true

	case "GameUser.number":
		if e.complexity.GameUser.Number == nil {
			break
		}

		return e.complexity.GameUser.Number(childComplexity), true

	case "GameUser.updatedAt":
		if e.complexity.GameUser.UpdatedAt == nil {
			break
		}

		return e.complexity.GameUser.UpdatedAt(childComplexity), true

	case "GameUser.updatedBy":
		if e.complexity.GameUser.UpdatedBy == nil {
			break
		}

		return e.complexity.GameUser.UpdatedBy(childComplexity), true

	case "GameUser.user":
		if e.complexity.GameUser.User == nil {
			break
		}

		return e.complexity.GameUser.User(childComplexity), true

	case "GameUser.userID":
		if e.complexity.GameUser.UserID == nil {
			break
		}

		return e.complexity.GameUser.UserID(childComplexity), true

	case "Mission.capacity":
		if e.complexity.Mission.Capacity == nil {
			break
		}

		return e.complexity.Mission.Capacity(childComplexity), true

	case "Mission.createdAt":
		if e.complexity.Mission.CreatedAt == nil {
			break
		}

		return e.complexity.Mission.CreatedAt(childComplexity), true

	case "Mission.createdBy":
		if e.complexity.Mission.CreatedBy == nil {
			break
		}

		return e.complexity.Mission.CreatedBy(childComplexity), true

	case "Mission.deletedAt":
		if e.complexity.Mission.DeletedAt == nil {
			break
		}

		return e.complexity.Mission.DeletedAt(childComplexity), true

	case "Mission.failed":
		if e.complexity.Mission.Failed == nil {
			break
		}

		return e.complexity.Mission.Failed(childComplexity), true

	case "Mission.game":
		if e.complexity.Mission.Game == nil {
			break
		}

		return e.complexity.Mission.Game(childComplexity), true

	case "Mission.gameID":
		if e.complexity.Mission.GameID == nil {
			break
		}

		return e.complexity.Mission.GameID(childComplexity), true

	case "Mission.id":
		if e.complexity.Mission.ID == nil {
			break
		}

		return e.complexity.Mission.ID(childComplexity), true

	case "Mission.leader":
		if e.complexity.Mission.Leader == nil {
			break
		}

		return e.complexity.Mission.Leader(childComplexity), true

	case "Mission.sequence":
		if e.complexity.Mission.Sequence == nil {
			break
		}

		return e.complexity.Mission.Sequence(childComplexity), true

	case "Mission.squads":
		if e.complexity.Mission.Squads == nil {
			break
		}

		return e.complexity.Mission.Squads(childComplexity), true

	case "Mission.status":
		if e.complexity.Mission.Status == nil {
			break
		}

		return e.complexity.Mission.Status(childComplexity), true

	case "Mission.updatedAt":
		if e.complexity.Mission.UpdatedAt == nil {
			break
		}

		return e.complexity.Mission.UpdatedAt(childComplexity), true

	case "Mission.updatedBy":
		if e.complexity.Mission.UpdatedBy == nil {
			break
		}

		return e.complexity.Mission.UpdatedBy(childComplexity), true

	case "Mission.votes":
		if e.complexity.Mission.Votes == nil {
			break
		}

		return e.complexity.Mission.Votes(childComplexity), true

	case "Mutation.closeRoom":
		if e.complexity.Mutation.CloseRoom == nil {
			break
		}

		args, err := ec.field_Mutation_closeRoom_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CloseRoom(childComplexity, args["req"].(model.RoomRequest)), true

	case "Mutation.createCard":
		if e.complexity.Mutation.CreateCard == nil {
			break
		}

		args, err := ec.field_Mutation_createCard_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateCard(childComplexity, args["req"].(ent.CreateCardInput)), true

	case "Mutation.createGame":
		if e.complexity.Mutation.CreateGame == nil {
			break
		}

		args, err := ec.field_Mutation_createGame_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateGame(childComplexity, args["req"].(model.RoomRequest)), true

	case "Mutation.createRoom":
		if e.complexity.Mutation.CreateRoom == nil {
			break
		}

		args, err := ec.field_Mutation_createRoom_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateRoom(childComplexity, args["req"].(ent.CreateRoomInput)), true

	case "Mutation.joinRoom":
		if e.complexity.Mutation.JoinRoom == nil {
			break
		}

		args, err := ec.field_Mutation_joinRoom_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.JoinRoom(childComplexity, args["req"].(ent.CreateRoomUserInput)), true

	case "Mutation.joinRoomByShortCode":
		if e.complexity.Mutation.JoinRoomByShortCode == nil {
			break
		}

		args, err := ec.field_Mutation_joinRoomByShortCode_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.JoinRoomByShortCode(childComplexity, args["req"].(model.JoinRoomInput)), true

	case "Mutation.leaveRoom":
		if e.complexity.Mutation.LeaveRoom == nil {
			break
		}

		args, err := ec.field_Mutation_leaveRoom_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.LeaveRoom(childComplexity, args["req"].(ent.CreateRoomUserInput)), true

	case "PageInfo.endCursor":
		if e.complexity.PageInfo.EndCursor == nil {
			break
		}

		return e.complexity.PageInfo.EndCursor(childComplexity), true

	case "PageInfo.hasNextPage":
		if e.complexity.PageInfo.HasNextPage == nil {
			break
		}

		return e.complexity.PageInfo.HasNextPage(childComplexity), true

	case "PageInfo.hasPreviousPage":
		if e.complexity.PageInfo.HasPreviousPage == nil {
			break
		}

		return e.complexity.PageInfo.HasPreviousPage(childComplexity), true

	case "PageInfo.startCursor":
		if e.complexity.PageInfo.StartCursor == nil {
			break
		}

		return e.complexity.PageInfo.StartCursor(childComplexity), true

	case "Query.cards":
		if e.complexity.Query.Cards == nil {
			break
		}

		return e.complexity.Query.Cards(childComplexity), true

	case "Query.gameUsers":
		if e.complexity.Query.GameUsers == nil {
			break
		}

		return e.complexity.Query.GameUsers(childComplexity), true

	case "Query.games":
		if e.complexity.Query.Games == nil {
			break
		}

		return e.complexity.Query.Games(childComplexity), true

	case "Query.missions":
		if e.complexity.Query.Missions == nil {
			break
		}

		return e.complexity.Query.Missions(childComplexity), true

	case "Query.node":
		if e.complexity.Query.Node == nil {
			break
		}

		args, err := ec.field_Query_node_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Node(childComplexity, args["id"].(string)), true

	case "Query.nodes":
		if e.complexity.Query.Nodes == nil {
			break
		}

		args, err := ec.field_Query_nodes_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Nodes(childComplexity, args["ids"].([]string)), true

	case "Query.records":
		if e.complexity.Query.Records == nil {
			break
		}

		return e.complexity.Query.Records(childComplexity), true

	case "Query.roomUsers":
		if e.complexity.Query.RoomUsers == nil {
			break
		}

		return e.complexity.Query.RoomUsers(childComplexity), true

	case "Query.rooms":
		if e.complexity.Query.Rooms == nil {
			break
		}

		return e.complexity.Query.Rooms(childComplexity), true

	case "Query.squads":
		if e.complexity.Query.Squads == nil {
			break
		}

		return e.complexity.Query.Squads(childComplexity), true

	case "Query.votes":
		if e.complexity.Query.Votes == nil {
			break
		}

		return e.complexity.Query.Votes(childComplexity), true

	case "Query._service":
		if e.complexity.Query.__resolve__service == nil {
			break
		}

		return e.complexity.Query.__resolve__service(childComplexity), true

	case "Query._entities":
		if e.complexity.Query.__resolve_entities == nil {
			break
		}

		args, err := ec.field_Query__entities_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.__resolve_entities(childComplexity, args["representations"].([]map[string]interface{})), true

	case "Record.createdAt":
		if e.complexity.Record.CreatedAt == nil {
			break
		}

		return e.complexity.Record.CreatedAt(childComplexity), true

	case "Record.createdBy":
		if e.complexity.Record.CreatedBy == nil {
			break
		}

		return e.complexity.Record.CreatedBy(childComplexity), true

	case "Record.deletedAt":
		if e.complexity.Record.DeletedAt == nil {
			break
		}

		return e.complexity.Record.DeletedAt(childComplexity), true

	case "Record.id":
		if e.complexity.Record.ID == nil {
			break
		}

		return e.complexity.Record.ID(childComplexity), true

	case "Record.room":
		if e.complexity.Record.Room == nil {
			break
		}

		return e.complexity.Record.Room(childComplexity), true

	case "Record.roomID":
		if e.complexity.Record.RoomID == nil {
			break
		}

		return e.complexity.Record.RoomID(childComplexity), true

	case "Record.score":
		if e.complexity.Record.Score == nil {
			break
		}

		return e.complexity.Record.Score(childComplexity), true

	case "Record.updatedAt":
		if e.complexity.Record.UpdatedAt == nil {
			break
		}

		return e.complexity.Record.UpdatedAt(childComplexity), true

	case "Record.updatedBy":
		if e.complexity.Record.UpdatedBy == nil {
			break
		}

		return e.complexity.Record.UpdatedBy(childComplexity), true

	case "Record.user":
		if e.complexity.Record.User == nil {
			break
		}

		return e.complexity.Record.User(childComplexity), true

	case "Record.userID":
		if e.complexity.Record.UserID == nil {
			break
		}

		return e.complexity.Record.UserID(childComplexity), true

	case "Room.closed":
		if e.complexity.Room.Closed == nil {
			break
		}

		return e.complexity.Room.Closed(childComplexity), true

	case "Room.createdAt":
		if e.complexity.Room.CreatedAt == nil {
			break
		}

		return e.complexity.Room.CreatedAt(childComplexity), true

	case "Room.createdBy":
		if e.complexity.Room.CreatedBy == nil {
			break
		}

		return e.complexity.Room.CreatedBy(childComplexity), true

	case "Room.deletedAt":
		if e.complexity.Room.DeletedAt == nil {
			break
		}

		return e.complexity.Room.DeletedAt(childComplexity), true

	case "Room.gameOn":
		if e.complexity.Room.GameOn == nil {
			break
		}

		return e.complexity.Room.GameOn(childComplexity), true

	case "Room.games":
		if e.complexity.Room.Games == nil {
			break
		}

		return e.complexity.Room.Games(childComplexity), true

	case "Room.id":
		if e.complexity.Room.ID == nil {
			break
		}

		return e.complexity.Room.ID(childComplexity), true

	case "Room.name":
		if e.complexity.Room.Name == nil {
			break
		}

		return e.complexity.Room.Name(childComplexity), true

	case "Room.records":
		if e.complexity.Room.Records == nil {
			break
		}

		return e.complexity.Room.Records(childComplexity), true

	case "Room.roomUsers":
		if e.complexity.Room.RoomUsers == nil {
			break
		}

		return e.complexity.Room.RoomUsers(childComplexity), true

	case "Room.updatedAt":
		if e.complexity.Room.UpdatedAt == nil {
			break
		}

		return e.complexity.Room.UpdatedAt(childComplexity), true

	case "Room.updatedBy":
		if e.complexity.Room.UpdatedBy == nil {
			break
		}

		return e.complexity.Room.UpdatedBy(childComplexity), true

	case "RoomUser.createdAt":
		if e.complexity.RoomUser.CreatedAt == nil {
			break
		}

		return e.complexity.RoomUser.CreatedAt(childComplexity), true

	case "RoomUser.createdBy":
		if e.complexity.RoomUser.CreatedBy == nil {
			break
		}

		return e.complexity.RoomUser.CreatedBy(childComplexity), true

	case "RoomUser.deletedAt":
		if e.complexity.RoomUser.DeletedAt == nil {
			break
		}

		return e.complexity.RoomUser.DeletedAt(childComplexity), true

	case "RoomUser.id":
		if e.complexity.RoomUser.ID == nil {
			break
		}

		return e.complexity.RoomUser.ID(childComplexity), true

	case "RoomUser.room":
		if e.complexity.RoomUser.Room == nil {
			break
		}

		return e.complexity.RoomUser.Room(childComplexity), true

	case "RoomUser.roomID":
		if e.complexity.RoomUser.RoomID == nil {
			break
		}

		return e.complexity.RoomUser.RoomID(childComplexity), true

	case "RoomUser.updatedAt":
		if e.complexity.RoomUser.UpdatedAt == nil {
			break
		}

		return e.complexity.RoomUser.UpdatedAt(childComplexity), true

	case "RoomUser.updatedBy":
		if e.complexity.RoomUser.UpdatedBy == nil {
			break
		}

		return e.complexity.RoomUser.UpdatedBy(childComplexity), true

	case "RoomUser.user":
		if e.complexity.RoomUser.User == nil {
			break
		}

		return e.complexity.RoomUser.User(childComplexity), true

	case "RoomUser.userID":
		if e.complexity.RoomUser.UserID == nil {
			break
		}

		return e.complexity.RoomUser.UserID(childComplexity), true

	case "Squad.createdAt":
		if e.complexity.Squad.CreatedAt == nil {
			break
		}

		return e.complexity.Squad.CreatedAt(childComplexity), true

	case "Squad.createdBy":
		if e.complexity.Squad.CreatedBy == nil {
			break
		}

		return e.complexity.Squad.CreatedBy(childComplexity), true

	case "Squad.deletedAt":
		if e.complexity.Squad.DeletedAt == nil {
			break
		}

		return e.complexity.Squad.DeletedAt(childComplexity), true

	case "Squad.id":
		if e.complexity.Squad.ID == nil {
			break
		}

		return e.complexity.Squad.ID(childComplexity), true

	case "Squad.mission":
		if e.complexity.Squad.Mission == nil {
			break
		}

		return e.complexity.Squad.Mission(childComplexity), true

	case "Squad.missionID":
		if e.complexity.Squad.MissionID == nil {
			break
		}

		return e.complexity.Squad.MissionID(childComplexity), true

	case "Squad.rat":
		if e.complexity.Squad.Rat == nil {
			break
		}

		return e.complexity.Squad.Rat(childComplexity), true

	case "Squad.updatedAt":
		if e.complexity.Squad.UpdatedAt == nil {
			break
		}

		return e.complexity.Squad.UpdatedAt(childComplexity), true

	case "Squad.updatedBy":
		if e.complexity.Squad.UpdatedBy == nil {
			break
		}

		return e.complexity.Squad.UpdatedBy(childComplexity), true

	case "Squad.user":
		if e.complexity.Squad.User == nil {
			break
		}

		return e.complexity.Squad.User(childComplexity), true

	case "Squad.userID":
		if e.complexity.Squad.UserID == nil {
			break
		}

		return e.complexity.Squad.UserID(childComplexity), true

	case "Subscription.getMissionsByGame":
		if e.complexity.Subscription.GetMissionsByGame == nil {
			break
		}

		args, err := ec.field_Subscription_getMissionsByGame_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Subscription.GetMissionsByGame(childComplexity, args["req"].(model.GameRequest)), true

	case "Subscription.getRoomOngoingGame":
		if e.complexity.Subscription.GetRoomOngoingGame == nil {
			break
		}

		args, err := ec.field_Subscription_getRoomOngoingGame_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Subscription.GetRoomOngoingGame(childComplexity, args["req"].(model.RoomRequest)), true

	case "Subscription.GetRoomUser":
		if e.complexity.Subscription.GetRoomUser == nil {
			break
		}

		return e.complexity.Subscription.GetRoomUser(childComplexity), true

	case "Subscription.getRoomUsers":
		if e.complexity.Subscription.GetRoomUsers == nil {
			break
		}

		args, err := ec.field_Subscription_getRoomUsers_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Subscription.GetRoomUsers(childComplexity, args["req"].(*model.RoomRequest)), true

	case "User.id":
		if e.complexity.User.ID == nil {
			break
		}

		return e.complexity.User.ID(childComplexity), true

	case "User.name":
		if e.complexity.User.Name == nil {
			break
		}

		return e.complexity.User.Name(childComplexity), true

	case "User.phone":
		if e.complexity.User.Phone == nil {
			break
		}

		return e.complexity.User.Phone(childComplexity), true

	case "Vote.createdAt":
		if e.complexity.Vote.CreatedAt == nil {
			break
		}

		return e.complexity.Vote.CreatedAt(childComplexity), true

	case "Vote.createdBy":
		if e.complexity.Vote.CreatedBy == nil {
			break
		}

		return e.complexity.Vote.CreatedBy(childComplexity), true

	case "Vote.deletedAt":
		if e.complexity.Vote.DeletedAt == nil {
			break
		}

		return e.complexity.Vote.DeletedAt(childComplexity), true

	case "Vote.id":
		if e.complexity.Vote.ID == nil {
			break
		}

		return e.complexity.Vote.ID(childComplexity), true

	case "Vote.mission":
		if e.complexity.Vote.Mission == nil {
			break
		}

		return e.complexity.Vote.Mission(childComplexity), true

	case "Vote.missionID":
		if e.complexity.Vote.MissionID == nil {
			break
		}

		return e.complexity.Vote.MissionID(childComplexity), true

	case "Vote.pass":
		if e.complexity.Vote.Pass == nil {
			break
		}

		return e.complexity.Vote.Pass(childComplexity), true

	case "Vote.updatedAt":
		if e.complexity.Vote.UpdatedAt == nil {
			break
		}

		return e.complexity.Vote.UpdatedAt(childComplexity), true

	case "Vote.updatedBy":
		if e.complexity.Vote.UpdatedBy == nil {
			break
		}

		return e.complexity.Vote.UpdatedBy(childComplexity), true

	case "Vote.user":
		if e.complexity.Vote.User == nil {
			break
		}

		return e.complexity.Vote.User(childComplexity), true

	case "Vote.userID":
		if e.complexity.Vote.UserID == nil {
			break
		}

		return e.complexity.Vote.UserID(childComplexity), true

	case "_Service.sdl":
		if e.complexity._Service.SDL == nil {
			break
		}

		return e.complexity._Service.SDL(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputCardOrder,
		ec.unmarshalInputCardWhereInput,
		ec.unmarshalInputCreateCardInput,
		ec.unmarshalInputCreateGameInput,
		ec.unmarshalInputCreateGameUserInput,
		ec.unmarshalInputCreateMissionInput,
		ec.unmarshalInputCreateRecordInput,
		ec.unmarshalInputCreateRoomInput,
		ec.unmarshalInputCreateRoomUserInput,
		ec.unmarshalInputCreateSquadInput,
		ec.unmarshalInputCreateVoteInput,
		ec.unmarshalInputGameOrder,
		ec.unmarshalInputGameRequest,
		ec.unmarshalInputGameUserOrder,
		ec.unmarshalInputGameUserWhereInput,
		ec.unmarshalInputGameWhereInput,
		ec.unmarshalInputJoinRoomInput,
		ec.unmarshalInputMissionOrder,
		ec.unmarshalInputMissionWhereInput,
		ec.unmarshalInputRecordOrder,
		ec.unmarshalInputRecordWhereInput,
		ec.unmarshalInputRoomOrder,
		ec.unmarshalInputRoomRequest,
		ec.unmarshalInputRoomUserOrder,
		ec.unmarshalInputRoomUserWhereInput,
		ec.unmarshalInputRoomWhereInput,
		ec.unmarshalInputSquadOrder,
		ec.unmarshalInputSquadWhereInput,
		ec.unmarshalInputUpdateCardInput,
		ec.unmarshalInputUpdateGameInput,
		ec.unmarshalInputUpdateGameUserInput,
		ec.unmarshalInputUpdateMissionInput,
		ec.unmarshalInputUpdateRecordInput,
		ec.unmarshalInputUpdateRoomInput,
		ec.unmarshalInputUpdateRoomUserInput,
		ec.unmarshalInputUpdateSquadInput,
		ec.unmarshalInputUpdateVoteInput,
		ec.unmarshalInputVoteOrder,
		ec.unmarshalInputVoteWhereInput,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Query(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}
	case ast.Subscription:
		next := ec._Subscription(ctx, rc.Operation.SelectionSet)

		var buf bytes.Buffer
		return func(ctx context.Context) *graphql.Response {
			buf.Reset()
			data := next(ctx)

			if data == nil {
				return nil
			}
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

//go:embed "avalon_backend.graphql" "subscription.graphql"
var sourcesFS embed.FS

func sourceData(filename string) string {
	data, err := sourcesFS.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("codegen problem: %s not available", filename))
	}
	return string(data)
}

var sources = []*ast.Source{
	{Name: "avalon_backend.graphql", Input: sourceData("avalon_backend.graphql"), BuiltIn: false},
	{Name: "subscription.graphql", Input: sourceData("subscription.graphql"), BuiltIn: false},
	{Name: "federation/directives.graphql", Input: `
	scalar _Any
	scalar _FieldSet

	directive @external on FIELD_DEFINITION
	directive @requires(fields: _FieldSet!) on FIELD_DEFINITION
	directive @provides(fields: _FieldSet!) on FIELD_DEFINITION
	directive @extends on OBJECT | INTERFACE

	directive @key(fields: _FieldSet!, resolvable: Boolean = true) repeatable on OBJECT | INTERFACE
	directive @link(import: [String!], url: String!) repeatable on SCHEMA
	directive @shareable on OBJECT | FIELD_DEFINITION
	directive @tag(name: String!) repeatable on FIELD_DEFINITION | INTERFACE | OBJECT | UNION | ARGUMENT_DEFINITION | SCALAR | ENUM | ENUM_VALUE | INPUT_OBJECT | INPUT_FIELD_DEFINITION
	directive @override(from: String!) on FIELD_DEFINITION
	directive @inaccessible on SCALAR | OBJECT | FIELD_DEFINITION | ARGUMENT_DEFINITION | INTERFACE | UNION | ENUM | ENUM_VALUE | INPUT_OBJECT | INPUT_FIELD_DEFINITION
`, BuiltIn: true},
	{Name: "federation/entity.graphql", Input: `
# a union of all types that use the @key directive
union _Entity = User

# fake type to build resolver interfaces for users to implement
type Entity {
		findUserByID(id: ID!,): User!

}

type _Service {
  sdl: String
}

extend type Query {
  _entities(representations: [_Any!]!): [_Entity]!
  _service: _Service!
}
`, BuiltIn: true},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
