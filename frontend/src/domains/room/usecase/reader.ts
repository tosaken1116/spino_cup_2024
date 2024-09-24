import useSWR from "swr";
import { useRoomRepository } from "../repository";
import type { GetRoomProps } from "../types/model";
import { roomCacheKeyGenerator } from "./cache";

export const useGetRoom = (props: GetRoomProps) => {
	const repository = useRoomRepository();
	return useSWR(
		roomCacheKeyGenerator.getRoom(props),
		() => repository.getRoom(props),
		{
			suspense: true,
		},
	);
};
