import { useState } from "react";
import { QRCodeIcon } from "../../../../components/icons/qrcode";
import { Canvas } from "../../../../components/ui/Canvas";
import Modal from "../../../../components/ui/Modal";
import QRCode from "../../../../components/ui/QRCode";
import type { ScreenAction } from "../../../../libs/wsClients";

type Props = Omit<ScreenAction, "type">;
export const ScreenController = ({ positions }: Props) => {
	const [isOpen, setIsOpen] = useState(false);
	return (
		<div className="w-full h-full flex flex-row gap-4">
			<div
				className="relative border-4 border-yellow-200 rounded-xl"
				style={{
					width: `${1300}px`,
					height: `${750}px`,
				}}
			>
				<Canvas
					circles={positions
						.filter((position) => position.isClicked)
						.map((position) => ({
							x: position.x * 1300,
							y: position.y * 750,
							color: position.color,
							size: position.penSize,
						}))}
					screenSize={{
						width: 1300,
						height: 750,
					}}
				/>
				{positions.map((position) => (
					<div
						key={position.user.id}
						className="absolute flex flex-row gap-2 items-center justify-center"
						style={{
							left: `${position.x * 1300}px`,
							top: `${position.y * 750}px`,
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
			<div className="w-1/4">
				<button type="button" onClick={() => setIsOpen(true)}>
					<QRCodeIcon />
				</button>
				<Modal isOpen={isOpen} onClose={() => setIsOpen(false)}>
					<QRCode url={window.location.href} />
				</Modal>
				<p>現在の参加者</p>
				{positions.map((position) => (
					<div
						key={position.user.id}
						className="flex flex-row gap-2 items-center"
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
