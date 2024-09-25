import clsx from "clsx";

type Props = {
	className?: string;
} & React.SVGProps<SVGSVGElement>;
export const SquareArrowUpLeft = ({ className, ...props }: Props) => (
	<svg
		width="24"
		height="24"
		viewBox="0 0 24 24"
		fill="none"
		stroke="currentColor"
		strokeWidth="2"
		strokeLinecap="round"
		strokeLinejoin="round"
		className={clsx("lucide lucide-square-arrow-up-left", className)}
		{...props}
	>
		<title>SquareArrowUpLeft</title>

		<rect width="18" height="18" x="3" y="3" rx="2" />
		<path d="M8 16V8h8" />
		<path d="M16 16 8 8" />
	</svg>
);
