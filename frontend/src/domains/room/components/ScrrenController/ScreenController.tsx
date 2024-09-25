import { useState } from "react";
import { Canvas } from "../../../../components/ui/Canvas";
import type { ScreenAction } from "../../../../libs/wsClients";

type Props = Omit<ScreenAction, "type">;
export const ScreenController = ({ positions }: Props) => {
	const [screenSize, setScreenSize] = useState({ width: 0, height: 0 });
	return (
		<div>
			<label>
				横幅
				<input
					onChange={(e) => {
						setScreenSize((prev) => ({
							...prev,
							width: Number(e.target.value),
						}));
					}}
				/>
			</label>
			<label>
				高さ
				<input
					onChange={(e) => {
						setScreenSize((prev) => ({
							...prev,
							height: Number(e.target.value),
						}));
					}}
				/>
			</label>
			<div
				className="relative border border-black rounded-sm"
				style={{
					width: `${screenSize.width}px`,
					height: `${screenSize.height}px`,
				}}
			>
				<Canvas
					circles={positions.map((position) => ({
						x: position.x * screenSize.width,
						y: position.y * screenSize.height,
						color: position.color,
					}))}
					screenSize={screenSize}
				/>
				{positions.map((position) => (
					<div
						key={position.id}
						className="absolute"
						style={{
							left: `${position.x * screenSize.width}px`,
							top: `${position.y * screenSize.height}px`,
						}}
					>
						<p>{position.id}</p>
						<p className="text-xl">{position.isClicked ? "🖕" : "🫵"}</p>
						<div
							className="w-4 h-4 rounded-full"
							style={{
								backgroundColor: position.color,
							}}
						/>
					</div>
				))}
			</div>
		</div>
	);
};
