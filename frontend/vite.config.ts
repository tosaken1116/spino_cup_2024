import { TanStackRouterVite } from "@tanstack/router-vite-plugin";
import react from "@vitejs/plugin-react-swc";
import dotenv from "dotenv";
import { defineConfig } from "vitest/config";
// https://vitejs.dev/config/
export default defineConfig({
  test: {
    include: ["src/**/*.test.ts"],
    exclude: ["src/**/*.test.tsx", "src/**/*.d.ts"],
    coverage: {
      include: ["src/**/*.ts"],
      exclude: ["src/**/*.test.ts", "src/**/*.test.tsx", "src/**/*.d.ts"],
    },
    environment: "jsdom",
    env: dotenv.config({ path: ".env.test" }).parsed,
  },
  plugins: [react(), TanStackRouterVite()],
});
