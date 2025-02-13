import {
  describe,
  it,
  expect,
  beforeEach,
  beforeAll,
  afterEach,
  afterAll,
} from "vitest";
import {
  renderWithProviders,
  userEvent,
  queryClient,
  router,
} from "../test-utils";

import { server } from "../mocks/server";
import { http, HttpResponse } from "msw";
import EnvConfig from "../utils/env.config";
import { MockData } from "../mocks/mock.data";
import MockAPIResponse from "../mocks/mock.api.response";

beforeAll(() => {
  MockData.seed();
  EnvConfig.apiBaseUrl = "/api";
  server.listen();
});
afterEach(() => {
  server.resetHandlers();
  queryClient.clear();
});
afterAll(() => server.close());

describe("UsersTable", () => {
  beforeEach(() => {
    server.use(
      http.get("/api/users", ({ request }) => {
        const url = new URL(request.url);
        const pageNumber = url.searchParams.get("pageNumber");
        const pageSize = url.searchParams.get("pageSize");

        const data = {
          users: Array.from({ length: Number(pageSize) || 10 }, (_, i) => ({
            id: String(i + 1 + (Number(pageNumber) - 1) * 10),
            firstname: `User${i + 1}`,
            lastname: "Doe",
            email: `user${i + 1}@example.com`,
            address: {
              street: `${i + 1} Main St`,
              city: "Anytown",
              state: "CA",
              zipCode: "12345",
            },
          })),
          pagination: {
            currentPage: Number(pageNumber) || 1,
            pageSize: Number(pageSize) || 10,
            totalPages: 6,
            totalItems: 15,
            hasNext: Number(pageNumber) < 2,
            hasPrev: Number(pageNumber) > 1,
          },
        };
        const response = MockAPIResponse({ data });
        return HttpResponse.json(response);
      }),
    );
  });

  it("should render user table with pagination", async () => {
    const { findByText, getByRole } = renderWithProviders("/");

    await findByText("User1 Doe");
    expect(!!getByRole("button", { name: /next page/i }));
  });

  it("should navigate between pages", async () => {
    const { getByRole } = renderWithProviders("/");

    await userEvent.click(getByRole("button", { name: /next page/i }));
    expect(!!getByRole("button", { name: "Page 3" })); // based on the design & number of pages
  });

  it("should navigate to user posts", async () => {
    const { findByText } = renderWithProviders("/");

    const firstUser = await findByText("User1 Doe");
    await userEvent.click(firstUser);

    expect(router.history.location.pathname).toBe("/users/1/posts");
  });
});
