import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import {
  createRouter,
  RouterProvider,
  createMemoryHistory,
} from "@tanstack/react-router";
import userEvent from "@testing-library/user-event";
import { render, type RenderResult } from "@testing-library/react";

import { routeTree } from "./router";

export const router = createRouter({
  routeTree,
  history: createMemoryHistory(),
});

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: false,
    },
  },
});

export const renderWithProviders = (
  routePath: string,
  initialParams?: Record<string, string>,
): RenderResult => {
  // Navigate to the initial route with params
  router.navigate({
    to: routePath,
    params: initialParams,
  });

  return render(
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>,
  );
};

export { userEvent, queryClient };
