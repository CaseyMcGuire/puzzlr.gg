package routes

type RouteName string

const (
	Home           RouteName = "Home"
	TicTacToeIndex RouteName = "TicTacToeIndex"
	UserProfile    RouteName = "UserProfile"
	Login          RouteName = "Login"
	Register       RouteName = "Register"
)

type Route struct {
	Path string
	Name RouteName
}

var PageRoutes = map[RouteName]Route{
	Home:           {Path: "/", Name: Home},
	TicTacToeIndex: {Path: "/game/tictactoe", Name: TicTacToeIndex},
	UserProfile:    {Path: "/user/{id}", Name: UserProfile},
	Login:          {Path: "/login", Name: Login},
	Register:       {Path: "/register", Name: Register},
}
