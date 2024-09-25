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
				"w-28 h-28 relative bg-blue-300 shadow-lg transition-shadow shadow-black before:contents-[''] bg-gradient-to-br from-blue-700 to-blue-900 duration-300 before:bg-yellow-200 before before:absolute before:left-1/2 before:top-1/2 before:-translate-x-1/2 before:-translate-y-1/2 before:w-36 before:rounded-full before:-z-10 before:block before:h-36 flex items-center justify-center rounded-full",
				{
					"bg-blue-500 shadow-transparent ": isClicked,
				},
			)}
		>
			<BruchIcon width={64} height={64} className="z-10" />
		</button>
	);
};
