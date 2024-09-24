import { describe, expect, test } from "vitest";
import { calculateScreenPosition } from ".";

describe("calculateScreenPosition", () => {
	test("should calculate screen position from front", () => {
		const max = { alpha: 90, beta: 90 };
		const min = { alpha: -90, beta: -90 };
		const current = { alpha: 0, beta: 0 };
		const center = { alpha: 0, beta: 0 };
		const screenSize = { width: 1920, height: 1080 };
		const result = calculateScreenPosition({
			max,
			min,
			current,
			center,
			screenSize,
		});

		expect(result).toEqual({ x: 960, y: 540 });
	});
	test("should calculate screen position from right", () => {
		const max = { alpha: 45, beta: 45 };
		const min = { alpha: -45, beta: -45 };
		const current = { alpha: 15, beta: 0 };
		const center = { alpha: 0, beta: 0 };
		const screenSize = { width: 1920, height: 1080 };
		const result = calculateScreenPosition({
			max,
			min,
			current,
			center,
			screenSize,
		});

		expect(result).toEqual({ x: (1920 / 6) * 4, y: 540 });
	});
	test("should calculate screen position from left", () => {
		const max = { alpha: 45, beta: 45 };
		const min = { alpha: -45, beta: -45 };
		const current = { alpha: -15, beta: 0 };
		const center = { alpha: 0, beta: 0 };
		const screenSize = { width: 1920, height: 1080 };
		const result = calculateScreenPosition({
			max,
			min,
			current,
			center,
			screenSize,
		});

		expect(result).toEqual({ x: (1920 / 6) * 2, y: 540 });
	});

	test("should calculate screen position from top", () => {
		const max = { alpha: 45, beta: 45 };
		const min = { alpha: -45, beta: -45 };
		const current = { alpha: 0, beta: 15 };
		const center = { alpha: 0, beta: 0 };
		const screenSize = { width: 1920, height: 1080 };
		const result = calculateScreenPosition({
			max,
			min,
			current,
			center,
			screenSize,
		});

		expect(result).toEqual({ x: 960, y: (1080 / 6) * 4 });
	});
	test("should calculate screen position from bottom", () => {
		const max = { alpha: 45, beta: 45 };
		const min = { alpha: -45, beta: -45 };
		const current = { alpha: 0, beta: -15 };
		const center = { alpha: 0, beta: 0 };
		const screenSize = { width: 1920, height: 1080 };
		const result = calculateScreenPosition({
			max,
			min,
			current,
			center,
			screenSize,
		});

		expect(result).toEqual({ x: 960, y: (1080 / 6) * 2 });
	});
});
