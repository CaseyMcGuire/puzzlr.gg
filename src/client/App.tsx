import {renderComponent} from "util/ReactPageUtils";
import {createBrowserRouter, RouterProvider} from "react-router";
import LoginPage from "pages/LoginPage/LoginPage";
import RegisterPage from "pages/RegisterPage/RegisterPage";
import {RelayEnvironmentProvider} from "react-relay";
import {RelayConfig} from "relay/RelayConfig";
import {Suspense} from "react";
import IndexPage from "pages/IndexPage/IndexPage";
import TicTacToeIndexPage from "pages/TicTacToeIndexPage/TicTacToeIndexPage";
import UserProfilePage from "pages/UserProfilePage/UserProfilePage";

const router = createBrowserRouter([
  {
    path: '/',
    element: <IndexPage />
  },
  {
    path: '/tictactoe',
    element: <TicTacToeIndexPage />
  },
  {
    path: '/user/:id',
    element: <UserProfilePage />
  },
  {
    path: '/login',
    element: <LoginPage />
  },
  {
    path: '/register',
    element: <RegisterPage />
  }
])

export default function App() {
  return (
    <RelayEnvironmentProvider environment={RelayConfig.getEnvironment()} >
      <Suspense fallback={null}>
        <RouterProvider router={router} />
      </Suspense>
    </RelayEnvironmentProvider>
  )
}

renderComponent(<App />);
