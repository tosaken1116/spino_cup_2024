import { describe, expect, test } from "vitest";
import { getBaseUrl } from ".";

describe("getBaseUrl", () => {
	test("should return PROD_BASE_URL", () => {
		const baseUrl = getBaseUrl();
		expect(baseUrl).toBe("https://api.example.com");
	});
});
