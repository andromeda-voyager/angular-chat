import { Channel } from "./channel";
import { User } from "./user";

export interface Server {
    id: number
    name: string
    image: string
    description: string 
    role: Role
    roles: Role[]
    channels: Channel[]
}

export interface NewServer {
    name:string
    description:string
}
    // updateRole(r: Role) {
    //     let role = this.roles.find(function (role) {
    //         return role.id == r.id;
    //     });
    //     if (role) {
    //         role = r;
    //     } else {
    //         this.roles.push(r);
    //     }
    // }

export interface Role {
    id?: number
    name: string
    serverPermissions: number
    ChannelPermissions: ChannelPermissions[]
}

export interface ChannelPermissions {
	channelID: number   
	value:     number 
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
    servers?: Server[]
    serverID: number
    channelID: number;
    channel?: Channel
    role?: Role
}