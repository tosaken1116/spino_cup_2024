import { useRef } from "react";
import type { ScreenAction } from "../../../../libs/wsClients";

type Props = Omit<ScreenAction, "type">;
export const ScreenController = ({ handleChangeScreen }: Props) => {
	const ref = useRef<HTMLFormElement>(null);
	return (
		<div>
			<form
        ref={ref}
				onSubmit={() => {
					handleChangeScreen({
						width: Number(ref.current?.elements),
						height: Number(ref.current?.elements),
					});
				}}
			>
				<label>
					横幅
					<input />
				</label>
				<label>
					高さ
					<input />
				</label>
				<button type="submit">決定</button>
			</form>
		</div>
	);
};
