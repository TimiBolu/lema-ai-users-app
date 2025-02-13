import { ApiResponse } from "../types/apiresponse";
import { Post } from "../types/post";
import EnvConfig from "../utils/env.config";
import { ApiResponseSchemaFn, PostResponseSchema } from "./api.schema";

export const fetchPosts = async (
  userId: string,
): Promise<ApiResponse<{ posts: Array<Post> }>> => {
  const response = await fetch(
    `${EnvConfig.apiBaseUrl}/posts?userId=${userId}`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${EnvConfig.jwtToken}`,
      },
      credentials: "include",
    },
  );

  if (!response.ok) {
    throw new Error("Failed to fetch posts");
  }

  const data = await response.json();
  const parsed = ApiResponseSchemaFn(PostResponseSchema).safeParse(data);
  if (!parsed.success) {
    throw new Error("Invalid response format");
  }

  return parsed.data;
};
