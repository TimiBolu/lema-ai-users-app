import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { createRouter, RouterProvider } from "@tanstack/react-router";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

import { routeTree } from "./router/index.tsx";
import { QueryErrorBoundary } from "./components/ErrorMgr/QueryErrorBoundary.tsx";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
      retry: 3,
      staleTime: 1000 * 60 * 5, // 5 minutes
    },
  },
});

export const Router = createRouter({ routeTree });

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof Router;
  }
}

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <QueryErrorBoundary>
        <RouterProvider router={Router} />
      </QueryErrorBoundary>
    </QueryClientProvider>
  </StrictMode>,
);
