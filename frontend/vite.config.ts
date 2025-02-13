import dotenv from "dotenv";
import { defineConfig } from "vitest/config";
import react from "@vitejs/plugin-react";
import tailwindcss from "@tailwindcss/vite";
import { TanStackRouterVite } from "@tanstack/router-plugin/vite";

dotenv.config({ path: ".env.test" });
// https://vite.dev/config/
export default defineConfig({
  plugins: [
    TanStackRouterVite({ autoCodeSplitting: true }),
    react(),
    tailwindcss(),
  ],
  test: {
    globals: true,
    environment: "happy-dom",
    setupFiles: "./src/test-utils.tsx",
    coverage: {
      provider: "v8",
      include: ["src/**/*.{ts,tsx}"],
      exclude: ["**/*.test.tsx", "**/types.ts"],
    },
  },
});
