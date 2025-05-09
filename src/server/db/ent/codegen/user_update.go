// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"puzzlr.gg/src/server/db/ent/codegen/game"
	"puzzlr.gg/src/server/db/ent/codegen/gameplayer"
	"puzzlr.gg/src/server/db/ent/codegen/predicate"
	"puzzlr.gg/src/server/db/ent/codegen/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetHashedPassword sets the "hashed_password" field.
func (uu *UserUpdate) SetHashedPassword(s string) *UserUpdate {
	uu.mutation.SetHashedPassword(s)
	return uu
}

// SetNillableHashedPassword sets the "hashed_password" field if the given value is not nil.
func (uu *UserUpdate) SetNillableHashedPassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetHashedPassword(*s)
	}
	return uu
}

// AddGameIDs adds the "games" edge to the Game entity by IDs.
func (uu *UserUpdate) AddGameIDs(ids ...int) *UserUpdate {
	uu.mutation.AddGameIDs(ids...)
	return uu
}

// AddGames adds the "games" edges to the Game entity.
func (uu *UserUpdate) AddGames(g ...*Game) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.AddGameIDs(ids...)
}

// AddWonGameIDs adds the "won_games" edge to the Game entity by IDs.
func (uu *UserUpdate) AddWonGameIDs(ids ...int) *UserUpdate {
	uu.mutation.AddWonGameIDs(ids...)
	return uu
}

// AddWonGames adds the "won_games" edges to the Game entity.
func (uu *UserUpdate) AddWonGames(g ...*Game) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.AddWonGameIDs(ids...)
}

// AddCurrentTurnGameIDs adds the "current_turn_games" edge to the Game entity by IDs.
func (uu *UserUpdate) AddCurrentTurnGameIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCurrentTurnGameIDs(ids...)
	return uu
}

// AddCurrentTurnGames adds the "current_turn_games" edges to the Game entity.
func (uu *UserUpdate) AddCurrentTurnGames(g ...*Game) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.AddCurrentTurnGameIDs(ids...)
}

// AddGamePlayerIDs adds the "game_player" edge to the GamePlayer entity by IDs.
func (uu *UserUpdate) AddGamePlayerIDs(ids ...int) *UserUpdate {
	uu.mutation.AddGamePlayerIDs(ids...)
	return uu
}

// AddGamePlayer adds the "game_player" edges to the GamePlayer entity.
func (uu *UserUpdate) AddGamePlayer(g ...*GamePlayer) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.AddGamePlayerIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearGames clears all "games" edges to the Game entity.
func (uu *UserUpdate) ClearGames() *UserUpdate {
	uu.mutation.ClearGames()
	return uu
}

// RemoveGameIDs removes the "games" edge to Game entities by IDs.
func (uu *UserUpdate) RemoveGameIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveGameIDs(ids...)
	return uu
}

// RemoveGames removes "games" edges to Game entities.
func (uu *UserUpdate) RemoveGames(g ...*Game) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.RemoveGameIDs(ids...)
}

// ClearWonGames clears all "won_games" edges to the Game entity.
func (uu *UserUpdate) ClearWonGames() *UserUpdate {
	uu.mutation.ClearWonGames()
	return uu
}

// RemoveWonGameIDs removes the "won_games" edge to Game entities by IDs.
func (uu *UserUpdate) RemoveWonGameIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveWonGameIDs(ids...)
	return uu
}

// RemoveWonGames removes "won_games" edges to Game entities.
func (uu *UserUpdate) RemoveWonGames(g ...*Game) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.RemoveWonGameIDs(ids...)
}

// ClearCurrentTurnGames clears all "current_turn_games" edges to the Game entity.
func (uu *UserUpdate) ClearCurrentTurnGames() *UserUpdate {
	uu.mutation.ClearCurrentTurnGames()
	return uu
}

// RemoveCurrentTurnGameIDs removes the "current_turn_games" edge to Game entities by IDs.
func (uu *UserUpdate) RemoveCurrentTurnGameIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCurrentTurnGameIDs(ids...)
	return uu
}

// RemoveCurrentTurnGames removes "current_turn_games" edges to Game entities.
func (uu *UserUpdate) RemoveCurrentTurnGames(g ...*Game) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.RemoveCurrentTurnGameIDs(ids...)
}

