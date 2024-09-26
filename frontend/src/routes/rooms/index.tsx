import { createFileRoute } from "@tanstack/react-router";
import { RoomList } from "../../components/page/RoomList";

export const Route = createFileRoute("/rooms/")({
	component: RoomList,
});
