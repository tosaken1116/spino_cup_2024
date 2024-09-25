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
	const add = max.alpha < min.alpha ? 360 : 0;
	const left = max.alpha + add;
	const right = min.alpha;
	const x =
		(current.alpha + (current.alpha < 180 ? add : 0) - right) / (left - right);
	const y = (max.beta - current.beta) / (max.beta - min.beta);
	return { x: 1 - x, y: 1 - y };
};
