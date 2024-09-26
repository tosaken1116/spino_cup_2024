import { Config } from "tailwindcss";

const config: Config = {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      keyframes: {
        "fade-in": {
          from: {
            opacity: "0",
          },
          to: {
            opacity: "1",
          },
        },
        jump: {
          "0%": {
            transform: "translateY(0px)",
          },
          "50%": {
            transform: "translateY(-10px)",
          },
          "100%": {
            transform: "translateY(0px)",
          },
        },
      },
      animation: {
        jump: "jump 0.5s ease-in-out 1",
        "fade-in": "fade-in 0.5s ease-in-out backwards",
      },
    },
  },
  plugins: [],
};

export default config;
