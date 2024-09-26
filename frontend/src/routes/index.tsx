import { createFileRoute } from "@tanstack/react-router";
import { TopPage } from "../components/page/TopPage";

type ProductSearch = {
	redirectURL?: string;
};

export const Route = createFileRoute("/")({
	component: TopPage,
	validateSearch: (search: Record<string, unknown>): ProductSearch => {
		return {
			redirectURL: search.redirectURL as string,
		};
	},
});
