import { createFileRoute } from "@tanstack/react-router";
import { RoomDetail } from "../../../components/page/RoomDetail";

export const Route = createFileRoute("/rooms/$id/")({
	component: () => (
		<div>
			<RoomDetail />
		</div>
	),
});
