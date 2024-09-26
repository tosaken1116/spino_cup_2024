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

// export const calculateScreenPosition = (props: Props) => {
//   const { max, min, current } = props;
//   const add = max.alpha < min.alpha ? 360 : 0;
//   const left = max.alpha + add;
//   const right = min.alpha;
//   const x =
//     (current.alpha + (current.alpha < 180 ? add : 0) - right) / (left - right);
//   const y = (max.beta - current.beta) / (max.beta - min.beta);
//   return { x: 2 * (1 - x), y: 1 - y };
// };
function normalizeAngle(angle: number): number {
	return (angle + 360) % 360;
}

function angleDifference(minAngle: number, maxAngle: number): number {
	const diff = normalizeAngle(maxAngle - minAngle);
	return diff > 180 ? diff - 360 : diff;
}

function calculateAnglePosition(
	minAngle: number,
	maxAngle: number,
	targetAngle: number,
): number {
	const normalizedMin = normalizeAngle(minAngle);
	const normalizedMax = normalizeAngle(maxAngle);
	const normalizedTarget = normalizeAngle(targetAngle);

	const totalDiff = Math.abs(angleDifference(normalizedMin, normalizedMax));

	let targetDiff = angleDifference(normalizedMin, normalizedTarget);
	if (targetDiff < 0) {
		targetDiff += 360; // 負の差を修正
	}

	const ratio = targetDiff / totalDiff;

	return ratio;
}
export const calculateScreenPosition = (props: Props) => {
	const x = calculateAnglePosition(
		props.min.alpha,
		props.max.alpha,
		props.current.alpha,
	);
	const y = calculateAnglePosition(
		props.min.beta,
		props.max.beta,
		props.current.beta,
	);
	return {
		x: Math.max(Math.min(1 - Math.round(x * 10000) / 10000, 1), 0),
		y: Math.max(Math.min(1 - Math.round(y * 10000) / 10000, 1), 0),
	};
};
