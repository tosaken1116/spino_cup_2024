import type { Meta } from "@storybook/react";
import { ChangePenSizeSlider } from ".";

export default {
	title: "ui/ChangePenSizeSlider",
	component: ChangePenSizeSlider,
} satisfies Meta<typeof ChangePenSizeSlider>;

export const Default = {
	args: {
		defaultValue: 5,
		onChange: (value: number) => {
			console.log(value);
		},
	},
};