// ClearGamePlayer clears all "game_player" edges to the GamePlayer entity.
func (uu *UserUpdate) ClearGamePlayer() *UserUpdate {
	uu.mutation.ClearGamePlayer()
	return uu
}

// RemoveGamePlayerIDs removes the "game_player" edge to GamePlayer entities by IDs.
func (uu *UserUpdate) RemoveGamePlayerIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveGamePlayerIDs(ids...)
	return uu
}

// RemoveGamePlayer removes "game_player" edges to GamePlayer entities.
func (uu *UserUpdate) RemoveGamePlayer(g ...*GamePlayer) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.RemoveGamePlayerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`codegen: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.HashedPassword(); ok {
		if err := user.HashedPasswordValidator(v); err != nil {
			return &ValidationError{Name: "hashed_password", err: fmt.Errorf(`codegen: validator failed for field "User.hashed_password": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.HashedPassword(); ok {
		_spec.SetField(user.FieldHashedPassword, field.TypeString, value)
	}
	if uu.mutation.GamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GamesTable,
			Columns: user.GamesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedGamesIDs(); len(nodes) > 0 && !uu.mutation.GamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GamesTable,
			Columns: user.GamesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GamesTable,
			Columns: user.GamesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.WonGamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WonGamesTable,
			Columns: []string{user.WonGamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedWonGamesIDs(); len(nodes) > 0 && !uu.mutation.WonGamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WonGamesTable,
			Columns: []string{user.WonGamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.WonGamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WonGamesTable,
			Columns: []string{user.WonGamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.CurrentTurnGamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CurrentTurnGamesTable,
			Columns: []string{user.CurrentTurnGamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedCurrentTurnGamesIDs(); len(nodes) > 0 && !uu.mutation.CurrentTurnGamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CurrentTurnGamesTable,
			Columns: []string{user.CurrentTurnGamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CurrentTurnGamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CurrentTurnGamesTable,
			Columns: []string{user.CurrentTurnGamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.GamePlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.GamePlayerTable,
			Columns: []string{user.GamePlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gameplayer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedGamePlayerIDs(); len(nodes) > 0 && !uu.mutation.GamePlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.GamePlayerTable,
			Columns: []string{user.GamePlayerColumn},
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
	if nodes := uu.mutation.GamePlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.GamePlayerTable,
			Columns: []string{user.GamePlayerColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetHashedPassword sets the "hashed_password" field.
func (uuo *UserUpdateOne) SetHashedPassword(s string) *UserUpdateOne {
	uuo.mutation.SetHashedPassword(s)
	return uuo
}

// SetNillableHashedPassword sets the "hashed_password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableHashedPassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetHashedPassword(*s)
	}
	return uuo
}

// AddGameIDs adds the "games" edge to the Game entity by IDs.
func (uuo *UserUpdateOne) AddGameIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddGameIDs(ids...)
	return uuo
}

// AddGames adds the "games" edges to the Game entity.
func (uuo *UserUpdateOne) AddGames(g ...*Game) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.AddGameIDs(ids...)
}

// AddWonGameIDs adds the "won_games" edge to the Game entity by IDs.
func (uuo *UserUpdateOne) AddWonGameIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddWonGameIDs(ids...)
	return uuo
}

// AddWonGames adds the "won_games" edges to the Game entity.
func (uuo *UserUpdateOne) AddWonGames(g ...*Game) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.AddWonGameIDs(ids...)
}

// AddCurrentTurnGameIDs adds the "current_turn_games" edge to the Game entity by IDs.
func (uuo *UserUpdateOne) AddCurrentTurnGameIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCurrentTurnGameIDs(ids...)
	return uuo
}

// AddCurrentTurnGames adds the "current_turn_games" edges to the Game entity.
func (uuo *UserUpdateOne) AddCurrentTurnGames(g ...*Game) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.AddCurrentTurnGameIDs(ids...)
}

// AddGamePlayerIDs adds the "game_player" edge to the GamePlayer entity by IDs.
func (uuo *UserUpdateOne) AddGamePlayerIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddGamePlayerIDs(ids...)
	return uuo
}

// AddGamePlayer adds the "game_player" edges to the GamePlayer entity.
func (uuo *UserUpdateOne) AddGamePlayer(g ...*GamePlayer) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.AddGamePlayerIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearGames clears all "games" edges to the Game entity.
func (uuo *UserUpdateOne) ClearGames() *UserUpdateOne {
	uuo.mutation.ClearGames()
	return uuo
}

