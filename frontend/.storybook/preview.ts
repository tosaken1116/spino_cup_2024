import type { Preview } from "@storybook/react";
import { withScreenshot } from "storycap";

export const decorators = [withScreenshot];

const preview: Preview = {
  parameters: {
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i,
      },
    },
    screenshot: {
      fullPage: false,
      delay: 0,
      viewports: {
        desktop: { width: 1920, height: 1080 },
        tablet: { width: 768, height: 1024 },
        mobile: { width: 360, height: 800, isMobile: true, hasTouch: true },
      },
    },
  },
};

export default preview;
