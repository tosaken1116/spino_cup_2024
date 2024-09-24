import { useParams } from "@tanstack/react-router";
import { ScreenDetailPage } from "./page";
export const ScreenDetail = () => {
	const { id } = useParams({ strict: false });
	if (id === undefined) {
		return <div>Invalid ID</div>;
	}
	return <ScreenDetailPage id={id} />;
};
