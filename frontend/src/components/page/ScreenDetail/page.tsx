import { useScreenUserWSClient } from "../../../libs/wsClients";

type Props = {
	id: string;
};
export const ScreenDetailPage = (props: Props) => {
	const { positions } = useScreenUserWSClient(props.id);
	if (positions.length === 0) {
		return <div>connection is 0</div>;
	}
	return (
		<div>
			{positions.map((position) => {
				return (
					<div key={position.id}>
						<p>{position.id}</p>
						<p>{position.x}</p>
						<p>{position.y}</p>
						<p>isClicked:{position.isClicked}</p>
					</div>
				);
			})}
		</div>
	);
};
