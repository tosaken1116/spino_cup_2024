import { useNavigate } from "@tanstack/react-router";
import { useRoomUsecase } from "../../../domains/room/usecase";
import { getBaseUrl } from "../../../libs/baseUrl";

export const RoomTop = () => {
	const navigation = useNavigate();
	const handleClick = async () => {
		const res = await fetch(`${getBaseUrl()}/v1/rooms`, {
			method: "POST",
			body: JSON.stringify({ name: "test" }),
			headers: {
				"Content-Type": "application/json",
			},
			mode: "cors",
		});
		const json = await res.json();
		navigation({
			to: "/room/$id",
			params: {
				id: json.room.id,
			},
		});
	};
	return (
		<button type="button" onClick={handleClick}>
			click to generate room
		</button>
	);
};
