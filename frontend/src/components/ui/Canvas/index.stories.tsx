import type { Meta } from "@storybook/react";
import { Canvas } from ".";

export default {
	title: "ui/Canvas",
	component: Canvas,
} satisfies Meta<typeof Canvas>;

export const Default = () => {
	return (
		<Canvas
			circles={[
				{ x: 50, y: 50, color: "red", size: 10 },
				{ x: 100, y: 100, color: "blue", size: 5 },
			]}
			screenSize={{
				width: 200,
				height: 200,
			}}
		/>
	);
};
