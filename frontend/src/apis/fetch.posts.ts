import { Post } from "../types/post";
import EnvConfig from "../utils/env.config";
import { PostResponseSchema } from "./api.schema";

export const fetchPosts = async (userId: string): Promise<Array<Post>> => {
  const response = await fetch(
    `${EnvConfig.apiBaseUrl}/posts?userId=${userId}`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    },
  );

  if (!response.ok) {
    throw new Error("Failed to fetch posts");
  }

  const data = await response.json();
  const parsed = PostResponseSchema.safeParse(data);
  if (!parsed.success) {
    throw new Error("Invalid response format");
  }

  return parsed.data;
};
