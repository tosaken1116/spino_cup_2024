const DEV_BASE_URL = import.meta.env.VITE_DEV_BASE_URL;
const PROD_BASE_URL = import.meta.env.VITE_PROD_BASE_URL;

if (DEV_BASE_URL === undefined) {
	throw new Error("DEV_BASE_URL is not defined");
}

if (PROD_BASE_URL === undefined) {
	throw new Error("PROD_BASE_URL is not defined");
}

export const getBaseUrl = () => {
	if (process.env.NODE_ENV === "development") {
		return DEV_BASE_URL;
	}
	return PROD_BASE_URL;
};
