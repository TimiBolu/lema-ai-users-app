import { toast } from "react-hot-toast";
import { ApiResponseSchemaFn, PostSchema } from "./api.schema";
import EnvConfig from "../utils/env.config";
import { ApiResponse } from "../types/apiresponse";
import { Post } from "../types/post";

export const createPost = async (
  userId: string,
  title: string,
  body: string,
): Promise<ApiResponse<Post>> => {
  try {
    const response = await fetch(`${EnvConfig.apiBaseUrl}/posts`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${EnvConfig.jwtToken}`,
      },
      body: JSON.stringify({ title, body, userId }),
      credentials: "include",
    });

    if (!response.ok) {
      throw new Error("Failed to create post");
    }

    const data = await response.json();
    const parsed = ApiResponseSchemaFn(PostSchema).safeParse(data);
    if (!parsed.success) {
      throw new Error("Invalid response format");
    }

    toast.success("Post created successfully!");
    return parsed.data;
  } catch (error) {
    toast.error("Failed to create post.");
    throw error;
  }
};
