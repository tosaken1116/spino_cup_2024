import { useCallback, useMemo, useRef, useState } from "react";
import { useRoomWSClient } from "../../generated/wsClient/room";
import type {
	ScreenSize,
	UserPosition,
} from "../../generated/wsClient/room/model";
import { getBaseUrl } from "../baseUrl";

export type UserAction = {
	type: "user";
	color: string;
	isClicked: boolean;
	screenSize: ScreenSize;
	handleChangePointerPosition: (payload: Pick<UserPosition, "x" | "y">) => void;
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
	const position = useRef<Omit<UserPosition, "id">>({
		x: 0,
		y: 0,
		color: "#000000",
		isClicked: false,
		penSize: 1,
	});
	const [screen, setScreen] = useState({ width: 0, height: 0 });
	const [positions, setPositions] = useState<UserPosition[]>([]);
	const { handleChangeCurrentPosition, handleChangeCurrentScreen } =
		useRoomWSClient({
			baseUrl: `${baseUrl}/rooms/${roomId}/join`,
			ChangeScreenSizeToUser: (payload) => {
				setScreen(payload);
			},
			ChangeUserPosition: (payload) => {
				setPositions((prev) => {
					const map = new Map<string, UserPosition>();
					for (const item of prev) {
						map.set(item.id, item);
					}

					for (const item of payload) {
						map.set(item.id, item);
					}
					return Array.from(map.values());
				});
			},
			JoinRoom: ({ userId, ownerId }) => {
				setUserId(userId);
				setOwnerId(ownerId);
			},
		});

	const handleChangePointerPosition = useCallback(
		(payload: Pick<UserPosition, "x" | "y">) => {
			if (!userId) return;
			handleChangeCurrentPosition({
				id: userId,
				...position.current,
				...payload,
			});
			position.current = { ...position.current, ...payload };
		},
		[handleChangeCurrentPosition, userId],
	);
	const handleChangePointerColor = useCallback(
		(color: string) => {
			if (!userId) return;
			handleChangeCurrentPosition({ id: userId, ...position.current, color });
			position.current = { ...position.current, color };
		},
		[handleChangeCurrentPosition, userId],
	);

	const handleClickPointer = useCallback(
		(isClicked: boolean) => {
			if (!userId) return;
			handleChangeCurrentPosition({
				id: userId,
				...position.current,
				isClicked,
			});
			position.current = { ...position.current, isClicked };
		},
		[handleChangeCurrentPosition, userId],
	);

	const userActions = useMemo<UserAction>(
		() => ({
			type: "user",
			screenSize: screen,
			color: position.current.color,
			isClicked: position.current.isClicked,
			handleChangePointerPosition,
			handleChangePointerColor,
			handleClickPointer,
		}),
		[
			handleChangePointerPosition,
			handleChangePointerColor,
			handleClickPointer,
			screen,
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
