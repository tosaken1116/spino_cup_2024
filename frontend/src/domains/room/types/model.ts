export type CreateRoomProps = {
	name: string;
	description: string;
};

export type UpdateRoomProps = {
	id: string;
	name: string;
	description: string;
};

export type JoinRoomProps = {
	id: string;
};

export type GetRoomProps = {
	id: string;
};
