import type {
	CreateRoomRequest,
	GetRoomRequest,
	UpdateRoomRequest,
} from "../../../generated/domain/room/schema";

export const roomCacheKeyGenerator = {
	createRoom: (req: CreateRoomRequest) => {
		return JSON.stringify(req);
	},
	getRoom: (req: GetRoomRequest) => {
		return JSON.stringify(req);
	},
	updateRoom: (req: UpdateRoomRequest) => {
		return JSON.stringify(req);
	},
	joinRoom: (req: GetRoomRequest) => {
		return JSON.stringify(req);
	},
} as const;
