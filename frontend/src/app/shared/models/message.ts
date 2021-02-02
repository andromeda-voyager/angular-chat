export interface Message {
     accountID: number
     text: string
     media: string
     timePosted: Date
 }
 
 export interface NewMessage {
     channelID: number
     text: string
     media?: string
 }