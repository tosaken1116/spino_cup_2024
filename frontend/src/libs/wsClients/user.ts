import { useState } from "react";
import { useRoomWSClient } from "../../generated/wsClient/room";
import type { UserPosition } from "../../generated/wsClient/room/model";
import { getBaseUrl } from "../baseUrl";

export const useRoomUserWSClient = (roomId: string) => {
	const baseUrl = getBaseUrl("ws");
	const [position, setPosition] = useState<UserPosition>({
		x: 0,
		y: 0,
		z: 0,
		color: "#000000",
		isClicked: false,
		id: "test-id",
	});
	const { connection } = useRoomWSClient({
		baseUrl: `${baseUrl}/room/${roomId}/join`,
		ChangeCurrentPosition: () => {},
		ChangeCurrentScreen: () => {},
		ChangeUserPosition: () => {},
	});

	const handleClick = () => {
		if (connection) {
			connection.send(
				JSON.stringify({
					type: "ChangeUserPosition",
					payload: {
						...position,
					},
				}),
			);
			setPosition((prev) => ({
				...prev,
				isClicked: !prev.isClicked,
			}));
		}
	};
	return { handleClick, position, isConnected: connection?.CONNECTING };
};
