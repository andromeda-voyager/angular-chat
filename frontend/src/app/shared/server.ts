export interface Server {
    name: string
    description: string
    serverImageUrl: string
}

export interface Invite {
    code: string
    serverID?: number
}
