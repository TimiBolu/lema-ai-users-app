import React from "react";
import "./loader.css";

interface LoaderProps {
  size?: number; // Default to 24px
  color?: string; // Default to currentColor
  className?: string; // Additional custom styling
}

const Loader: React.FC<LoaderProps> = ({
  size = 24,
  color = "currentColor",
  className = "",
}) => {
  const dotSize = size * 0.25; // Each dot is 25% of the loader size
  const spacing = size * 0.15; // Adjusts the distance between dots

  return (
    <div
      className={`loader ${className}`}
      style={{
        width: size,
        height: size,
        position: "relative", // Ensure relative positioning
      }}
    >
      {[...Array(4)].map((_, i) => (
        <div
          key={i}
          style={{
            width: dotSize,
            height: dotSize,
            backgroundColor: color,
            position: "absolute",
            left: i * (dotSize + spacing),
            top: "50%",
            transform: "translateY(-50%)", // Center vertically
          }}
          className={`dot dot-${i + 1}`}
        />
      ))}
    </div>
  );
};

export default Loader;
