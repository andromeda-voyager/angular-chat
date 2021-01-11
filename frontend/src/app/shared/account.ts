import { Connection, Server } from "./server";

export interface Account {
    id: number
    username: string
    password: string
    email: string
    avatarURL?: string
    connections: Connection[]
    // friends: string[]
}


export interface NewAccount {
    username: string
    password: string
    email: string
    avatarURL?: string
    code: string
}

