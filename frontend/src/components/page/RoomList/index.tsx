import { useNavigate } from "@tanstack/react-router";
import { RoomList } from "../../../domains/room/components/RoomList";
import { getBaseUrl } from "../../../libs/baseUrl";

export const RoomListPage = () => {
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
      to: "/rooms/$id",
      params: {
        id: json.room.id,
      },
    });
  };
  return (
    <div className="w-full h-screen flex flex-row ">

      <RoomList />
      <button type="button" onClick={handleClick}>
        click to generate room
      </button>
    </div>
  );
};
