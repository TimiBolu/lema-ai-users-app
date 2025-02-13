import EnvConfig from "../utils/env.config";
import { UsersResponse } from "../types/user";
import { ApiResponseSchemaFn, UsersResponseSchema } from "./api.schema";
import { ApiResponse } from "../types/apiresponse";

export const fetchUsers = async (
  page: number,
  pageSize = 4,
): Promise<ApiResponse<UsersResponse>> => {
  const response = await fetch(
    `${EnvConfig.apiBaseUrl}/users?pageNumber=${page}&pageSize=${pageSize}`,
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
    throw new Error("Failed to fetch users");
  }

  const data = await response.json();
  const parsed = ApiResponseSchemaFn(UsersResponseSchema).safeParse(data);
  console.log({ data });
  console.log({ parsed });
  if (!parsed.success) {
    throw new Error("Invalid response format");
  }

  return parsed.data;
};
