import react from "@vitejs/plugin-react-swc";
import { defineConfig } from "vitest/config";
import { TanStackRouterVite } from "@tanstack/router-vite-plugin";

// https://vitejs.dev/config/
export default defineConfig({
  test: {
    include: ["src/**/*.ts"],
    exclude: ["src/**/*.test.ts", "src/**/*.test.tsx", "src/**/*.d.ts"],
    coverage: {
      include: ["src/**/*.ts"],
      exclude: ["src/**/*.test.ts", "src/**/*.test.tsx", "src/**/*.d.ts"],
    },
  },
  plugins: [react(), TanStackRouterVite()],
});
