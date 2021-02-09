import { Member } from "./server";

export interface Message {
    channelID: number
    id: number
    member?: Member
    text: string
    media: string
    timePosted: Date
}
 
export interface NewMessage {
    channelID: number
    text: string
    media: string
}