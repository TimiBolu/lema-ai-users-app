import { User } from "../types/user";
import { ApiResponseSchemaFn, UserSchema } from "./api.schema";
import EnvConfig from "../utils/env.config";
import { ApiResponse } from "../types/apiresponse";

export const fetchUserById = async (
  userId: string,
): Promise<ApiResponse<User>> => {
  const response = await fetch(`${EnvConfig.apiBaseUrl}/users/${userId}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${EnvConfig.jwtToken}`,
    },
    credentials: "include",
  });

  if (!response.ok) {
    throw new Error("Failed to fetch user");
  }

  const data = await response.json();
  const parsed = ApiResponseSchemaFn(UserSchema).safeParse(data);
  if (!parsed.success) {
    throw new Error("Invalid response format");
  }

  return parsed.data;
};
