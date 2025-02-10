import { Post } from "../types/post";
import { User } from "../types/user";

const mockUsersData: User[] = [
  {
    id: "1",
    firstname: "John",
    lastname: "Doe",
    email: "john.doe@example.com",
    address: {
      street: "123 Main St",
      city: "Springfield",
      state: "IL",
      zipCode: "62701",
    },
  },
  {
    id: "2",
    firstname: "Jane",
    lastname: "Smith",
    email: "jane.smith@example.com",
    address: {
      street: "456 Oak St",
      city: "Chicago",
      state: "IL",
      zipCode: "60602",
    },
  },
  {
    id: "3",
    firstname: "Michael",
    lastname: "Johnson",
    email: "michael.johnson@example.com",
    address: {
      street: "789 Pine St",
      city: "Peoria",
      state: "IL",
      zipCode: "61602",
    },
  },
];

const mockPostsData: Post[] = [
  {
    id: "1",
    userId: "1",
    title: "First Post by John",
    body: "This is the body of the first post by John.",
    createdAt: "2025-02-09T10:00:00Z",
  },
  {
    id: "2",
    userId: "1",
    title: "Second Post by John",
    body: "This is the body of the second post by John.",
    createdAt: "2025-02-09T11:00:00Z",
  },
  {
    id: "3",
    userId: "1",
    title: "Third Post by John",
    body: "This is the body of the third post by John.",
    createdAt: "2025-02-09T12:00:00Z",
  },
  {
    id: "4",
    userId: "2",
    title: "First Post by Jane",
    body: "This is the body of the first post by Jane.",
    createdAt: "2025-02-09T13:00:00Z",
  },
  {
    id: "5",
    userId: "2",
    title: "Second Post by Jane",
    body: "This is the body of the second post by Jane.",
    createdAt: "2025-02-09T14:00:00Z",
  },
  {
    id: "6",
    userId: "2",
    title: "Third Post by Jane",
    body: "This is the body of the third post by Jane.",
    createdAt: "2025-02-09T15:00:00Z",
  },
  {
    id: "7",
    userId: "3",
    title: "First Post by Michael",
    body: "This is the body of the first post by Michael.",
    createdAt: "2025-02-09T16:00:00Z",
  },
  {
    id: "8",
    userId: "3",
    title: "Second Post by Michael",
    body: "This is the body of the second post by Michael.",
    createdAt: "2025-02-09T17:00:00Z",
  },
  {
    id: "9",
    userId: "3",
    title: "Third Post by Michael",
    body: "This is the body of the third post by Michael.",
    createdAt: "2025-02-09T18:00:00Z",
  },
];

const clone = <T>(data: T) => {
  return JSON.parse(JSON.stringify(data)) as T;
};

let _mockUsersData = clone(mockUsersData);
let _mockPostsData = clone(mockPostsData);

let everIncreasingPostsId = _mockPostsData.length;

export const MockData = {
  seed: () => {
    _mockUsersData = clone(mockUsersData);
    _mockPostsData = clone(mockPostsData);
  },
  getUsers: () => {
    return [..._mockUsersData];
  },
  getUser: (userId: string) => {
    return _mockUsersData.find((user) => user.id === userId);
  },
  getPosts: () => {
    return [..._mockPostsData];
  },
  setUsers: (usersData: User[]) => {
    _mockUsersData = usersData;
  },
  setPosts: (postsData: Post[]) => {
    _mockPostsData = postsData;
  },
  insertPost: (title: string, body: string, userId: string) => {
    const newPost = {
      id: String(everIncreasingPostsId++),
      userId,
      title,
      body,
      createdAt: new Date().toISOString(),
    };
    _mockPostsData.push(newPost);
    return newPost;
  },
  removePost: (predicate: (post: Post) => boolean) => {
    _mockPostsData = _mockPostsData.filter(predicate);
  },
};
