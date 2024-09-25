type Props = {
	max: {
		alpha: number;
		beta: number;
	};
	min: {
		alpha: number;
		beta: number;
	};
	center?: {
		alpha: number;
		beta: number;
	};
	current: {
		alpha: number;
		beta: number;
	};
	screenSize: {
		width: number;
		height: number;
	};
};

export const calculateScreenPosition = (props: Props) => {
	const { max, min, current } = props;

	const x = Math.tan(
		(Math.abs(max.alpha - current.alpha) / Math.abs(max.alpha - min.alpha)) *
			Math.PI,
	);

	const y = Math.tan(
		(Math.abs(max.beta - current.beta) / Math.abs(max.beta - min.beta)) *
			Math.PI,
	);

	return { x: Math.min(Math.max(x, 0), 1), y: Math.min(Math.max(y, 0), 1) };
};
