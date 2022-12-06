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
	Game() GameResolver
	GameUser() GameUserResolver
	Query() QueryResolver
	Room() RoomResolver
	RoomUser() RoomUserResolver
	CardWhereInput() CardWhereInputResolver
	CreateGameInput() CreateGameInputResolver
	CreateGameUserInput() CreateGameUserInputResolver
	CreateRoomInput() CreateRoomInputResolver
	CreateRoomUserInput() CreateRoomUserInputResolver
	GameUserWhereInput() GameUserWhereInputResolver
	GameWhereInput() GameWhereInputResolver
	RoomUserWhereInput() RoomUserWhereInputResolver
	RoomWhereInput() RoomWhereInputResolver
	UpdateGameInput() UpdateGameInputResolver
	UpdateGameUserInput() UpdateGameUserInputResolver
	UpdateRoomInput() UpdateRoomInputResolver
	UpdateRoomUserInput() UpdateRoomUserInputResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Card struct {
		CreatedAt func(childComplexity int) int
		CreatedBy func(childComplexity int) int
		DeletedAt func(childComplexity int) int
		ID        func(childComplexity int) int
		Name      func(childComplexity int) int
		UpdatedAt func(childComplexity int) int
		UpdatedBy func(childComplexity int) int
	}

	Game struct {
		CreatedAt func(childComplexity int) int
		CreatedBy func(childComplexity int) int
		DeletedAt func(childComplexity int) int
		GameUsers func(childComplexity int) int
		ID        func(childComplexity int) int
		UpdatedAt func(childComplexity int) int
		UpdatedBy func(childComplexity int) int
	}

	GameUser struct {
		CreatedAt func(childComplexity int) int
		CreatedBy func(childComplexity int) int
		DeletedAt func(childComplexity int) int
		Game      func(childComplexity int) int
		GameID    func(childComplexity int) int
		ID        func(childComplexity int) int
		UpdatedAt func(childComplexity int) int
		UpdatedBy func(childComplexity int) int
		UserID    func(childComplexity int) int
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
		Node               func(childComplexity int, id string) int
		Nodes              func(childComplexity int, ids []string) int
		RoomUsers          func(childComplexity int) int
		Rooms              func(childComplexity int) int
		__resolve__service func(childComplexity int) int
	}

	Room struct {
		CreatedAt func(childComplexity int) int
		CreatedBy func(childComplexity int) int
		DeletedAt func(childComplexity int) int
		ID        func(childComplexity int) int
		Name      func(childComplexity int) int
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

	case "GameUser.userID":
		if e.complexity.GameUser.UserID == nil {
			break
		}

		return e.complexity.GameUser.UserID(childComplexity), true

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

	case "Query._service":
		if e.complexity.Query.__resolve__service == nil {
			break
		}

		return e.complexity.Query.__resolve__service(childComplexity), true

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

	case "RoomUser.userID":
		if e.complexity.RoomUser.UserID == nil {
			break
		}

		return e.complexity.RoomUser.UserID(childComplexity), true

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
		ec.unmarshalInputCreateRoomInput,
		ec.unmarshalInputCreateRoomUserInput,
		ec.unmarshalInputGameOrder,
		ec.unmarshalInputGameUserOrder,
		ec.unmarshalInputGameUserWhereInput,
		ec.unmarshalInputGameWhereInput,
		ec.unmarshalInputRoomOrder,
		ec.unmarshalInputRoomUserOrder,
		ec.unmarshalInputRoomUserWhereInput,
		ec.unmarshalInputRoomWhereInput,
		ec.unmarshalInputUpdateCardInput,
		ec.unmarshalInputUpdateGameInput,
		ec.unmarshalInputUpdateGameUserInput,
		ec.unmarshalInputUpdateRoomInput,
		ec.unmarshalInputUpdateRoomUserInput,
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

//go:embed "avalon_backend.graphql"
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
type _Service {
  sdl: String
}

extend type Query {
  _service: _Service!
}
`, BuiltIn: true},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
