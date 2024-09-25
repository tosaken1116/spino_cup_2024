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
				{ x: 50, y: 50, color: "red" },
				{ x: 100, y: 100, color: "blue" },
			]}
			screenSize={{
				width: 200,
				height: 200,
			}}
		/>
	);
};
