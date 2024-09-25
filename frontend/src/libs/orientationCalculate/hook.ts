import { useCallback, useEffect, useRef, useState } from "react";
import type { ScreenSize } from "../../generated/wsClient/room/model";
import { calculateScreenPosition } from "../calculateScreenPosition";

export const useOrientationCalculate = ({
	handleChangePointerPosition,
	screenSize,
}: {
	screenSize: ScreenSize;
	handleChangePointerPosition: (props: { x: number; y: number }) => void;
}) => {
	const [leftTopOrientation, setLeftTopOrientation] = useState({
		alpha: 0,
		beta: 0,
	});
	const [permissionGranted, setPermissionGranted] = useState(false);
	const [rightBottomOrientation, setRightTopOrientation] = useState({
		alpha: 0,
		beta: 0,
	});
	const leftTopOrientationFetch = useRef<boolean>(false);
	const rightBottomOrientationFetch = useRef<boolean>(false);

	const handleSetLeftTopPoint = useCallback(() => {
		leftTopOrientationFetch.current = true;
	}, []);

	const handleSetRightBottomPoint = useCallback(() => {
		rightBottomOrientationFetch.current = true;
	}, []);

	const setCurrentPointer = (props: {
		alpha: number;
		beta: number;
		gamma: number;
	}) => {
		if (rightBottomOrientationFetch.current === true) {
			setRightTopOrientation({
				alpha: props.alpha,
				beta: props.beta,
			});
			rightBottomOrientationFetch.current = false;
		}
		if (leftTopOrientationFetch.current === true) {
			setLeftTopOrientation({
				alpha: props.alpha,
				beta: props.beta,
			});
			leftTopOrientationFetch.current = false;
		}
		handleChangePointerPosition(
			calculateScreenPosition({
				current: {
					alpha: props.alpha,
					beta: props.beta,
				},
				max: leftTopOrientation,
				min: rightBottomOrientation,
				screenSize,
			}),
		);
	};

	const handlePermissionGranted = async () => {
		// @ts-ignore
		await DeviceMotionEvent.requestPermission()
			// @ts-ignore
			.then((permissionState) => {
				if (permissionState === "granted") {
					setPermissionGranted(true);
				}
			})
			.catch(console.error);
	};

	// biome-ignore lint/correctness/useExhaustiveDependencies: <explanation>
	useEffect(() => {
		// @ts-ignore
		window.addEventListener("deviceorientation", setCurrentPointer);
		return () => {
			// @ts-ignore
			window.removeEventListener("deviceorientation", setCurrentPointer);
		};
	}, [permissionGranted, leftTopOrientation, rightBottomOrientation]);

	return {
		leftTopOrientation,
		rightBottomOrientation,
		permissionGranted,
		handleSetLeftTopPoint,
		handleSetRightBottomPoint,
		handlePermissionGranted,
	};
};
