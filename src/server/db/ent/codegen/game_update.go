// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"puzzlr.gg/src/server/db/ent/codegen/game"
	"puzzlr.gg/src/server/db/ent/codegen/gameplayer"
	"puzzlr.gg/src/server/db/ent/codegen/predicate"
	"puzzlr.gg/src/server/db/ent/codegen/user"
)

// GameUpdate is the builder for updating Game entities.
type GameUpdate struct {
	config
	hooks    []Hook
	mutation *GameMutation
}

// Where appends a list predicates to the GameUpdate builder.
func (gu *GameUpdate) Where(ps ...predicate.Game) *GameUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetUpdateTime sets the "update_time" field.
func (gu *GameUpdate) SetUpdateTime(t time.Time) *GameUpdate {
	gu.mutation.SetUpdateTime(t)
	return gu
}

// SetBoard sets the "board" field.
func (gu *GameUpdate) SetBoard(s [][]string) *GameUpdate {
	gu.mutation.SetBoard(s)
	return gu
}

// AppendBoard appends s to the "board" field.
func (gu *GameUpdate) AppendBoard(s [][]string) *GameUpdate {
	gu.mutation.AppendBoard(s)
	return gu
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (gu *GameUpdate) AddUserIDs(ids ...int) *GameUpdate {
	gu.mutation.AddUserIDs(ids...)
	return gu
}

// AddUser adds the "user" edges to the User entity.
func (gu *GameUpdate) AddUser(u ...*User) *GameUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.AddUserIDs(ids...)
}

// SetWinnerID sets the "winner" edge to the User entity by ID.
func (gu *GameUpdate) SetWinnerID(id int) *GameUpdate {
	gu.mutation.SetWinnerID(id)
	return gu
}

// SetNillableWinnerID sets the "winner" edge to the User entity by ID if the given value is not nil.
func (gu *GameUpdate) SetNillableWinnerID(id *int) *GameUpdate {
	if id != nil {
		gu = gu.SetWinnerID(*id)
	}
	return gu
}

// SetWinner sets the "winner" edge to the User entity.
func (gu *GameUpdate) SetWinner(u *User) *GameUpdate {
	return gu.SetWinnerID(u.ID)
}

// SetCurrentTurnID sets the "current_turn" edge to the User entity by ID.
func (gu *GameUpdate) SetCurrentTurnID(id int) *GameUpdate {
	gu.mutation.SetCurrentTurnID(id)
	return gu
}

// SetNillableCurrentTurnID sets the "current_turn" edge to the User entity by ID if the given value is not nil.
func (gu *GameUpdate) SetNillableCurrentTurnID(id *int) *GameUpdate {
	if id != nil {
		gu = gu.SetCurrentTurnID(*id)
	}
	return gu
}

// SetCurrentTurn sets the "current_turn" edge to the User entity.
func (gu *GameUpdate) SetCurrentTurn(u *User) *GameUpdate {
	return gu.SetCurrentTurnID(u.ID)
}

// AddGamePlayerIDs adds the "game_player" edge to the GamePlayer entity by IDs.
func (gu *GameUpdate) AddGamePlayerIDs(ids ...int) *GameUpdate {
	gu.mutation.AddGamePlayerIDs(ids...)
	return gu
}

// AddGamePlayer adds the "game_player" edges to the GamePlayer entity.
func (gu *GameUpdate) AddGamePlayer(g ...*GamePlayer) *GameUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.AddGamePlayerIDs(ids...)
}

// Mutation returns the GameMutation object of the builder.
func (gu *GameUpdate) Mutation() *GameMutation {
	return gu.mutation
}

// ClearUser clears all "user" edges to the User entity.
func (gu *GameUpdate) ClearUser() *GameUpdate {
	gu.mutation.ClearUser()
	return gu
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (gu *GameUpdate) RemoveUserIDs(ids ...int) *GameUpdate {
	gu.mutation.RemoveUserIDs(ids...)
	return gu
}

// RemoveUser removes "user" edges to User entities.
func (gu *GameUpdate) RemoveUser(u ...*User) *GameUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.RemoveUserIDs(ids...)
}

// ClearWinner clears the "winner" edge to the User entity.
func (gu *GameUpdate) ClearWinner() *GameUpdate {
	gu.mutation.ClearWinner()
	return gu
}

// ClearCurrentTurn clears the "current_turn" edge to the User entity.
func (gu *GameUpdate) ClearCurrentTurn() *GameUpdate {
	gu.mutation.ClearCurrentTurn()
	return gu
}

// ClearGamePlayer clears all "game_player" edges to the GamePlayer entity.
func (gu *GameUpdate) ClearGamePlayer() *GameUpdate {
	gu.mutation.ClearGamePlayer()
	return gu
}

