// Code generated by protoc-gen-ts. DO NOT EDIT.
import { http, HttpResponse } from 'msw'
import { getBaseUrl } from "../../libs/baseUrl";
import type * as RoomSchema from '../apiclient/domain/room/schema';

export const apiMockClientBase = (baseUrl: string) => ({
  room: {
		createRoom: (data: RoomSchema.CreateRoomResponse) =>
			http.post(`${baseUrl}/v1/rooms`, () => {
				return HttpResponse.json(data)
			}),
		getRoom: (data: RoomSchema.GetRoomResponse) =>
			http.get(`${baseUrl}/v1/rooms/:id`, () => {
				return HttpResponse.json(data)
			}),
		updateRoom: (data: RoomSchema.UpdateRoomResponse) =>
			http.put(`${baseUrl}/v1/rooms/:id`, () => {
				return HttpResponse.json(data)
			}),
		joinRoom: (data: RoomSchema.JoinRoomResponse) =>
			http.post(`${baseUrl}/v1/rooms/:id/join`, () => {
				return HttpResponse.json(data)
			}),
		listRoom: (data: RoomSchema.ListRoomResponse) =>
			http.get(`${baseUrl}/v1/rooms`, () => {
				return HttpResponse.json(data)
			}),
  },
});
export const apiMockClient = apiMockClientBase(getBaseUrl())
