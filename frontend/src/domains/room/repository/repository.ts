import { useMemo } from "react";
import type { ApiClient } from "../../../generated/apiclient";
import type {
	CreateRoomRequest,
	GetRoomRequest,
	UpdateRoomRequest,
} from "../../../generated/apiclient/domain/room/schema";
import { useApiClient } from "../../../libs/apiClient";

export const RoomRepository = {};

export const useRoomRepository = () => {
	const apiClient = useApiClient();
	const repository = useMemo(
		() => createRoomRepository(apiClient),
		[apiClient],
	);

	return repository;
};

const createRoomRepository = (apiClient: ApiClient) => ({
	createRoom: async (req: CreateRoomRequest) => {
		const res = await apiClient.room.createRoom(req);
		return res;
	},
	getRoom: async (req: GetRoomRequest) => {
		const res = await apiClient.room.getRoom(req);
		return res;
	},
	updateRoom: async (req: UpdateRoomRequest) => {
		const res = await apiClient.room.updateRoom(req);
		return res;
	},
	joinRoom: async (req: GetRoomRequest) => {
		const res = await apiClient.room.joinRoom(req);
		return res;
	},
	listRoom: async () => {
		const res = await apiClient.room.listRoom();
		return res;
	},
});
