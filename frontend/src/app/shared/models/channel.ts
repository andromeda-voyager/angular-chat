import { Post } from "./post";

export interface Channel {
    id: number
    name: string
    posts: Post[]
}

export interface ChannelPermissions {
    channelID: number
    permissions: number
}

