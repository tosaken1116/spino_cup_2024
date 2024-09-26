import { Link } from "@tanstack/react-router";
import { Suspense } from "react";
import { useListRoom } from "../../usecase/reader";

export const RoomList = () => {
	return (
		<Suspense fallback={<p>...loading</p>}>
			<RoomListRender />
		</Suspense>
	);
};

const RoomListRender = () => {
	const { data } = useListRoom();
	return (
		<ul className="gap-2 flex flex-col py-4 h-full overflow-y-scroll px-2">
			{data.rooms.map((room, index) => (
				<Link
					as="li"
					to={"/rooms/$id"}
					params={{ id: room.id }}
					key={room.id}
					className="animate-fade-in rounded-sm rotate-6 relative text-black w-64 bg-yellow-100 after:contents-[''] after:bg-black after:w-4 after:h-4 after:rounded-full after:absolute after:left-4 after:top-1/2 after:-translate-x-1/2 after:-translate-y-1/2 hover:-translate-y-2 duration-200"
					style={{
						animationDelay: `${index * 100}ms`,
					}}
				>
					<p>{room.name}</p>
					<div className="flex flex-row gap-2">
						<p>{room.ownerId}</p>
					</div>
				</Link>
			))}
		</ul>
	);
};
