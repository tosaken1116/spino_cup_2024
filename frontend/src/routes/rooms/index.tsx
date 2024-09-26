import { createFileRoute } from "@tanstack/react-router";
import { RoomListPage } from "../../components/page/RoomList";
import { AuthProvider } from "../../libs/auth";

export const Route = createFileRoute("/rooms/")({
	component: () => (
		<AuthProvider>
			<RoomListPage />
		</AuthProvider>
	),
});
