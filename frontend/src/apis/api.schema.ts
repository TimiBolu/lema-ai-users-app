import { z } from "zod";

export const UserSchema = z.object({
  id: z.string(),
  firstname: z.string(),
  lastname: z.string(),
  email: z.string().email(),
  address: z.object({
    street: z.string(),
    city: z.string(),
    state: z.string(),
    zipCode: z.string(),
  }),
});

export const UsersResponseSchema = z.object({
  users: z.array(UserSchema),
  pagination: z.object({
    totalPages: z.number(),
    currentPage: z.number(),
    pageSize: z.number(),
    totalItems: z.number(),
    hasNext: z.boolean(),
    hasPrev: z.boolean(),
  }),
});

export const PostSchema = z.object({
  id: z.string(),
  userId: z.string(),
  title: z.string(),
  body: z.string(),
  createdAt: z.string(),
});

export const PostResponseSchema = z.object({
  posts: z.array(PostSchema),
});

export const ApiResponseSchemaFn = <T extends z.ZodTypeAny>(dataSchema: T) =>
  z.object({
    success: z.boolean(),
    message: z.string(),
    data: dataSchema,
  });
