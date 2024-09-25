import { useState, type FC } from "react";
type Props = {
	defaultValue?: number;
	onChange: (value: number) => void;
};

export const ChangePenSizeSlider: FC<Props> = ({ defaultValue = 5, onChange }) => {
	const [value, setValue] = useState(defaultValue);
	return (
		<div className="flex flex-col gap-4">
			<span>ペンの太さ {value} px</span>
			<div className="relative w-fit">
				<label>
					<input
						type="range"
						min={1}
						max={20}
						value={value}
						onChange={(e) => {
							setValue(Number(e.target.value));
							onChange(Number(e.target.value))
						}}
					/>
				</label>
				<div
					className="absolute -bottom-4 translate-y-1/2 left-1/2 -translate-x-1/2 bg-red-200  rounded-full border-2 border-gray-400 z-10"
					style={{
						width: `${value}px`,
						height: `${value}px`,
					}}
				/>
			</div>
		</div>
	);
};
