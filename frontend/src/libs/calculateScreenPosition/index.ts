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

	const x =
		Math.abs(max.alpha - current.alpha) / Math.abs(max.alpha - min.alpha);
	const y = Math.abs(max.beta - current.beta) / Math.abs(max.beta - min.beta);

	return { x: Math.min(Math.max(x, 0), 1), y: Math.min(Math.max(y, 0), 1) };
};
