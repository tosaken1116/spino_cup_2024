import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
	component: () => <div className="text-red-600">Hello /!</div>,
});
