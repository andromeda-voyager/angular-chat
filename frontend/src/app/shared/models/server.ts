import { Channel } from "./channel";
import { Message } from "./message";

export enum UpdateType {
    MESSAGE = "Message",
    CHANNEL = "Channel",
    ROLE = "Role",
    MEMBER = "Member"
}

export enum UpdateEvent {
    DELETE = "Delete",
    MODIFY = "Modify",
    NEW = "New",
}

export interface Server {
    id: number
    name: string
    image: string
    description: string 
    role: Role
    roles: Role[]
    channels: Channel[]
    members: Member[]
}

export interface Member {
    id: number
    serverID: number
    alias: string
    avatar: string
    role: Role
}

export interface NewServer {
    name:string
    description:string
}

export interface Role {
    id: number
    name: string
    permissions: number
    channelOverrides: ChannelOverride[]
}

export interface ChannelOverride {
	channelID: number   
	permissions:     number 
}

export interface Invite {
    code: string
}

export interface ServerRequest {
    serverID: number
    channel?: Channel
    role?: Role
    roles?: Role[]
}

export interface Update {
    type: UpdateType
    event: UpdateEvent
    server?: Server[]
    channel?: Channel
    role?: Role
    messages?: Message[]
}