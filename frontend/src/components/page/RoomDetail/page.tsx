import { useRoomUserWSClient } from "../../../libs/wsClients";

type Props = {
	id: string;
};
export const RoomDetailPage = (props: Props) => {
	const { handleClick, isConnected } = useRoomUserWSClient(props.id);
	return (
		<div>
			{isConnected ? (
				<div>
					<p>connected</p>
				</div>
			) : (
				<div>
					<p>not connected</p>
				</div>
			)}
			<button type="button" onClick={handleClick}>
				Click me
			</button>
			<p>
				if click this button, you can see the screen click state will change
			</p>
		</div>
	);
};
