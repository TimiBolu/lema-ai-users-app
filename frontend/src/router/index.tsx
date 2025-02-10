import { Outlet, createRoute, createRootRoute } from "@tanstack/react-router";
import { Toaster } from "react-hot-toast";
import { TanStackRouterDevtools } from "@tanstack/router-devtools";

import "../index.css";
import AppLayout from "../layout";
import UsersTable from "../pages/UsersTable.tsx";
import PostManager from "../pages/PostManager.tsx";
import EnvConfig from "../utils/env.config.ts";

const rootRoute = createRootRoute({
  component: () => (
    <>
      <AppLayout>
        <Toaster
          position="top-right"
          toastOptions={{
            style: {
              background: "#333",
              color: "#fff",
              borderRadius: "8px",
              padding: "12px 16px",
            },
            success: { icon: "✅" },
            error: { icon: "❌" },
          }}
        />
        <Outlet />
      </AppLayout>
      {EnvConfig.apiEnv === "development" && <TanStackRouterDevtools />}
    </>
  ),
});

const indexRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/",
  component: UsersTable,
});

const postsRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/users/$userId/posts",
  component: PostManager,
});

export const routeTree = rootRoute.addChildren([indexRoute, postsRoute]);
