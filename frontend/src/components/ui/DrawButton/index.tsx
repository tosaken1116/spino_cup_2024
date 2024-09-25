import clsx from "clsx";
import { useState } from "react";
import { BruchIcon } from "../../icons/brush";

type Props = {
	onChangePointer: (isClicked: boolean) => void;
};
export const DrawButton = ({ onChangePointer }: Props) => {
	const [isClicked, setIsClicked] = useState(false);
	const handleChange = (isClicked: boolean) => {
		setIsClicked(isClicked);
		onChangePointer(isClicked);
	};
	return (
		<button
			type="button"
			onTouchStart={() => {
				handleChange(true);
			}}
			onTouchEnd={() => {
				handleChange(false);
			}}
			onMouseDown={() => {
				handleChange(true);
			}}
			onMouseUp={() => {
				handleChange(false);
			}}
			className={clsx(
				"w-full z-10 aspect-square relative shadow-lg transition-shadow shadow-black bg-gradient-to-br from-blue-700 to-blue-900 duration-300 flex items-center justify-center rounded-full",
				{
					"shadow-transparent ": isClicked,
				},
			)}
		>
			<BruchIcon width={64} height={64} className="z-10" />
		</button>
	);
};
