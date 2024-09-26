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
					circles={positions
						.filter((position) => position.isClicked)
						.map((position) => ({
							x: position.x * screenSize.width,
							y: position.y * screenSize.height,
							color: position.color,
							size: position.penSize,
						}))}
					screenSize={screenSize}
				/>
				{positions.map((position) => (
					<div
						key={position.user.id}
						className="absolute flex flex-row gap-2 items-center justify-center"
						style={{
							left: `${position.x * screenSize.width}px`,
							top: `${position.y * screenSize.height}px`,
						}}
					>
						<img
							alt={position.user.name}
							height={32}
							width={32}
							className="rounded-full"
							src={position.user.avatarUrl}
						/>
						<p className="font-semibold">{position.user.name}</p>
						<div
							className="rounded-full"
							style={{
								width: `${position.penSize}px`,
								height: `${position.penSize}px`,
								backgroundColor: position.color,
							}}
						/>
					</div>
				))}
			</div>
		</div>
	);
};
