import { Message } from "./message";

export interface Channel {
    id: number
    name: string
    messages: Message[]
}

export interface ChannelPermissions {
    channelID: number
    permissions: number
}

