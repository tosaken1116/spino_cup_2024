import { createFileRoute } from "@tanstack/react-router";
import { RoomListPage } from "../../components/page/RoomList";
import { WithAuth } from "../../libs/auth";

export const Route = createFileRoute("/rooms/")({
	component: () => (
		<WithAuth>
			<RoomListPage />
		</WithAuth>
	),
});
