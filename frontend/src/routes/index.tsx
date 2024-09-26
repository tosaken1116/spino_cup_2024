import { createFileRoute } from "@tanstack/react-router";
import { TopPage } from "../components/page/TopPage";

export const Route = createFileRoute("/")({
	component: TopPage,
});
