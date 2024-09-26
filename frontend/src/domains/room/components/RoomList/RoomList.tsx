import { Link } from "@tanstack/react-router";
import { Suspense } from "react";
import { useListRoom } from "../../usecase/reader";

export const RoomList = () => {
  return (
    <Suspense fallback={<p>...loading</p>}>
      <RoomListRender />
    </Suspense>
  );
};

const RoomListRender = () => {
  const { data } = useListRoom();
  return (
    <ul className="gap-2 flex  flex-col py-8 h-full overflow-y-scroll px-2">
      {data.rooms.map((room, index) => (
        <Link
          as="li"
          to={"/rooms/$id"}
          params={{ id: room.id }}
          key={room.id}
          className="animate-fade-in px-8 rounded-sm rotate-6 relative text-black w-72 bg-yellow-100 after:contents-[''] after:bg-black after:w-4 after:h-4 after:rounded-full after:absolute after:left-4 after:top-1/2 after:-translate-x-1/2 after:-translate-y-1/2 hover:-translate-y-2 duration-200"
          style={{
            animationDelay: `${index * 100}ms`,
          }}
        >
          <p className=" border-b-4 border-dotted border-spacing-3">
            {room.name}
          </p>
          <div className="flex flex-row gap-2 items-center">
            {/* @ts-ignore 時間ない */}
            <img width={32} height={32} className="rounded-full" src={room.owner?.avatarUrl }/>
            {/* @ts-ignore 時間ない */}
            <p className="font-bold">{room.owner?.name }</p>
          </div>
        </Link>
      ))}
    </ul>
  );
};
