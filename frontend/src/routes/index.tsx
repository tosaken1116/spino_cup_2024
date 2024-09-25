import { createFileRoute } from "@tanstack/react-router";
import { RoomTop } from "../components/page/RoomTop";

export const Route = createFileRoute("/")({
	component: RoomTop,
});
