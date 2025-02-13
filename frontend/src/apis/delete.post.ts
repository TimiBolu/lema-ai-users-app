import { toast } from "react-hot-toast";
import EnvConfig from "../utils/env.config";

export const deletePost = async (postId: string) => {
  try {
    const response = await fetch(`${EnvConfig.apiBaseUrl}/posts/${postId}`, {
      method: "DELETE",
      credentials: "include",
      headers: {
        Authorization: `Bearer ${EnvConfig.jwtToken}`,
      },
    });

    if (!response.ok) {
      throw new Error("Failed to delete post");
    }

    toast.success("Post deleted successfully!");
  } catch (error) {
    toast.error("Failed to delete post.");
    throw error;
  }
};
