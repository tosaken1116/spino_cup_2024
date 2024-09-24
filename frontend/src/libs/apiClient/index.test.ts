import { renderHook } from "@testing-library/react-hooks";
import { describe, expect, test } from "vitest";
import { useApiClient } from ".";
describe("useApiClient", () => {
	test("generate api client correctly", () => {
		const { result } = renderHook(() => useApiClient());
		expect(result.current.room).toBeDefined();
	});
});
