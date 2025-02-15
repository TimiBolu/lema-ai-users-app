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
  const [titleError, setTitleError] = useState("");
  const [bodyError, setBodyError] = useState("");
  const isPostDisbaled = !title || !body || !!titleError || !!bodyError;

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

      <div className="flex flex-wrap gap-[23px] w-full justify-center sm:justify-start">
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
            className="w-full md:w-[679px] max-w-[679px] lg:max-w-[679px] h-auto min-h-[483px]
                       bg-white p-4 sm:p-6 md:p-[24px] rounded-[8px]
                       shadow-[0px_4px_8px_-3px_rgba(0,0,0,0.1),1px_10px_20px_-5px_rgba(0,0,0,0.08)]"
            onClick={(e) => e.stopPropagation()}
          >
            <h2 className="text-2xl sm:text-3xl md:text-[36px]/[43.57px] text-[#181D27] font-[500] mb-4 sm:mb-6 md:mb-[24px]">
              New Post
            </h2>

            <div className="mb-4 sm:mb-6 md:mb-[24px]">
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
                className={`px-3 sm:px-4 md:px-[1rem] text-sm sm:text-base md:text-[14px]/[21px]
                          w-full border border-[#E2E8F0] placeholder-[#94A3B8] p-2 rounded-[4px]
                          focus:outline-none ${titleError ? "border-red-500" : "border-[#E2E8F0]"}`}
                value={title}
                maxLength={100}
                onChange={(e) => {
                  const value = e.target.value;
                  setTitle(value);

                  const isValid = /^[a-zA-Z0-9\s.,!?'"-]+$/.test(value);
                  setTitleError(
                    !isValid && value !== ""
                      ? "Contains invalid special characters"
                      : "",
                  );
                }}
              />
              <div className="flex justify-between mt-1">
                {titleError && (
                  <span className="text-red-500 text-sm">{titleError}</span>
                )}
                <span className="text-sm text-gray-500 ml-auto">
                  {title.length}/100
                </span>
              </div>
            </div>

            <div className="mb-3 sm:mb-4 md:mb-[12px]">
              <label
                className="block text-base sm:text-lg md:text-[18px] font-[500] text-[#535862] mb-2 sm:mb-3 md:mb-[10px]"
                htmlFor="body"
              >
                Post content
              </label>
              <textarea
                id="body"
                placeholder="Write something mind-blowing"
                className={`resize-none px-3 h-[179px]  sm:px-4 md:px-[1rem] text-sm sm:text-base md:text-[14px]/[21px]
                          w-full border border-[#E2E8F0] placeholder-[#94A3B8] p-2 rounded-[4px]
                          focus:outline-none ${bodyError ? "border-red-500" : "border-[#E2E8F0]"}`}
                rows={4}
                value={body}
                maxLength={500}
                onChange={(e) => {
                  const value = e.target.value;
                  setBody(value);
                  setBodyError(
                    value.length >= 500 ? "Maximum 500 characters reached" : "",
                  );
                }}
              />
              <div className="flex justify-between mt-1">
                {bodyError && (
                  <span className="text-red-500 text-sm">{bodyError}</span>
                )}
                <span className="text-sm text-gray-500 ml-auto">
                  {body.length}/500
                </span>
              </div>
            </div>

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
                className={`flex items-end gap-2 bg-[#334155] w-fit text-white font-[600] px-3 sm:px-4
                  py-1.5 sm:py-2 rounded disabled:opacity-50 disabled:cursor-not-allowed cursor-pointer`}
                disabled={isPostDisbaled}
              >
                <p>Publish</p>
                {isCreatingPost && (
                  <Loader color="#FFFFFF" className="mr-[0.5rem]" />
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
