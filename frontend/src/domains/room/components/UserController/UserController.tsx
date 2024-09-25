import { ColorPicker } from "../../../../components/ui/ColorPicker";
import type { UserAction } from "../../../../libs/wsClients";

type Props = Omit<UserAction, "type">;

export const UserController = ({
	isClicked,
	color,
	handleChangePointerColor,
	handleChangePointerPosition,
	handleClickPointer,
}: Props) => {
	return (
		<div>
			<button type="button" className="w-32 h-32 rounded-full">
				クリック
			</button>
			<ColorPicker onChangeColor={handleChangePointerColor} />
		</div>
	);
};
