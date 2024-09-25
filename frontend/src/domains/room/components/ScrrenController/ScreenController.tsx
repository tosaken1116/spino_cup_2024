import { useRef } from "react";
import type { ScreenAction } from "../../../../libs/wsClients";

type Props = Omit<ScreenAction, "type">;
export const ScreenController = ({ handleChangeScreen, positions }: Props) => {
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
			<div>
				{positions.map((position) => (
					<div key={position.id}>
						<p>{position.id}</p>
						<p className="text-xl">{position.isClicked ? "🖕" : "🫵"}</p>
						<div
							className="w-4 h-4 rounded-full"
							style={{
								backgroundColor: position.color,
							}}
						/>
					</div>
				))}
			</div>
		</div>
	);
};