// RemoveGamePlayerIDs removes the "game_player" edge to GamePlayer entities by IDs.
func (gu *GameUpdate) RemoveGamePlayerIDs(ids ...int) *GameUpdate {
	gu.mutation.RemoveGamePlayerIDs(ids...)
	return gu
}

// RemoveGamePlayer removes "game_player" edges to GamePlayer entities.
func (gu *GameUpdate) RemoveGamePlayer(g ...*GamePlayer) *GameUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.RemoveGamePlayerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GameUpdate) Save(ctx context.Context) (int, error) {
	gu.defaults()
	return withHooks(ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GameUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GameUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GameUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GameUpdate) defaults() {
	if _, ok := gu.mutation.UpdateTime(); !ok {
		v := game.UpdateDefaultUpdateTime()
		gu.mutation.SetUpdateTime(v)
	}
}

func (gu *GameUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(game.Table, game.Columns, sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.UpdateTime(); ok {
		_spec.SetField(game.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := gu.mutation.Board(); ok {
		_spec.SetField(game.FieldBoard, field.TypeJSON, value)
	}
	if value, ok := gu.mutation.AppendedBoard(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, game.FieldBoard, value)
		})
	}
	if gu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   game.UserTable,
			Columns: game.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedUserIDs(); len(nodes) > 0 && !gu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   game.UserTable,
			Columns: game.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   game.UserTable,
			Columns: game.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.WinnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.WinnerTable,
			Columns: []string{game.WinnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.WinnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.WinnerTable,
			Columns: []string{game.WinnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.CurrentTurnCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.CurrentTurnTable,
			Columns: []string{game.CurrentTurnColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.CurrentTurnIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.CurrentTurnTable,
			Columns: []string{game.CurrentTurnColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.GamePlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   game.GamePlayerTable,
			Columns: []string{game.GamePlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gameplayer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedGamePlayerIDs(); len(nodes) > 0 && !gu.mutation.GamePlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   game.GamePlayerTable,
			Columns: []string{game.GamePlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gameplayer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.GamePlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   game.GamePlayerTable,
			Columns: []string{game.GamePlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gameplayer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{game.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GameUpdateOne is the builder for updating a single Game entity.
type GameUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GameMutation
}

// SetUpdateTime sets the "update_time" field.
func (guo *GameUpdateOne) SetUpdateTime(t time.Time) *GameUpdateOne {
	guo.mutation.SetUpdateTime(t)
	return guo
}

// SetBoard sets the "board" field.
func (guo *GameUpdateOne) SetBoard(s [][]string) *GameUpdateOne {
	guo.mutation.SetBoard(s)
	return guo
}

// AppendBoard appends s to the "board" field.
func (guo *GameUpdateOne) AppendBoard(s [][]string) *GameUpdateOne {
	guo.mutation.AppendBoard(s)
	return guo
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (guo *GameUpdateOne) AddUserIDs(ids ...int) *GameUpdateOne {
	guo.mutation.AddUserIDs(ids...)
	return guo
}

// AddUser adds the "user" edges to the User entity.
func (guo *GameUpdateOne) AddUser(u ...*User) *GameUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.AddUserIDs(ids...)
}

// SetWinnerID sets the "winner" edge to the User entity by ID.
func (guo *GameUpdateOne) SetWinnerID(id int) *GameUpdateOne {
	guo.mutation.SetWinnerID(id)
	return guo
}

// SetNillableWinnerID sets the "winner" edge to the User entity by ID if the given value is not nil.
func (guo *GameUpdateOne) SetNillableWinnerID(id *int) *GameUpdateOne {
	if id != nil {
		guo = guo.SetWinnerID(*id)
	}
	return guo
}

// SetWinner sets the "winner" edge to the User entity.
func (guo *GameUpdateOne) SetWinner(u *User) *GameUpdateOne {
	return guo.SetWinnerID(u.ID)
}

// SetCurrentTurnID sets the "current_turn" edge to the User entity by ID.
func (guo *GameUpdateOne) SetCurrentTurnID(id int) *GameUpdateOne {
	guo.mutation.SetCurrentTurnID(id)
	return guo
}

// SetNillableCurrentTurnID sets the "current_turn" edge to the User entity by ID if the given value is not nil.
func (guo *GameUpdateOne) SetNillableCurrentTurnID(id *int) *GameUpdateOne {
	if id != nil {
		guo = guo.SetCurrentTurnID(*id)
	}
	return guo
}

// SetCurrentTurn sets the "current_turn" edge to the User entity.
func (guo *GameUpdateOne) SetCurrentTurn(u *User) *GameUpdateOne {
	return guo.SetCurrentTurnID(u.ID)
}

// AddGamePlayerIDs adds the "game_player" edge to the GamePlayer entity by IDs.
func (guo *GameUpdateOne) AddGamePlayerIDs(ids ...int) *GameUpdateOne {
	guo.mutation.AddGamePlayerIDs(ids...)
	return guo
}

// AddGamePlayer adds the "game_player" edges to the GamePlayer entity.
func (guo *GameUpdateOne) AddGamePlayer(g ...*GamePlayer) *GameUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.AddGamePlayerIDs(ids...)
}

// Mutation returns the GameMutation object of the builder.
func (guo *GameUpdateOne) Mutation() *GameMutation {
	return guo.mutation
}

// ClearUser clears all "user" edges to the User entity.
func (guo *GameUpdateOne) ClearUser() *GameUpdateOne {
	guo.mutation.ClearUser()
	return guo
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (guo *GameUpdateOne) RemoveUserIDs(ids ...int) *GameUpdateOne {
	guo.mutation.RemoveUserIDs(ids...)
	return guo
}

// RemoveUser removes "user" edges to User entities.
func (guo *GameUpdateOne) RemoveUser(u ...*User) *GameUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.RemoveUserIDs(ids...)
}

// ClearWinner clears the "winner" edge to the User entity.
func (guo *GameUpdateOne) ClearWinner() *GameUpdateOne {
	guo.mutation.ClearWinner()
	return guo
}

// ClearCurrentTurn clears the "current_turn" edge to the User entity.
func (guo *GameUpdateOne) ClearCurrentTurn() *GameUpdateOne {
	guo.mutation.ClearCurrentTurn()
	return guo
}

// ClearGamePlayer clears all "game_player" edges to the GamePlayer entity.
func (guo *GameUpdateOne) ClearGamePlayer() *GameUpdateOne {
	guo.mutation.ClearGamePlayer()
	return guo
}

// RemoveGamePlayerIDs removes the "game_player" edge to GamePlayer entities by IDs.
func (guo *GameUpdateOne) RemoveGamePlayerIDs(ids ...int) *GameUpdateOne {
	guo.mutation.RemoveGamePlayerIDs(ids...)
	return guo
}

// RemoveGamePlayer removes "game_player" edges to GamePlayer entities.
func (guo *GameUpdateOne) RemoveGamePlayer(g ...*GamePlayer) *GameUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.RemoveGamePlayerIDs(ids...)
}

// Where appends a list predicates to the GameUpdate builder.
func (guo *GameUpdateOne) Where(ps ...predicate.Game) *GameUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GameUpdateOne) Select(field string, fields ...string) *GameUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Game entity.
func (guo *GameUpdateOne) Save(ctx context.Context) (*Game, error) {
	guo.defaults()
	return withHooks(ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GameUpdateOne) SaveX(ctx context.Context) *Game {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GameUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GameUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GameUpdateOne) defaults() {
	if _, ok := guo.mutation.UpdateTime(); !ok {
		v := game.UpdateDefaultUpdateTime()
		guo.mutation.SetUpdateTime(v)
	}
}

func (guo *GameUpdateOne) sqlSave(ctx context.Context) (_node *Game, err error) {
	_spec := sqlgraph.NewUpdateSpec(game.Table, game.Columns, sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`codegen: missing "Game.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, game.FieldID)
		for _, f := range fields {
			if !game.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("codegen: invalid field %q for query", f)}
			}
			if f != game.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.UpdateTime(); ok {
		_spec.SetField(game.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := guo.mutation.Board(); ok {
		_spec.SetField(game.FieldBoard, field.TypeJSON, value)
	}
	if value, ok := guo.mutation.AppendedBoard(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, game.FieldBoard, value)
		})
	}
	if guo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   game.UserTable,
			Columns: game.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedUserIDs(); len(nodes) > 0 && !guo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   game.UserTable,
			Columns: game.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   game.UserTable,
			Columns: game.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.WinnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.WinnerTable,
			Columns: []string{game.WinnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.WinnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.WinnerTable,
			Columns: []string{game.WinnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.CurrentTurnCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.CurrentTurnTable,
			Columns: []string{game.CurrentTurnColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.CurrentTurnIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   game.CurrentTurnTable,
			Columns: []string{game.CurrentTurnColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.GamePlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   game.GamePlayerTable,
			Columns: []string{game.GamePlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gameplayer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedGamePlayerIDs(); len(nodes) > 0 && !guo.mutation.GamePlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   game.GamePlayerTable,
			Columns: []string{game.GamePlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gameplayer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.GamePlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   game.GamePlayerTable,
			Columns: []string{game.GamePlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gameplayer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Game{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{game.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}
