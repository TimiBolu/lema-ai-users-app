import React from "react";
import TrashIcon from "./icons/TrashIcon";

type Post = {
  id: string;
  title: string;
  body: string;
};

type PostCardProps = {
  post: Post;
  onDelete: (postId: string) => void;
};

const PostCard: React.FC<PostCardProps> = ({ post, onDelete }) => {
  return (
    <div className="relative border border-[#D5D7DA] p-6 text-[#535862] w-[270px] h-[293px] rounded-lg shadow-md">
      <button
        aria-label={"delete-post-" + post.id}
        onClick={() => onDelete(post.id)}
        className="absolute top-2 right-2 cursor-pointer"
      >
        <TrashIcon />
      </button>

      {/* Title with 2-line ellipsis */}
      <p
        className="text-lg font-medium mb-4 overflow-hidden text-ellipsis"
        style={{
          display: "-webkit-box",
          WebkitLineClamp: 2,
          WebkitBoxOrient: "vertical",
          wordBreak: "break-word", // Ensures long words wrap
        }}
      >
        {post.title}
      </p>

      {/* Body with 9-line ellipsis */}
      <p
        className="text-sm overflow-hidden text-ellipsis leading-tight"
        style={{
          display: "-webkit-box",
          WebkitLineClamp: 9,
          WebkitBoxOrient: "vertical",
          wordBreak: "break-word",
        }}
      >
        {post.body}
      </p>
    </div>
  );
};

export default PostCard;
