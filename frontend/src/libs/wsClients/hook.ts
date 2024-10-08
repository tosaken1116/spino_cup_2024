import { useCallback, useMemo, useRef, useState } from "react";
import { useRoomWSClient } from "../../generated/wsClient/room";
import type {
	ScreenSize,
	UserPosition,
	UserPositionToScreen,
} from "../../generated/wsClient/room/model";
import { useAuthContext } from "../auth/providers";
import { getBaseUrl } from "../baseUrl";

export type UserAction = {
	type: "user";
	color: string;
	isClicked: boolean;
	screenSize: ScreenSize;
	handleChangePointerPosition: (payload: Pick<UserPosition, "x" | "y">) => void;
	handleChangePointerColor: (color: string) => void;
	handleClickPointer: (isClicked: boolean) => void;
	handleChangePenSize: (penSize: number) => void;
};

export type ScreenAction = {
	type: "screen";
	positions: UserPositionToScreen[];
	handleChangeScreen: (payload: { width: number; height: number }) => void;
};

export const useRoomUserWSClient = (
	roomId: string,
): UserAction | ScreenAction | null => {
	const baseUrl = getBaseUrl("ws");
	const [userId, setUserId] = useState<string | null>(null);
	const [ownerId, setOwnerId] = useState<string | null>(null);
	const { token } = useAuthContext();
	const position = useRef<Omit<UserPosition, "id">>({
		x: 0,
		y: 0,
		color: "#000000",
		isClicked: false,
		penSize: 1,
	});
	const [screen, setScreen] = useState({ width: 0, height: 0 });
	const [positions, setPositions] = useState<UserPositionToScreen[]>([]);
	const { handleChangeCurrentPosition, handleChangeCurrentScreen } =
		useRoomWSClient({
			baseUrl: `${baseUrl}/rooms/${roomId}/join?token=${token}`,
			ChangeScreenSizeToUser: (payload) => {
				setScreen(payload);
			},
			ChangeUserPosition: (payload) => {
				setPositions((prev) => {
					const map = new Map<string, UserPositionToScreen>();
					for (const item of prev) {
						map.set(item.user.id, item);
					}

					for (const item of payload) {
						map.set(item.user.id, item);
					}
					return Array.from(map.values());
				});
			},
			JoinRoom: ({ userId, ownerId }) => {
				setUserId(userId);
				setOwnerId(ownerId);
			},
			LeaveRoom: ({ id }) => {
				setPositions((prev) => prev.filter((item) => item.user.id !== id));
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
	const handleChangePenSize = useCallback(
		(penSize: number) => {
			if (!userId) return;
			handleChangeCurrentPosition({
				id: userId,
				...position.current,
				penSize,
			});
			position.current = { ...position.current, penSize };
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
			handleChangePenSize,
		}),
		[
			handleChangePointerPosition,
			handleChangePointerColor,
			handleClickPointer,
			handleChangePenSize,
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
