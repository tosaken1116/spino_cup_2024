import { createFileRoute } from "@tanstack/react-router";
import { Suspense } from "react";
import { useGetRoom } from "../../domains/room/usecase";

export const Route = createFileRoute("/room/")({
	component: () => (
		<Suspense fallback="loading">
			<Inner />
		</Suspense>
	),
});

const Inner = () => {
	const { data } = useGetRoom({ id: "1" });
	return <div>{data.room.description}</div>;
};
