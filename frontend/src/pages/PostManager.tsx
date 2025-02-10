import { FC, useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { useParams, useNavigate } from "@tanstack/react-router";

import TrashIcon from "../components/icons/TrashIcon";

import { fetchPosts, deletePost, createPost } from "../apis";
import DirectionBtn from "../components/DirectionBtn";
import { fetchUserById } from "../apis/fetch.user.by.id";
import AddIcon from "../components/icons/AddIcon";
import { Post } from "../types/post";
import Loader from "../components/Loader";

type EmptyCardProps = {
  onClick: () => void;
};
const EmptyCard: FC<EmptyCardProps> = ({ onClick }) => {
  return (
    <div
      onClick={onClick}
      className="w-[270px] h-[293px] rounded-[8px] flex justify-center items-center cursor-pointer"
      style={{
        backgroundImage: `url("data:image/svg+xml,%3csvg width='100%25' height='100%25' xmlns='http://www.w3.org/2000/svg'%3e%3crect width='100%25' height='100%25' fill='none' rx='8' ry='8' stroke='%23D5D7DAFF' stroke-width='2' stroke-dasharray='6%2c 14' stroke-dashoffset='0' stroke-linecap='square'/%3e%3c/svg%3e")`,
        backgroundRepeat: "no-repeat",
        backgroundSize: "cover",
      }}
    >
      <div className="flex flex-col items-center gap-[8px]">
        <AddIcon />
        <p className="text-[#717680] text-[14px]/[20px] font-[600]">New Post</p>
      </div>
    </div>
  );
};

type PostCardProps = {
  post: Post;
  onDelete: (postId: string) => void;
};
const PostCard: FC<PostCardProps> = ({ post, onDelete }) => {
  return (
    <div className="relative border-1 border-[#D5D7DA] p-[24px] text-[#535862] w-[270px] h-[293px] rounded-[8px] shadow-[0px_2px_4px_-2px_rgba(10,13,18,0.06),0px_4px_8px_-2px_rgba(10,13,18,0.1)]">
      <button
        aria-label={"delete-post-" + post.id}
        onClick={() => onDelete(post.id)}
        className="absolute top-[10px] right-[10px] cursor-pointer"
      >
        <TrashIcon />
      </button>
      <p className="text-[18px]/[20px] font-[500] mb-[1rem]">{post.title}</p>
      <p
        className="text-[14px]/[20px]  overflow-hidden text-ellipsis"
        style={{
          display: "-webkit-box",
          WebkitLineClamp: 9,
          WebkitBoxOrient: "vertical",
        }}
      >
        {post.body}
      </p>
    </div>
  );
};

const PostManager = () => {
  const navigate = useNavigate();
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isCreatingPost, setIsCreatingPost] = useState(false);

  const { userId } = useParams({ strict: false });
  const {
    data: posts,
    refetch,
    isPending,
  } = useQuery({
    queryKey: ["posts", userId],
    queryFn: () => fetchPosts(userId!),
  });

  const { data: user } = useQuery({
    queryKey: ["user", userId],
    queryFn: () => fetchUserById(userId!),
  });

  const handleDeletePost = async (postId: string) => {
    await deletePost(postId);

    refetch();
  };

  const handleAddPost = async () => {
    setIsCreatingPost(true);
    try {
      await createPost(userId!, title, body);

      setIsModalOpen(false);
      setTitle("");
      setBody("");
      refetch();
    } finally {
      setIsCreatingPost(false);
    }
  };

  const noOfPosts = posts?.length;

  return isPending ? (
    <div className="flex w-full h-[50dvh] justify-center items-center">
      <Loader color="#bdb1c6" size={50} />
    </div>
  ) : (
    <div className="space-y-4 pb-8">
      <DirectionBtn
        direction="back"
        text="Back to Users"
        onClick={() => {
          navigate({ to: "/" });
        }}
      />

      <p className="text-[36px]/[43.57px] text-[#181D27] my-[1rem]">
        {user?.firstname} {user?.lastname}
      </p>

      <p className="text-[#535862] text-[14px]/[20px] mb-[24px]">
        {user?.email} • {noOfPosts} {noOfPosts === 1 ? "Post" : "Posts"}
      </p>

      <div className="flex flex-wrap gap-[23px] w-full justify-center md:justify-start">
        <EmptyCard onClick={() => setIsModalOpen(true)} />
        {posts?.map((post) => (
          <PostCard key={post.id} post={post} onDelete={handleDeletePost} />
        ))}
      </div>

      {/* Modal */}
      {isModalOpen && (
        <div
          className="fixed inset-0 backdrop-blur-[1px] flex items-center justify-center p-4"
          onClick={() => setIsModalOpen(false)}
        >
          <div
            className="w-full max-w-[95%] md:max-w-[90%] lg:max-w-[679px] h-auto min-h-[483px]
                       bg-white p-4 sm:p-6 md:p-[24px] rounded-[8px]
                       shadow-[0px_4px_8px_-3px_rgba(0,0,0,0.1),1px_10px_20px_-5px_rgba(0,0,0,0.08)]"
            onClick={(e) => e.stopPropagation()}
          >
            <h2 className="text-2xl sm:text-3xl md:text-[36px]/[43.57px] text-[#181D27] font-[500] mb-4 sm:mb-6 md:mb-[24px]">
              New Post
            </h2>
            <label
              className="block text-base sm:text-lg md:text-[18px] font-[500] text-[#535862] mb-2 sm:mb-3 md:mb-[10px]"
              htmlFor="title"
            >
              Post title
            </label>
            <input
              id="title"
              type="text"
              placeholder="Give your post a title"
              className="px-3 sm:px-4 md:px-[1rem] text-sm sm:text-base md:text-[14px]/[21px]
                        text-[400] w-full border border-[#E2E8F0] placeholder-[#94A3B8]
                        p-2 rounded-[4px] mb-4 sm:mb-6 md:mb-[24px] focus:outline-none"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
            />
            <label
              className="block text-base sm:text-lg md:text-[18px] font-[500] text-[#535862] mb-2 sm:mb-3 md:mb-[10px]"
              htmlFor="body"
            >
              Post content
            </label>
            <textarea
              id="body"
              placeholder="Write something mind-blowing"
              className="px-3 sm:px-4 md:px-[1rem] text-sm sm:text-base md:text-[14px]/[21px]
                        font-[400] w-full border border-[#E2E8F0] h-[140px] sm:h-[160px] md:h-[179px]
                        placeholder-[#94A3B8] p-2 rounded-[4px] mb-3 sm:mb-4 md:mb-[12px] focus:outline-none"
              rows={4}
              value={body}
              onChange={(e) => setBody(e.target.value)}
            />
            <div className="flex justify-end text-sm sm:text-base md:text-[14px] space-x-3">
              <button
                onClick={() => setIsModalOpen(false)}
                className="cursor-pointer border border-[#E2E8F0] font-[400] text-[#334155]
                          px-3 sm:px-4 py-1.5 sm:py-2 rounded"
              >
                Cancel
              </button>
              <button
                onClick={handleAddPost}
                className="flex items-center gap-2 cursor-pointer bg-[#334155] w-fit text-white
                          font-[600] px-3 sm:px-4 py-1.5 sm:py-2 rounded"
                disabled={!title || !body}
              >
                <p>Publish</p>
                {isCreatingPost && (
                  <Loader color="#FFFFFF" className="mr-[0.5rem] mb-1" />
                )}
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default PostManager;
