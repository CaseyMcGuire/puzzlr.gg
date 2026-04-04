// AUTO-GENERATED — do not edit manually
// Run `go run bin/generate_client_routes.go` to regenerate

export const routes = {
  Login: {
    path: "/login",
    build: () => "/login",
  },
  Register: {
    path: "/register",
    build: () => "/register",
  },
  Home: {
    path: "/",
    build: () => "/",
  },
  TicTacToeIndex: {
    path: "/game/tictactoe",
    build: () => "/game/tictactoe",
  },
  UserProfile: {
    path: "/user/:id",
    build: (id: string) => `/user/${id}`,
  },
} as const;
