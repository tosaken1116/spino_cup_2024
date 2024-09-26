import { getRouteApi } from "@tanstack/react-router";
import { useFirebaseLogin } from "../../../libs/auth/providers";
import { Route } from "../../../routes/__root";
import { GithubLogo } from "../../icons/GithubLogo";
import { BruchIcon } from "../../icons/brush";

export const TopPage = () => {
	const { login } = useFirebaseLogin();
	const routeApi = getRouteApi(Route.fullPath);
	const filters = routeApi.useSearch();
	return (
		<div className="w-full h-full flex flex-col md:flex-row">
			<div className="w-full h-full flex justify-center items-center">
				<div
					className="animate-jump"
					style={{ animationDuration: "300ms", animationDelay: "100ms" }}
				>
					<p className="text-blue-400 md:text-8xl text-4xl font-bold rotate-2">お</p>
				</div>

				<div
					className=" animate-jump"
					style={{ animationDuration: "300ms", animationDelay: "150ms" }}
				>
					<p className="text-red-400 md:text-9xl text-5xl font-bold -rotate-2">絵</p>
				</div>

				<div
					className=" animate-jump"
					style={{ animationDuration: "300ms", animationDelay: "200ms" }}
				>
					<p className="text-orange-400 md:text-7xl text-3xl font-bold rotate-3">描</p>
				</div>

				<div
					className=" animate-jump"
					style={{ animationDuration: "300ms", animationDelay: "250ms" }}
				>
					<p className="text-green-400 md:text-8xl text-4xl font-bold rotate-6">き</p>
				</div>

				<div
					className=" animate-jump"
					style={{ animationDuration: "300ms", animationDelay: "300ms" }}
				>
					<p className="md:text-5xl text-4xl font-bold">の</p>
				</div>
				<div
					className=" animate-jump"
					style={{ animationDuration: "300ms", animationDelay: "350ms" }}
				>
					<p className="text-blue-400 md:text-9xl text-4xl  font-bold relative">
						谷
						<BruchIcon
							className="absolute -left-3 top-0"
							width={72}
							height={72}
							fill="white"
						/>
					</p>
				</div>
			</div>
			<div className="w-full h-full flex justify-center items-center">
				<button
					onClick={() => login(filters.redirectURL ?? "/rooms")}
					type="button"
					className="flex flex-row md:text-2xl text-2xl gap-2 rounded-xl rounded4xl border-slate-500 border-2 px-4 py-2 items-center"
				>
					<GithubLogo fill="black" />
					Githubでログイン
				</button>
			</div>
		</div>
	);
};
