import type { Meta, StoryObj } from "@storybook/react";
import { UserController } from ".";

export default {
	title: "ui/UserController",
	component: UserController,
} satisfies Meta<typeof UserController>;

export const Default = {
	args: {},
	decorators: [
		(Story) => (
			<div
				style={{
					width: "320px",
					height: "568px",
					backgroundColor: "white",
					border: "2px black",
				}}
			>
				<Story />
			</div>
		),
	],
} satisfies StoryObj<typeof UserController>;
