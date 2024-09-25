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

function angleDifference(angle1: number, angle2: number): number {
	const normalizedAngle1 = normalizeAngle(angle1);
	const normalizedAngle2 = normalizeAngle(angle2);
	const diff = normalizeAngle(normalizedAngle2 - normalizedAngle1);
	return diff > 180 ? diff - 360 : diff;
}

const calculateAngleRatio = (
	angle1: number,
	angle2: number,
	targetAngle: number,
): number => {
	// 角度を正規化
	const normalizedAngle1 = normalizeAngle(angle1);
	const normalizedAngle2 = normalizeAngle(angle2);
	const normalizedTargetAngle = normalizeAngle(targetAngle);

	// 角度差を計算
	const totalDiff = Math.abs(
		angleDifference(normalizedAngle1, normalizedAngle2),
	);

	// angle1とtargetAngleの間の差を計算
	let targetDiff = angleDifference(normalizedAngle1, normalizedTargetAngle);
	if (targetDiff < 0) {
		targetDiff += 360; // 負の差を修正
	}

	const ratio = targetDiff / totalDiff;

	return ratio;
};
export const calculateScreenPosition = (props: Props) => {
	const x = calculateAngleRatio(
		props.min.alpha,
		props.max.alpha,
		props.current.alpha,
	);
	const y = calculateAngleRatio(
		props.min.beta,
		props.max.beta,
		props.current.beta,
	);
	return {
		x: Math.max(Math.min(Math.round(x * 100) / 100, 1), 0),
		y: Math.max(Math.min(Math.round(y * 100) / 100, 1), 0),
	};
};
