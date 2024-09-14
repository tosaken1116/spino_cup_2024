import react from "@vitejs/plugin-react-swc";
import { defineConfig } from "vitest/config";

// https://vitejs.dev/config/
export default defineConfig({
  test: {
    exclude: ["**/node_modules/**"],
    coverage: {
      exclude: ["**/node_modules/**", "**/*.config.*", "**/*.d.ts"],
    },
  },
  plugins: [react()],
});
