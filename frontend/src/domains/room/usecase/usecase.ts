import { useCallback } from "react";
import { mutate } from "swr";
import { useRoomRepository } from "../repository";
import type {
	CreateRoomProps,
	JoinRoomProps,
	UpdateRoomProps,
} from "../types/model";
import { roomCacheKeyGenerator } from "./cache";

export const useRoomUsecase = () => {
	const repository = useRoomRepository();
	const createRoom = useCallback(
		async (props: CreateRoomProps) => {
			const res = await repository.createRoom(props);
			await Promise.all([
				mutate(roomCacheKeyGenerator.getRoom({ id: res.room.id })),
				mutate(roomCacheKeyGenerator.listRoom()),
			]);
		},
		[repository.createRoom],
	);

	const updateRoom = useCallback(
		async (props: UpdateRoomProps) => {
			const res = await repository.updateRoom(props);
			await mutate(roomCacheKeyGenerator.getRoom({ id: res.room.id }));
		},
		[repository.updateRoom],
	);

	const joinRoom = useCallback(
		async (props: JoinRoomProps) => {
			const res = await repository.joinRoom(props);
			await mutate(roomCacheKeyGenerator.getRoom({ id: res.room.id }));
		},
		[repository.joinRoom],
	);
	return {
		createRoom,
		updateRoom,
		joinRoom,
	};
};
