import { Link, Outlet, createRootRoute } from "@tanstack/react-router";

export const Route = createRootRoute({
	component: () => (
		<div className="w-full overflow-x-hidden  overflow-y-hidden">
			<header className="px-4">
				<Link to="/" className="[&.active]:font-bold flex flex-row items-end">
					<p className="text-blue-400 text-4xl rotate-2">お</p>
					<p className="text-red-400 text-5xl -rotate-2">絵</p>
					<p className="text-orange-400 text-3xl rotate-3">描</p>
					<p className="text-green-400 text-4xl rotate-6">き</p>
					<p className="text-2xl">の</p>
					<p className="text-blue-400 text-5xl">谷</p>
				</Link>
			</header>
			<div className="w-full h-[calc(100vh-64px)]">
				<Outlet />
			</div>
		</div>
	),
});
