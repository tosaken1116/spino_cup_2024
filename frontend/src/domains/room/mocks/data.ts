import { v4 as generateId } from "uuid";
import type { Room } from "../../../generated/apiclient/domain/room/model";
export const generateRoomMock = (props?: Partial<Room>): Room => {
	return {
		id: generateId(),
		name: "ルーム1",
		description: "ルーム1の説明",
		ownerId: generateId(),
		memberIds: Array.from({ length: 5 }, () => generateId()),
		...props,
	};
};

export const generateRoomMocks = (count: number): Room[] => {
	return Array.from({ length: count }, () => generateRoomMock());
};
