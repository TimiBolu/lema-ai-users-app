import { useState } from "react";
import { useParams, useNavigate } from "@tanstack/react-router";

import { deletePost, createPost } from "../apis";
import DirectionBtn from "../components/DirectionBtn";
import Loader from "../components/Loader";
import PostCard from "../components/PostCard";
import EmptyCard from "../components/EmptyCard";
import { usePosts } from "../hooks/usePosts";
import { useSingleUser } from "../hooks/useSingleUser";

const PostManager = () => {
  const navigate = useNavigate();
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isCreatingPost, setIsCreatingPost] = useState(false);

  const { userId } = useParams({ strict: false });
  const { data: postsData, refetch, isPending } = usePosts(userId!);
  const posts = postsData?.data?.posts;

  const { data: userData } = useSingleUser(userId!);
  const user = userData?.data;

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
