import { useAuthUseCase } from "../../../libs/auth";

export const TopPage = () => {
	const { login } = useAuthUseCase();
	return (
		<div>
			<button onClick={() => login("/rooms")} type="button">
				ログイン
			</button>
		</div>
	);
};
