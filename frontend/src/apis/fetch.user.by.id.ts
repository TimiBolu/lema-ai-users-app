import { User } from "../types/user";
import { UserSchema } from "./api.schema";
import EnvConfig from "../utils/env.config";

export const fetchUserById = async (userId: string): Promise<User> => {
  const response = await fetch(`${EnvConfig.apiBaseUrl}/users/${userId}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });

  if (!response.ok) {
    throw new Error("Failed to fetch user");
  }

  const data = await response.json();
  const parsed = UserSchema.safeParse(data);
  if (!parsed.success) {
    throw new Error("Invalid response format");
  }

  return parsed.data;
};
