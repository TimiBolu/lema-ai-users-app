import { useQuery } from "@tanstack/react-query";
import { fetchPosts } from "../apis";

export const usePosts = (userId: string) => {
  return useQuery({
    queryKey: ["posts", userId],
    queryFn: () => fetchPosts(userId!),
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
