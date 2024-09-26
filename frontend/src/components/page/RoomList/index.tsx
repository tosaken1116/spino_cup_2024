import { useNavigate } from "@tanstack/react-router";
import { RoomList } from "../../../domains/room/components/RoomList";
import { useAuthContext } from "../../../libs/auth/providers";
import { getBaseUrl } from "../../../libs/baseUrl";

export const RoomListPage = () => {
	const navigation = useNavigate();
	const { token } = useAuthContext();
	const handleClick = async () => {
		const res = await fetch(`${getBaseUrl()}/v1/rooms`, {
			method: "POST",
			body: JSON.stringify({ name: "test" }),
			headers: {
				Authorization: `Bearer ${token}`,
				"Content-Type": "application/json",
			},
			mode: "cors",
		});
		const json = await res.json();
		navigation({
			to: "/rooms/$id",
			params: {
				id: json.room.id,
			},
		});
	};
	return (
		<div className="w-full h-[calc(100vh-64px)] flex md:flex-row flex-col ">
			<div className="w-full flex items-center justify-center pt-8 flex-col gap-8">
				<strong className="text-xl"> 現在作成済みのルーム</strong>
				<RoomList />
			</div>
			<div className="w-full flex justify-center md:visible  items-center">
				<button
					type="button"
					onClick={handleClick}
					className="rounded-md border-2 text-xl border-orange-500 bg-orange-300 text-red-900 font-bold px-4 py-2 hover:-translate-y-1 duration-300 transition-transform"
				>
					ルームを作成！
				</button>
			</div>
		</div>
	);
};
