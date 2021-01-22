export interface Server {
    id: number
    name: string
    description: string
    image: string
    posts: Post[]
}

export interface Channel {
	id:          number
	name:        string  
	permissions: number   
	posts:      Post[] 
}

export interface Role {
	// id:          number
	name:        string  
	serverPermissions: number   
}

export interface Invite {
    code: string
}

export interface Post {
    accountID: number
    text: string
    media: string
    timePosted: Date
}

export interface NewPost {
    channelID: number
    text: string
    media: string
}