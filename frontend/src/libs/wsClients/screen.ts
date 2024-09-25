import { useState } from "react";
import { useRoomWSClient } from "../../generated/wsClient/room";
import type { UserPosition } from "../../generated/wsClient/room/model";
import { getBaseUrl } from "../baseUrl";

export const useScreenUserWSClient = (roomId: string) => {
	const baseUrl = getBaseUrl("ws");
	const [positions, setPositions] = useState<UserPosition[]>([]);
	useRoomWSClient({
		baseUrl: `${baseUrl}/rooms/${roomId}`,
		ChangeCurrentPosition: () => {},
		ChangeCurrentScreen: () => {},
		ChangeUserPosition: (data) => {
			setPositions(data.userPositions);
		},
	});

	return { positions };
};
