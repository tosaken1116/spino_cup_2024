import { Link, Outlet, createRootRoute } from "@tanstack/react-router";

export const Route = createRootRoute({
	component: () => (
		<div>
			<header className="h-12">
				<Link to="/" className="[&.active]:font-bold">
					お絵描きの谷
				</Link>
			</header>
			<div  className="w-full h-[calc(100vh-48px)]">

			<Outlet />
			</div>
		</div>
	),
});
