const DEV_BASE_URL = import.meta.env.VITE_DEV_API_DOMAIN;
const PROD_BASE_URL = import.meta.env.VITE_PROD_API_DOMAIN;
const IS_DEV = import.meta.env.NODE_ENV === "development";
if (IS_DEV && DEV_BASE_URL === undefined) {
	throw new Error("DEV_BASE_URL is not defined");
}

if (!IS_DEV && PROD_BASE_URL === undefined) {
	throw new Error("PROD_BASE_URL is not defined");
}

export const getBaseUrl = (type: "ws" | "http" = "http"): string => {
	if (process.env.NODE_ENV === "development") {
		return `${type}://${DEV_BASE_URL}`;
	}
	return `${type}s://${PROD_BASE_URL}`;
};
