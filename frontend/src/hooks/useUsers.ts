import { useQuery } from "@tanstack/react-query";
import { fetchUsers } from "../apis/fetch.users";

export const useUsers = (page: number) => {
  return useQuery({
    queryKey: ["users", page],
    queryFn: () => fetchUsers(page),
    retry: (failureCount, error) => {
      // Don't retry on 404s or validation errors
      if (error) {
        return false;
      }
      return failureCount < 3;
    },
    staleTime: 1000 * 60 * 5,
  });
};
