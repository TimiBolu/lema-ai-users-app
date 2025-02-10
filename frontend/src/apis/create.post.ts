import { toast } from "react-hot-toast";
import { PostSchema } from "./api.schema";
import EnvConfig from "../utils/env.config";

export const createPost = async (
  userId: string,
  title: string,
  body: string,
) => {
  try {
    const response = await fetch(`${EnvConfig.apiBaseUrl}/posts`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ title, body, userId }),
      credentials: "include",
    });

    if (!response.ok) {
      throw new Error("Failed to create post");
    }

    const data = await response.json();
    const parsed = PostSchema.safeParse(data);
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
