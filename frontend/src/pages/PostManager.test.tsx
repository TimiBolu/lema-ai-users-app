import {
  it,
  expect,
  afterAll,
  describe,
  beforeAll,
  afterEach,
  beforeEach,
} from "vitest";
import "@testing-library/jest-dom";
import { http, HttpResponse } from "msw";

import { server } from "../mocks/server";
import EnvConfig from "../utils/env.config";
import { MockData } from "../mocks/mock.data";
import { renderWithProviders, userEvent, queryClient } from "../test-utils";

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

describe("PostManager", () => {
  beforeEach(() => {
    server.use(
      http.get("/api/users/:userId", ({ params }) => {
        const user = MockData.getUser(params.userId as string);
        return HttpResponse.json(user);
      }),
      http.get("/api/posts", ({ request }) => {
        const url = new URL(request.url);
        const userId = url.searchParams.get("userId");

        if (!userId) {
          return HttpResponse.json(
            { error: "User ID is required" },
            { status: 400 },
          );
        }

        return HttpResponse.json(
          MockData.getPosts().filter((data) => data.userId === userId),
        );
      }),
      http.post("/api/posts", async ({ request }) => {
        const { title, body, userId } = (await request.json()) as {
          title: string;
          body: string;
          userId: string;
        };
        const newPost = MockData.insertPost(title, body, userId);
        return HttpResponse.json(newPost);
      }),
      http.delete("/api/posts/:postId", async ({ params }) => {
        MockData.removePost((data) => data.id !== params.postId);
        return HttpResponse.json(null, { status: 200 });
      }),
    );
  });

  it("should render user info and posts", async () => {
    const { findByText } = renderWithProviders("/users/$userId/posts", {
      userId: "1",
    });

    await findByText(/John\s*Doe/);
    expect(await findByText(/john.doe@example.com/)).toBeInTheDocument();
    expect(await findByText("Second Post by John")).toBeInTheDocument();
    expect(await findByText(/3\s*Posts/)).toBeInTheDocument();
  });

  it("should create new post", async () => {
    const { findByText, getByPlaceholderText, getByRole } = renderWithProviders(
      "/users/$userId/posts",
      { userId: "1" },
    );

    await userEvent.click(await findByText("New Post"));

    await userEvent.type(
      getByPlaceholderText("Give your post a title"),
      "New Title",
    );
    await userEvent.type(
      getByPlaceholderText("Write something mind-blowing"),
      "New content",
    );

    await userEvent.click(getByRole("button", { name: /publish/i }));

    await findByText("New Title");
  });

  it("should handle post deletion", async () => {
    const { findByText, queryByText, getByRole } = renderWithProviders(
      "/users/$userId/posts",
      { userId: "1" },
    );

    await findByText("First Post by John");
    const deleteButton = getByRole("button", { name: "delete-post-1" }); // since the id is 1
    await userEvent.click(deleteButton);

    expect(queryByText("First Post by John")).not.toBeInTheDocument();
  });
});
