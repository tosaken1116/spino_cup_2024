import { apiMockClient } from "../../../generated/mswHandlers";
import { generateRoomMock, generateRoomMocks } from "./data";

export const roomHandlers = [
	apiMockClient.room.createRoom({ room: generateRoomMock() }),
	apiMockClient.room.getRoom({ room: generateRoomMock() }),
	apiMockClient.room.updateRoom({ room: generateRoomMock() }),
	apiMockClient.room.joinRoom({ room: generateRoomMock() }),
	apiMockClient.room.listRoom({ rooms: generateRoomMocks(5), total: 5 }),
];
