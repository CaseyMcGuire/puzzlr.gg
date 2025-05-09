// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"puzzlr.gg/src/server/db/ent/codegen/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"puzzlr.gg/src/server/db/ent/codegen/game"
	"puzzlr.gg/src/server/db/ent/codegen/gameplayer"
	"puzzlr.gg/src/server/db/ent/codegen/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Game is the client for interacting with the Game builders.
	Game *GameClient
	// GamePlayer is the client for interacting with the GamePlayer builders.
	GamePlayer *GamePlayerClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Game = NewGameClient(c.config)
	c.GamePlayer = NewGamePlayerClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("codegen: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("codegen: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Game:       NewGameClient(cfg),
		GamePlayer: NewGamePlayerClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Game:       NewGameClient(cfg),
		GamePlayer: NewGamePlayerClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Game.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Game.Use(hooks...)
	c.GamePlayer.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Game.Intercept(interceptors...)
	c.GamePlayer.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *GameMutation:
		return c.Game.mutate(ctx, m)
	case *GamePlayerMutation:
		return c.GamePlayer.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("codegen: unknown mutation type %T", m)
	}
}

// GameClient is a client for the Game schema.
type GameClient struct {
	config
}

// NewGameClient returns a client for the Game from the given config.
func NewGameClient(c config) *GameClient {
	return &GameClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `game.Hooks(f(g(h())))`.
func (c *GameClient) Use(hooks ...Hook) {
	c.hooks.Game = append(c.hooks.Game, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `game.Intercept(f(g(h())))`.
func (c *GameClient) Intercept(interceptors ...Interceptor) {
	c.inters.Game = append(c.inters.Game, interceptors...)
}

// Create returns a builder for creating a Game entity.
func (c *GameClient) Create() *GameCreate {
	mutation := newGameMutation(c.config, OpCreate)
	return &GameCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Game entities.
func (c *GameClient) CreateBulk(builders ...*GameCreate) *GameCreateBulk {
	return &GameCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *GameClient) MapCreateBulk(slice any, setFunc func(*GameCreate, int)) *GameCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &GameCreateBulk{err: fmt.Errorf("calling to GameClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*GameCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &GameCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Game.
func (c *GameClient) Update() *GameUpdate {
	mutation := newGameMutation(c.config, OpUpdate)
	return &GameUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GameClient) UpdateOne(ga *Game) *GameUpdateOne {
	mutation := newGameMutation(c.config, OpUpdateOne, withGame(ga))
	return &GameUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GameClient) UpdateOneID(id int) *GameUpdateOne {
	mutation := newGameMutation(c.config, OpUpdateOne, withGameID(id))
	return &GameUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Game.
func (c *GameClient) Delete() *GameDelete {
	mutation := newGameMutation(c.config, OpDelete)
	return &GameDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GameClient) DeleteOne(ga *Game) *GameDeleteOne {
	return c.DeleteOneID(ga.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *GameClient) DeleteOneID(id int) *GameDeleteOne {
	builder := c.Delete().Where(game.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GameDeleteOne{builder}
}

// Query returns a query builder for Game.
func (c *GameClient) Query() *GameQuery {
	return &GameQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeGame},
		inters: c.Interceptors(),
	}
}

// Get returns a Game entity by its id.
func (c *GameClient) Get(ctx context.Context, id int) (*Game, error) {
	return c.Query().Where(game.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GameClient) GetX(ctx context.Context, id int) *Game {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Game.
func (c *GameClient) QueryUser(ga *Game) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ga.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(game.Table, game.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, game.UserTable, game.UserPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ga.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWinner queries the winner edge of a Game.
func (c *GameClient) QueryWinner(ga *Game) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ga.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(game.Table, game.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, game.WinnerTable, game.WinnerColumn),
		)
		fromV = sqlgraph.Neighbors(ga.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCurrentTurn queries the current_turn edge of a Game.
func (c *GameClient) QueryCurrentTurn(ga *Game) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ga.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(game.Table, game.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, game.CurrentTurnTable, game.CurrentTurnColumn),
		)
		fromV = sqlgraph.Neighbors(ga.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGamePlayer queries the game_player edge of a Game.
func (c *GameClient) QueryGamePlayer(ga *Game) *GamePlayerQuery {
	query := (&GamePlayerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ga.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(game.Table, game.FieldID, id),
			sqlgraph.To(gameplayer.Table, gameplayer.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, game.GamePlayerTable, game.GamePlayerColumn),
		)
		fromV = sqlgraph.Neighbors(ga.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GameClient) Hooks() []Hook {
	return c.hooks.Game
}

// Interceptors returns the client interceptors.
func (c *GameClient) Interceptors() []Interceptor {
	return c.inters.Game
}

func (c *GameClient) mutate(ctx context.Context, m *GameMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&GameCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&GameUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&GameUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&GameDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("codegen: unknown Game mutation op: %q", m.Op())
	}
}

// GamePlayerClient is a client for the GamePlayer schema.
type GamePlayerClient struct {
	config
}

// NewGamePlayerClient returns a client for the GamePlayer from the given config.
func NewGamePlayerClient(c config) *GamePlayerClient {
	return &GamePlayerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `gameplayer.Hooks(f(g(h())))`.
func (c *GamePlayerClient) Use(hooks ...Hook) {
	c.hooks.GamePlayer = append(c.hooks.GamePlayer, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `gameplayer.Intercept(f(g(h())))`.
func (c *GamePlayerClient) Intercept(interceptors ...Interceptor) {
	c.inters.GamePlayer = append(c.inters.GamePlayer, interceptors...)
}

// Create returns a builder for creating a GamePlayer entity.
func (c *GamePlayerClient) Create() *GamePlayerCreate {
	mutation := newGamePlayerMutation(c.config, OpCreate)
	return &GamePlayerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GamePlayer entities.
func (c *GamePlayerClient) CreateBulk(builders ...*GamePlayerCreate) *GamePlayerCreateBulk {
	return &GamePlayerCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *GamePlayerClient) MapCreateBulk(slice any, setFunc func(*GamePlayerCreate, int)) *GamePlayerCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &GamePlayerCreateBulk{err: fmt.Errorf("calling to GamePlayerClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*GamePlayerCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &GamePlayerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GamePlayer.
func (c *GamePlayerClient) Update() *GamePlayerUpdate {
	mutation := newGamePlayerMutation(c.config, OpUpdate)
	return &GamePlayerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GamePlayerClient) UpdateOne(gp *GamePlayer) *GamePlayerUpdateOne {
	mutation := newGamePlayerMutation(c.config, OpUpdateOne, withGamePlayer(gp))
	return &GamePlayerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GamePlayerClient) UpdateOneID(id int) *GamePlayerUpdateOne {
	mutation := newGamePlayerMutation(c.config, OpUpdateOne, withGamePlayerID(id))
	return &GamePlayerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GamePlayer.
func (c *GamePlayerClient) Delete() *GamePlayerDelete {
	mutation := newGamePlayerMutation(c.config, OpDelete)
	return &GamePlayerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GamePlayerClient) DeleteOne(gp *GamePlayer) *GamePlayerDeleteOne {
	return c.DeleteOneID(gp.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *GamePlayerClient) DeleteOneID(id int) *GamePlayerDeleteOne {
	builder := c.Delete().Where(gameplayer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GamePlayerDeleteOne{builder}
}

// Query returns a query builder for GamePlayer.
func (c *GamePlayerClient) Query() *GamePlayerQuery {
	return &GamePlayerQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeGamePlayer},
		inters: c.Interceptors(),
	}
}

// Get returns a GamePlayer entity by its id.
func (c *GamePlayerClient) Get(ctx context.Context, id int) (*GamePlayer, error) {
	return c.Query().Where(gameplayer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GamePlayerClient) GetX(ctx context.Context, id int) *GamePlayer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a GamePlayer.
func (c *GamePlayerClient) QueryUser(gp *GamePlayer) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := gp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(gameplayer.Table, gameplayer.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, gameplayer.UserTable, gameplayer.UserColumn),
		)
		fromV = sqlgraph.Neighbors(gp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGame queries the game edge of a GamePlayer.
func (c *GamePlayerClient) QueryGame(gp *GamePlayer) *GameQuery {
	query := (&GameClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := gp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(gameplayer.Table, gameplayer.FieldID, id),
			sqlgraph.To(game.Table, game.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, gameplayer.GameTable, gameplayer.GameColumn),
		)
		fromV = sqlgraph.Neighbors(gp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GamePlayerClient) Hooks() []Hook {
	return c.hooks.GamePlayer
}

// Interceptors returns the client interceptors.
func (c *GamePlayerClient) Interceptors() []Interceptor {
	return c.inters.GamePlayer
}

func (c *GamePlayerClient) mutate(ctx context.Context, m *GamePlayerMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&GamePlayerCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&GamePlayerUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&GamePlayerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&GamePlayerDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("codegen: unknown GamePlayer mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGames queries the games edge of a User.
func (c *UserClient) QueryGames(u *User) *GameQuery {
	query := (&GameClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(game.Table, game.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.GamesTable, user.GamesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWonGames queries the won_games edge of a User.
func (c *UserClient) QueryWonGames(u *User) *GameQuery {
	query := (&GameClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(game.Table, game.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.WonGamesTable, user.WonGamesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCurrentTurnGames queries the current_turn_games edge of a User.
func (c *UserClient) QueryCurrentTurnGames(u *User) *GameQuery {
	query := (&GameClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(game.Table, game.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CurrentTurnGamesTable, user.CurrentTurnGamesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGamePlayer queries the game_player edge of a User.
func (c *UserClient) QueryGamePlayer(u *User) *GamePlayerQuery {
	query := (&GamePlayerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(gameplayer.Table, gameplayer.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.GamePlayerTable, user.GamePlayerColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("codegen: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Game, GamePlayer, User []ent.Hook
	}
	inters struct {
		Game, GamePlayer, User []ent.Interceptor
	}
)
