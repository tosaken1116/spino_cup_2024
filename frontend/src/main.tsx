import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { App } from "./App.tsx";
import "./index.css";

async function enableMocking() {
	if (
		import.meta.env.NODE_ENV !== "development" ||
		import.meta.env.USE_MOCK !== "true"
	) {
		return;
	}

	const { worker } = await import("./mocks/browser");

	// `worker.start()` returns a Promise that resolves
	// once the Service Worker is up and ready to intercept requests.
	return worker.start();
}
enableMocking().then(() => {
	// biome-ignore lint/style/noNonNullAssertion: 存在するからnullじゃない
	createRoot(document.getElementById("root")!).render(
		<StrictMode>
			<App />
		</StrictMode>,
	);
});
