import { createFileRoute } from "@tanstack/react-router";
import { RoomDetail } from "../../../components/page/RoomDetail";
import { WithAuth } from "../../../libs/auth";

export const Route = createFileRoute("/rooms/$id/")({
	component: () => (
		<WithAuth>
			<RoomDetail />
		</WithAuth>
	),
});
