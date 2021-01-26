import { Server } from "./server";

export interface User {
    id: number
    username: string
    password: string
    email: string
    avatarURL: string
    // friends: string[]
}


export interface Account {
    username: string
    password: string
    email: string
    code: string
}

export interface LoginResponse {
    servers: Server[]
    user: User
}

