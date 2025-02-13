import { FC } from "react";
import AddIcon from "./icons/AddIcon";

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

export default EmptyCard;
