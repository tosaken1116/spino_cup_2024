import { apiMockClient } from "../../../generated/mswHandlers";
import { generateRoomMock } from "./data";

export const roomHandlers = [
	apiMockClient.room.createRoom({ room: generateRoomMock() }),
	apiMockClient.room.getRoom({ room: generateRoomMock() }),
	apiMockClient.room.updateRoom({ room: generateRoomMock() }),
	apiMockClient.room.joinRoom({ room: generateRoomMock() }),
];
