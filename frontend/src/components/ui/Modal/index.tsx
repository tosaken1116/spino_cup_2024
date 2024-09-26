import type { ReactNode } from "react";

type Props = {
	isOpen: boolean;
	onClose: () => void;
	children: ReactNode;
};
const Modal = ({ isOpen, onClose, children }: Props) => {
	if (!isOpen) return null;

	return (
		// biome-ignore lint/a11y/useKeyWithClickEvents: <explanation>
		<div
			className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50"
			onClick={onClose}
		>
			{/* biome-ignore lint/a11y/useKeyWithClickEvents: <explanation> */}
			<div
				className="bg-white p-6 rounded-lg shadow-lg relative"
				onClick={(e) => e.stopPropagation()}
			>
				<button
					className="absolute top-2 right-2 text-gray-600 hover:text-gray-800"
					onClick={onClose}
					type="button"
				>
					×
				</button>
				{children}
			</div>
		</div>
	);
};

export default Modal;
