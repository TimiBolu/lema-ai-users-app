import clsx from "clsx";
import { FC } from "react";
import DirectionIcon from "./icons/DirectionIcon";

type DirectionBtnProps = {
  text?: string;
  disabled?: boolean;
  onClick: () => void;
  direction: "back" | "next";
};
const DirectionBtn: FC<DirectionBtnProps> = ({
  disabled,
  onClick,
  text,
  direction,
}) => {
  return (
    <button
      aria-label={`${direction} page`}
      role={"button"}
      // name=""
      disabled={disabled}
      className={clsx("flex items-center gap-[8px]", {
        "cursor-pointer": !disabled,
      })}
      onClick={disabled ? () => {} : onClick}
    >
      <DirectionIcon direction={direction} disabled={disabled} />
      <p
        className={clsx("hidden md:block text-sm font-bold text-[#535862]", {
          "opacity-50": disabled,
        })}
      >
        {text || (direction === "back" ? "Previous" : "Next")}
      </p>
    </button>
  );
};

export default DirectionBtn;
