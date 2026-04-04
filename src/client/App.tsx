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
import {routes} from "routes.generated";

const router = createBrowserRouter([
  {
    path: routes.Home.path,
    element: <IndexPage />
  },
  {
    path: routes.TicTacToeIndex.path,
    element: <TicTacToeIndexPage />
  },
  {
    path: routes.UserProfile.path,
    element: <UserProfilePage />
  },
  {
    path: routes.Login.path,
    element: <LoginPage />
  },
  {
    path: routes.Register.path,
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
