import { useQuery } from "@tanstack/react-query";
import { fetchUserById } from "../apis/fetch.user.by.id";

export const useSingleUser = (userId: string) => {
  return useQuery({
    queryKey: ["user", userId],
    queryFn: () => fetchUserById(userId!),
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