// RemoveGameIDs removes the "games" edge to Game entities by IDs.
func (uuo *UserUpdateOne) RemoveGameIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveGameIDs(ids...)
	return uuo
}

// RemoveGames removes "games" edges to Game entities.
func (uuo *UserUpdateOne) RemoveGames(g ...*Game) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.RemoveGameIDs(ids...)
}

// ClearWonGames clears all "won_games" edges to the Game entity.
func (uuo *UserUpdateOne) ClearWonGames() *UserUpdateOne {
	uuo.mutation.ClearWonGames()
	return uuo
}

// RemoveWonGameIDs removes the "won_games" edge to Game entities by IDs.
func (uuo *UserUpdateOne) RemoveWonGameIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveWonGameIDs(ids...)
	return uuo
}

// RemoveWonGames removes "won_games" edges to Game entities.
func (uuo *UserUpdateOne) RemoveWonGames(g ...*Game) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.RemoveWonGameIDs(ids...)
}

// ClearCurrentTurnGames clears all "current_turn_games" edges to the Game entity.
func (uuo *UserUpdateOne) ClearCurrentTurnGames() *UserUpdateOne {
	uuo.mutation.ClearCurrentTurnGames()
	return uuo
}

// RemoveCurrentTurnGameIDs removes the "current_turn_games" edge to Game entities by IDs.
func (uuo *UserUpdateOne) RemoveCurrentTurnGameIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCurrentTurnGameIDs(ids...)
	return uuo
}

// RemoveCurrentTurnGames removes "current_turn_games" edges to Game entities.
func (uuo *UserUpdateOne) RemoveCurrentTurnGames(g ...*Game) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.RemoveCurrentTurnGameIDs(ids...)
}

// ClearGamePlayer clears all "game_player" edges to the GamePlayer entity.
func (uuo *UserUpdateOne) ClearGamePlayer() *UserUpdateOne {
	uuo.mutation.ClearGamePlayer()
	return uuo
}

// RemoveGamePlayerIDs removes the "game_player" edge to GamePlayer entities by IDs.
func (uuo *UserUpdateOne) RemoveGamePlayerIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveGamePlayerIDs(ids...)
	return uuo
}

// RemoveGamePlayer removes "game_player" edges to GamePlayer entities.
func (uuo *UserUpdateOne) RemoveGamePlayer(g ...*GamePlayer) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.RemoveGamePlayerIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`codegen: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.HashedPassword(); ok {
		if err := user.HashedPasswordValidator(v); err != nil {
			return &ValidationError{Name: "hashed_password", err: fmt.Errorf(`codegen: validator failed for field "User.hashed_password": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`codegen: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("codegen: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.HashedPassword(); ok {
		_spec.SetField(user.FieldHashedPassword, field.TypeString, value)
	}
	if uuo.mutation.GamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GamesTable,
			Columns: user.GamesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedGamesIDs(); len(nodes) > 0 && !uuo.mutation.GamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GamesTable,
			Columns: user.GamesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GamesTable,
			Columns: user.GamesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.WonGamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WonGamesTable,
			Columns: []string{user.WonGamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedWonGamesIDs(); len(nodes) > 0 && !uuo.mutation.WonGamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WonGamesTable,
			Columns: []string{user.WonGamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.WonGamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WonGamesTable,
			Columns: []string{user.WonGamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.CurrentTurnGamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CurrentTurnGamesTable,
			Columns: []string{user.CurrentTurnGamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedCurrentTurnGamesIDs(); len(nodes) > 0 && !uuo.mutation.CurrentTurnGamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CurrentTurnGamesTable,
			Columns: []string{user.CurrentTurnGamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CurrentTurnGamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CurrentTurnGamesTable,
			Columns: []string{user.CurrentTurnGamesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.GamePlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.GamePlayerTable,
			Columns: []string{user.GamePlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gameplayer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedGamePlayerIDs(); len(nodes) > 0 && !uuo.mutation.GamePlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.GamePlayerTable,
			Columns: []string{user.GamePlayerColumn},
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
	if nodes := uuo.mutation.GamePlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.GamePlayerTable,
			Columns: []string{user.GamePlayerColumn},
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
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
