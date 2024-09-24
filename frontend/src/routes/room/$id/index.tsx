import { createFileRoute } from "@tanstack/react-router";
import { RoomDetail } from "../../../components/page/RoomDetail";

export const Route = createFileRoute("/room/$id/")({
	component: () => (
		<div>
			<RoomDetail />
		</div>
	),
});
