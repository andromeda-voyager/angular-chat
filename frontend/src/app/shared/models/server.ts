import { Channel } from "./channel";

export interface Server {
    id: number
    name: string
    description: string
    image: string
    roles: Role[]
    channels: Channel[]
}

export interface Role {
    id?: number
    name: string
    serverPermissions: number
}

export interface Invite {
    code: string
}

// ServerRequest .
export interface ServerRequest {
    serverID: number
    Channel?: Channel
    Role?: Role
    Roles?: Role[]
}