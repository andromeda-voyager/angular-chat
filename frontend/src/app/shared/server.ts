export interface Server {
    id: number
    name: string
    description: string
    imageURL: string
    posts: Post[]
}

export interface NewServer {
    name: string
    description: string
    imageURL: string
}

export interface Connection {
    server: Server
    userID: number
    alias: string
    permissions: number
}

export interface Invite {
    code: string
    serverID?: number
}

export interface Post {
    serverID: number
    userID: number
    text: string
    medialURL: string
    timePosted: Date
}

export interface NewPost {
    serverID: number
    text: string
    mediaURL: string
}