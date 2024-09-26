import { getRouteApi } from "@tanstack/react-router";
import { useFirebaseLogin } from "../../../libs/auth/providers";
import { Route } from "../../../routes/__root";

export const TopPage = () => {
	const { login } = useFirebaseLogin();
	const routeApi = getRouteApi(Route.fullPath);
	const filters = routeApi.useSearch();
	return (
		<div>
			<button
				onClick={() => login(filters.redirectURL ?? "/rooms")}
				type="button"
			>
				ログイン
			</button>
		</div>
	);
};
