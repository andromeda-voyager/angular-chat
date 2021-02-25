import { Member } from "./server";

export interface Message {
    channelID: number
    id: number
    member: Member
    text: string
    media: string
    timePosted: Date
    isEdited: boolean
    parentID: number
    isEditable: boolean
}
 
export interface NewMessage {
    channelID: number
    text: string
    media: string
}