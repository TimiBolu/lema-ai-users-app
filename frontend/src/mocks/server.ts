import { setupServer } from "msw/node";
import { http, HttpResponse } from "msw";
import { MockData } from "./mock.data";

const users = MockData.getUsers();
let posts = MockData.getPosts();

export const server = setupServer(
  http.get("/api/users", ({ request }) => {
    const url = new URL(request.url);
    const page = Number(url.searchParams.get("pageNumber")) || 1;
    const pageSize = 4;

    return HttpResponse.json({
      users: users.slice((page - 1) * pageSize, page * pageSize),
      pagination: {
        currentPage: page,
        pageSize,
        totalPages: Math.ceil(users.length / pageSize),
        totalItems: users.length,
        hasNext: page < Math.ceil(users.length / pageSize),
        hasPrev: page > 1,
      },
    });
  }),

  http.get("/api/users/:userId", ({ params }) => {
    const user = users.find((u) => u.id === params.userId);
    return user
      ? HttpResponse.json(user)
      : new HttpResponse(null, {
          status: 404,
        });
  }),

  http.get("/api/posts", ({ request }) => {
    const url = new URL(request.url);
    const userId = url.searchParams.get("userId");
    const userPosts = posts.filter((p) => p.userId === userId);
    return HttpResponse.json(userPosts);
  }),

  http.post("/api/posts", async ({ request }) => {
    const { userId, title, body } = (await request.json()) as {
      userId: string;
      title: string;
      body: string;
    };
    const newPost = {
      id: String(posts.length + 1),
      userId,
      title,
      body,
      createdAt: new Date().toISOString(),
    };
    posts.push(newPost);
    return HttpResponse.json(newPost);
  }),

  http.delete("/api/posts/:postId", ({ params }) => {
    posts = posts.filter((p) => p.id !== params.postId);
    return new HttpResponse(null, {
      status: 204,
    });
  }),
);
