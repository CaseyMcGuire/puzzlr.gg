// AUTO-GENERATED — do not edit manually
// Run `go run bin/generate_client_routes.go` to regenerate

export const routes = {
  Home: {
    path: "/",
    build: () => "/",
  },
  TicTacToeIndex: {
    path: "/tictactoe",
    build: () => "/tictactoe",
  },
  UserProfile: {
    path: "/user/:id",
    build: (id: string) => `/user/${id}`,
  },
  Login: {
    path: "/login",
    build: () => "/login",
  },
  Register: {
    path: "/register",
    build: () => "/register",
  },
} as const;
