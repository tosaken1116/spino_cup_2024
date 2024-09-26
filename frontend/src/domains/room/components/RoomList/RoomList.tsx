import { Suspense } from "react";
import { useListRoom } from "../../usecase/reader";

export const RoomList = () => {
	return (
		<Suspense>
			<RoomListRender />
		</Suspense>
	);
};

const RoomListRender = () => {
	const { data } = useListRoom();
	return (
		<ul>
			{data.rooms.map((room) => (
				<li key={room.id}>
					<p>{room.id}</p>
					<p>{room.name}</p>
				</li>
			))}
		</ul>
	);
};
