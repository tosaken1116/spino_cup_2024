import { useNavigate } from "@tanstack/react-router";
import { initializeApp } from "firebase/app";

import {
	GithubAuthProvider,
	type User,
	browserPopupRedirectResolver,
	browserSessionPersistence,
	initializeAuth,
	onAuthStateChanged,
	signInWithPopup,
} from "firebase/auth";
import { useEffect, useState } from "react";

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

export const useAuthUseCase = () => {
	const navigate = useNavigate();
	const [user, setUser] = useState<User | null>(null);
	const [loading, setLoading] = useState(true);
	const [token, setToken] = useState<string | null>(null);
	// useEffect(() => {
	// 	const check = async () => {
	// 		if (auth.currentUser) {
	// 			setUser(auth.currentUser);
	// 			const token = await auth.currentUser.getIdToken();
	// 			setToken(token);
	// 			setLoading(false);
	// 		}
	// 	};
	// 	check();
	// }, []);
	useEffect(() => {
		return onAuthStateChanged(auth, async (user) => {
			setUser(user);
			setLoading(false);
			const token = await user?.getIdToken();
			console.log(token);
			setToken(token ?? "");
		});
	}, []);
	const login = async (redirectPath: string) => {
		const res = await signInWithPopup(auth, provider);
		const token = await res.user.getIdToken();
		console.log(token);
		setToken(token);
		navigate({ to: redirectPath });
	};
	return {
		user,
		loading,
		token,
		login,
	};
};
