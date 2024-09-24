import { createFileRoute } from "@tanstack/react-router";
import { ScreenDetail } from "../../../components/page/ScreenDetail";

export const Route = createFileRoute("/screen/$id/")({
	component: () => (
		<div>
			<ScreenDetail />
		</div>
	),
});
