import clsx from "clsx";
import { useState } from "react";
import { ColorPicker } from "../../../../components/ui/ColorPicker";
import { useOrientationCalculate } from "../../../../libs/orientationCalculate";
import type { UserAction } from "../../../../libs/wsClients";

type Props = Omit<UserAction, "type">;

export const UserController = ({
	handleChangePointerColor,
	isClicked,
	screenSize,
	handleClickPointer,
	handleChangePointerPosition,
}: Props) => {
	const [position, setPosition] = useState({
		x: 0,
		y: 0,
	});
	const {
		leftTopOrientation,
		rightBottomOrientation,
		handlePermissionGranted,
		handleSetLeftTopPoint,
		handleSetRightBottomPoint,
		permissionGranted,
	} = useOrientationCalculate({
		screenSize,
		handleChangePointerPosition: (props) => {
			handleChangePointerPosition(props);
			setPosition(props);
		},
	});
	return (
		<div>
			{permissionGranted ? <p>Permission granted</p> : <p>Permission denied</p>}
			<div>
				<p>左上の座標</p>
				<p>α: {leftTopOrientation.alpha}</p>
				<p>β: {leftTopOrientation.beta}</p>
			</div>
			<div>
				<p>現在の座標</p>
				<p>x: {position.x * 100}</p>
				<p>y: {position.y * 100}</p>
			</div>
			<div className="relative w-[100px] h-[100px] border-[1px]">
				<div
					className="absolute w-2 h-2 rounded-full bg-red-200"
					style={{
						left: `${position.x * 100}px`,
						top: `${position.y * 100}px`,
					}}
				/>
			</div>

			<div>
				<p>右下の座標</p>
				<p>α: {rightBottomOrientation.alpha}</p>
				<p>β: {rightBottomOrientation.beta}</p>
			</div>
			<button type="button" onClick={handlePermissionGranted}>
				権限を設定
			</button>
			<button type="button" onClick={handleSetRightBottomPoint}>
				右下の座標を設定
			</button>
			<button type="button" onClick={handleSetLeftTopPoint}>
				左上の座標を設定
			</button>
			<button
				type="button"
				onTouchStart={() => handleClickPointer(true)}
				onTouchEnd={() => handleClickPointer(false)}
				onMouseDown={() => handleClickPointer(true)}
				onMouseUp={() => handleClickPointer(false)}
				className={clsx("w-32 h-32 rounded-full", {
					"bg-blue-500": isClicked,
					"bg-blue-300": !isClicked,
				})}
			>
				クリック
			</button>
			<ColorPicker onChangeColor={handleChangePointerColor} />
		</div>
	);
};
