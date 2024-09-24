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
	const { max, min, current, screenSize } = props;

	const x =
		(current.alpha / (max.alpha - min.alpha)) * screenSize.width +
		screenSize.width / 2;
	const y =
		(current.beta / (max.beta - min.beta)) * screenSize.height +
		screenSize.height / 2;

	return { x, y };
};
