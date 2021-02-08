import { Message } from "./message";

export interface Channel {
    serverID: number
    id: number
    name: string
    overrides: Override[]
}

export interface Override {
    roleID: number
    permissions: number
}