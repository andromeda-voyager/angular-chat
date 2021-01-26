import { Post } from "./post";

export interface Channel {
    id: number
    name: string
    posts: Post[]
    channelPermissions: ChannelPermissions[]
}

export interface ChannelPermissions {
    roleID: number
    permissions: number
}

export interface NewChannel {
    serverID: number
    name: string
    channelPermissions: ChannelPermissions[]
}