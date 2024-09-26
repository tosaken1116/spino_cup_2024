import { useParams } from "@tanstack/react-router";
import { WithAuth } from "../../../libs/auth";
import { RoomDetailPage } from "./page";
export const RoomDetail = () => {
	const { id } = useParams({ strict: false });
	if (id === undefined) {
		return <div>Invalid ID</div>;
	}
	return (
		<WithAuth>
			<RoomDetailPage id={id} />
		</WithAuth>
	);
};
