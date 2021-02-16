import { Channel } from "./channel";
import { Message } from "./message";
import { Role, Server } from "./server";

export interface Update {
    type: UpdateType
    event: UpdateEvent
    server?: Server
    channel?: Channel
    role?: Role
    message?: Message
}

export enum UpdateEvent {
    MESSAGE = "Message",
    CHANNEL = "Channel",
    ROLE = "Role",
    MEMBER = "Member"
}

export enum UpdateType {
    DELETE = "Delete",
    MODIFY = "Modify",
    NEW = "New",
}
