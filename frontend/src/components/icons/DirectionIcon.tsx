import { FC } from "react";
import clsx from "clsx";

type direction = "back" | "next";
type DirectionIconProps = {
  direction: direction;
  disabled?: boolean;
};
const DirectionIcon: FC<DirectionIconProps> = ({ direction, disabled }) => (
  <div
    className={clsx("cursor-pointer", {
      "opacity-50 pointer-events-none": disabled, // Faded & non-interactive if disabled
    })}
  >
    {direction === "back" ? (
      <svg
        width="20"
        height="20"
        viewBox="0 0 20 20"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M15.8334 10H4.16669M4.16669 10L10 15.8334M4.16669 10L10 4.16669"
          stroke="#717680"
          strokeWidth="1.67"
          strokeLinecap="round"
          strokeLinejoin="round"
        />
      </svg>
    ) : (
      <svg
        width="20"
        height="20"
        viewBox="0 0 20 20"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M4.16663 10H15.8333M15.8333 10L9.99996 4.16669M15.8333 10L9.99996 15.8334"
          stroke="#717680"
          strokeWidth="1.67"
          strokeLinecap="round"
          strokeLinejoin="round"
        />
      </svg>
    )}
  </div>
);

export default DirectionIcon;
