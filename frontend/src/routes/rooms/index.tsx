import { createFileRoute } from "@tanstack/react-router";
import { RoomListPage } from "../../components/page/RoomList";

export const Route = createFileRoute("/rooms/")({
  component: RoomListPage,
});
