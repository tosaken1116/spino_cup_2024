import { ScreenController } from "../../../domains/room/components/ScrrenController/ScreenController";
import { UserController } from "../../../domains/room/components/UserController";
import { useRoomUserWSClient } from "../../../libs/wsClients";

type Props = {
	id: string;
};
export const RoomDetailPage = (props: Props) => {
	const actions = useRoomUserWSClient(props.id);
	if (actions === null) {
		return <p>...connecting</p>;
	}
	if (actions.type === "user") {
		return (
			<div>
				<UserController {...actions} />
			</div>
		);
	}

	return (
		<div>
			<ScreenController {...actions} />
		</div>
	);
};
