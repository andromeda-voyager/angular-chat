export interface Post {
    accountID: number
    text: string
    media: string
    timePosted: Date
}

export interface NewPost {
    channelID: number
    serverID: number
    text: string
    media: string
}