// Code generated by protoc-gen-ts. DO NOT EDIT.
// source: proto/room/rpc/room.proto
import { Room } from '../../domain/room/model';
export type CreateRoomRequest = {
  name: string;
  description: string;
};
export type CreateRoomResponse = {
  room: Room;
};
export type GetRoomRequest = {
  id: string;
};
export type GetRoomResponse = {
  room: Room;
};
export type UpdateRoomRequest = {
  id: string;
  name: string;
  description: string;
};
export type UpdateRoomResponse = {
  room: Room;
};
export type JoinRoomRequest = {
  id: string;
};
export type JoinRoomResponse = {
  room: Room;
};
