import { useState } from "react";
import { ChangePenSizeSlider } from "../../../../components/ui/ChangePenSizeSlider";
import { ColorPicker } from "../../../../components/ui/ColorPicker";
import { DrawButton } from "../../../../components/ui/DrawButton";
import { useOrientationCalculate } from "../../../../libs/orientationCalculate";
import type { UserAction } from "../../../../libs/wsClients";

type Props = Omit<UserAction, "type">;

export const UserController = ({
	handleChangePointerColor,
	isClicked,
	screenSize,
	handleClickPointer,
	handleChangePenSize,
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
		<div className="flex flex-col gap-4 h-full">
			<div className="flex flex-row gap-2 h-full">
				<div className="w-full flex justify-center items-center">
					<div className="relative w-48 h-28 rounded-md border-2">
						<div
							className="absolute w-2 h-2"
							style={{
								left: `${position.x}px`,
								top: `${position.y}px`,
							}}
						/>
					</div>
				</div>
				<div className="w-1/4">
					<ChangePenSizeSlider onChange={handleChangePenSize} />
				</div>
			</div>
			<div className="flex flex-row items-end h-full p-4">
				<div className="w-2/3">
					<ColorPicker onChangeColor={handleChangePointerColor} />
				</div>
				<div className="w-1/3">
					<DrawButton onChangePointer={handleClickPointer} />
				</div>
			</div>
		</div>
	);
};
