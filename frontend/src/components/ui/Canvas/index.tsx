import { useEffect, useRef } from "react";

type Props = {
	circles: {
		x: number;
		y: number;
		color: string;
		size: number;
	}[];
	screenSize: {
		width: number;
		height: number;
	};
};
export const Canvas = (props: Props) => {
	const canvasRef = useRef<HTMLCanvasElement>(null);
	useEffect(() => {
		const canvas = canvasRef.current;
		if (canvas === null) {
			return;
		}
		const context = canvas.getContext("2d");
		if (context === null) {
			return;
		}
		for (const circle of props.circles) {
			context.beginPath();
			context.arc(circle.x, circle.y, circle.size, 0, 2 * Math.PI, false);
			context.fillStyle = circle.color;
			context.fill();
		}
	}, [props.circles]);
	return (
		<canvas
			ref={canvasRef}
			width={props.screenSize.width}
			height={props.screenSize.height}
		/>
	);
};
