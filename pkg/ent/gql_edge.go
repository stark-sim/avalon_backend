// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (c *Card) GameUsers(ctx context.Context) (result []*GameUser, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedGameUsers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.GameUsersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryGameUsers().All(ctx)
	}
	return result, err
}

func (ga *Game) GameUsers(ctx context.Context) (result []*GameUser, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ga.NamedGameUsers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ga.Edges.GameUsersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ga.QueryGameUsers().All(ctx)
	}
	return result, err
}

func (ga *Game) Missions(ctx context.Context) (result []*Mission, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ga.NamedMissions(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ga.Edges.MissionsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ga.QueryMissions().All(ctx)
	}
	return result, err
}

func (ga *Game) Room(ctx context.Context) (*Room, error) {
	result, err := ga.Edges.RoomOrErr()
	if IsNotLoaded(err) {
		result, err = ga.QueryRoom().Only(ctx)
	}
	return result, err
}

func (gu *GameUser) Game(ctx context.Context) (*Game, error) {
	result, err := gu.Edges.GameOrErr()
	if IsNotLoaded(err) {
		result, err = gu.QueryGame().Only(ctx)
	}
	return result, err
}

func (gu *GameUser) Card(ctx context.Context) (*Card, error) {
	result, err := gu.Edges.CardOrErr()
	if IsNotLoaded(err) {
		result, err = gu.QueryCard().Only(ctx)
	}
	return result, err
}

func (m *Mission) Game(ctx context.Context) (*Game, error) {
	result, err := m.Edges.GameOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryGame().Only(ctx)
	}
	return result, err
}

func (m *Mission) Squads(ctx context.Context) (result []*Squad, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = m.NamedSquads(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = m.Edges.SquadsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = m.QuerySquads().All(ctx)
	}
	return result, err
}

func (m *Mission) Votes(ctx context.Context) (result []*Vote, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = m.NamedVotes(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = m.Edges.VotesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = m.QueryVotes().All(ctx)
	}
	return result, err
}

func (r *Record) Room(ctx context.Context) (*Room, error) {
	result, err := r.Edges.RoomOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryRoom().Only(ctx)
	}
	return result, err
}

func (r *Room) RoomUsers(ctx context.Context) (result []*RoomUser, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedRoomUsers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.RoomUsersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryRoomUsers().All(ctx)
	}
	return result, err
}

func (r *Room) Games(ctx context.Context) (result []*Game, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedGames(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.GamesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryGames().All(ctx)
	}
	return result, err
}

func (r *Room) Records(ctx context.Context) (result []*Record, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedRecords(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.RecordsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryRecords().All(ctx)
	}
	return result, err
}

func (ru *RoomUser) Room(ctx context.Context) (*Room, error) {
	result, err := ru.Edges.RoomOrErr()
	if IsNotLoaded(err) {
		result, err = ru.QueryRoom().Only(ctx)
	}
	return result, err
}

func (s *Squad) Mission(ctx context.Context) (*Mission, error) {
	result, err := s.Edges.MissionOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryMission().Only(ctx)
	}
	return result, err
}

func (v *Vote) Mission(ctx context.Context) (*Mission, error) {
	result, err := v.Edges.MissionOrErr()
	if IsNotLoaded(err) {
		result, err = v.QueryMission().Only(ctx)
	}
	return result, err
}
