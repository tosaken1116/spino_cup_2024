import clsx from "clsx";

type Props = {
	className?: string;
} & React.SVGProps<SVGSVGElement>;
export const ImageDown = ({ className, ...props }: Props) => (
	<svg
		width="24"
		height="24"
		viewBox="0 0 24 24"
		fill="none"
		stroke="currentColor"
		strokeWidth="2"
		strokeLinecap="round"
		strokeLinejoin="round"
		className={clsx("lucide lucide-square-arrow-down-right", className)}
		{...props}
	>
		<title> ImageDown</title>

		<path d="M10.3 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10l-3.1-3.1a2 2 0 0 0-2.814.014L6 21" />
		<path d="m14 19 3 3v-5.5" />
		<path d="m17 22 3-3" />
		<circle cx="9" cy="9" r="2" />
	</svg>
);
