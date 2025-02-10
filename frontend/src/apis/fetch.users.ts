import EnvConfig from "../utils/env.config";
import { UsersResponse } from "../types/user";
import { UsersResponseSchema } from "./api.schema";

export const fetchUsers = async (
  page: number,
  pageSize = 4,
): Promise<UsersResponse> => {
  const response = await fetch(
    `${EnvConfig.apiBaseUrl}/users?pageNumber=${page}&pageSize=${pageSize}`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    },
  );

  if (!response.ok) {
    throw new Error("Failed to fetch users");
  }

  const data = await response.json();
  const parsed = UsersResponseSchema.safeParse(data);
  if (!parsed.success) {
    throw new Error("Invalid response format");
  }

  return parsed.data;
};
