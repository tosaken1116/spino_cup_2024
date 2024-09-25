import { useCallback, useMemo, useState } from "react";
import { useRoomWSClient } from "../../generated/wsClient/room";
import type { UserPosition } from "../../generated/wsClient/room/model";
import { getBaseUrl } from "../baseUrl";

export type UserAction = {
	type: "user";
	color: string;
	isClicked: boolean;
	handleChangePointerPosition: (
		payload: Omit<UserPosition, "id" | "color" | "isClicked">,
	) => void;
	handleChangePointerColor: (color: string) => void;
	handleClickPointer: (isClicked: boolean) => void;
};

export type ScreenAction = {
	type: "screen";
	positions: UserPosition[];
	handleChangeScreen: (payload: { width: number; height: number }) => void;
};

export const useRoomUserWSClient = (
	roomId: string,
): UserAction | ScreenAction | null => {
	const baseUrl = getBaseUrl("ws");
	const [userId, setUserId] = useState<string | null>(null);
	const [ownerId, setOwnerId] = useState<string | null>(null);
	const [position, setPosition] = useState<Omit<UserPosition, "id">>({
		x: 0,
		y: 0,
		color: "#000000",
		isClicked: false,
	});
	const [positions, setPositions] = useState<UserPosition[]>([]);
	const { handleChangeCurrentPosition, handleChangeCurrentScreen } =
		useRoomWSClient({
			baseUrl: `${baseUrl}/rooms/${roomId}/join`,
			ChangeUserPosition: (payload) => {
				const ids = positions.map((position) => position.id);
				setPositions((prev) =>
					prev.map((position) =>
						ids.includes(position.id) ? { ...position, ...payload } : position,
					),
				);
			},
			JoinRoom: ({ userId, ownerId }) => {
				setUserId(userId);
				setOwnerId(ownerId);
			},
		});

	const handleChangePointerPosition = useCallback(
		(payload: Omit<UserPosition, "id" | "color" | "isClicked">) => {
			if (!userId) return;
			handleChangeCurrentPosition({ id: userId, ...position, ...payload });
			setPosition({ ...position, ...payload });
		},
		[handleChangeCurrentPosition, userId, position],
	);
	const handleChangePointerColor = useCallback(
		(color: string) => {
			if (!userId) return;
			handleChangeCurrentPosition({ id: userId, ...position, color });
			setPosition({ ...position, color });
		},
		[handleChangeCurrentPosition, userId, position],
	);

	const handleClickPointer = useCallback(
		(isClicked: boolean) => {
			if (!userId) return;
			handleChangeCurrentPosition({ id: userId, ...position, isClicked });
			setPosition({ ...position, isClicked });
		},
		[handleChangeCurrentPosition, userId, position],
	);

	const userActions = useMemo<UserAction>(
		() => ({
			type: "user",
			color: position.color,
			isClicked: position.isClicked,
			handleChangePointerPosition,
			handleChangePointerColor,
			handleClickPointer,
		}),
		[
			handleChangePointerPosition,
			handleChangePointerColor,
			handleClickPointer,
			position,
		],
	);

	const screenActions = useMemo<ScreenAction>(
		() => ({
			type: "screen",
			positions,
			handleChangeScreen: handleChangeCurrentScreen,
		}),
		[handleChangeCurrentScreen, positions],
	);

	if (userId !== null && ownerId !== null && userId !== ownerId) {
		return userActions;
	}
	if (userId !== null && ownerId !== null && userId === ownerId) {
		return screenActions;
	}
	return null;
};
