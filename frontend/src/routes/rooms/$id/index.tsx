import { createFileRoute } from "@tanstack/react-router";
import { RoomDetail } from "../../../components/page/RoomDetail";
import { AuthProvider } from "../../../libs/auth";

export const Route = createFileRoute("/rooms/$id/")({
	component: () => (
		<AuthProvider>
			<RoomDetail />
		</AuthProvider>
	),
});
