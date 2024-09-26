import { type ReactNode, useNavigate } from "@tanstack/react-router";
import { type FC, createContext, useContext } from "react";
import { useAuthUseCase } from "./hook";

type Props = {
	children: ReactNode;
};

const authContext = createContext<{
	token: string | null;
}>({
	token: null,
});

export const useAuthContext = () => {
	return useContext(authContext);
};
const { Provider } = authContext;

export const WithAuth: FC<Props> = ({ children }) => {
	const { loading, token } = useAuthUseCase();
	const navigate = useNavigate();
	if (loading) {
		return <p>loading...</p>;
	}
	if (token === null || token==="") {
		navigate({ to: "/" });
	}
	return <Provider value={{ token }}>{children}</Provider>;
};
