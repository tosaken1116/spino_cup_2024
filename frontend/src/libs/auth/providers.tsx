import { type ReactNode, useNavigate } from "@tanstack/react-router";
import { initializeApp } from "firebase/app";
import { type FC, createContext, useContext, useEffect, useState } from "react";

import {
	GithubAuthProvider,
	browserPopupRedirectResolver,
	browserSessionPersistence,
	initializeAuth,
	onAuthStateChanged,
	signInWithPopup,
} from "firebase/auth";

type Props = {
	children: ReactNode;
};

const authContext = createContext<{
	token: string | undefined;
	isLoading: boolean;
}>({
	token: undefined,
	isLoading: true,
});

const AUTH_API_KEY = import.meta.env.VITE_AUTH_API_KEY;
const AUTH_DOMAIN = import.meta.env.VITE_AUTH_DOMAIN;

const config = {
	apiKey: AUTH_API_KEY,
	authDomain: AUTH_DOMAIN,
} as const;

const provider = new GithubAuthProvider();
const initApp = initializeApp(config);
const auth = initializeAuth(initApp, {
	persistence: browserSessionPersistence,
	popupRedirectResolver: browserPopupRedirectResolver,
});

export const useAuthContext = () => {
	return useContext(authContext);
};

export const useFirebaseLogin = () => {
	const navigate = useNavigate();
	return {
		login: async (path?: string) => {
			await signInWithPopup(auth, provider);
			if (path) {
				navigate({ to: path });
			}
		},
	};
};
const { Provider } = authContext;

export const AuthProvider: FC<Props> = ({ children }) => {
	const [token, setToken] = useState<string | undefined>(undefined);
	const [loading, setLoading] = useState(true);
	const navigate = useNavigate();
	useEffect(() => {
		const check = async () => {
			if (auth.currentUser) {
				console.log("check", auth.currentUser);
				const token = await auth.currentUser.getIdToken();
				setLoading(false);
				setToken(token);
			} else {
				const res = await signInWithPopup(auth, provider);
				setToken(await res.user.getIdToken());
				setLoading(false);
			}
		};
		check();
	}, []);

	useEffect(() => {
		return onAuthStateChanged(auth, async (user) => {
			setLoading(false);
			const token = await user?.getIdToken();
			console.log("onAuthStateChanged", token);
			setToken(token);
		});
	}, []);
	if (loading) {
		return <p>...loading</p>;
	}
	if (!token) {
		navigate({
			to: "/",
			search: {
				redirectURL: window.location.pathname,
			},
		});
	}

	return <Provider value={{ token, isLoading: loading }}>{children}</Provider>;
};
