import { Injectable } from '@angular/core';
import { Observable, Observer, Subject } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Message, NewMessage } from '../models/message';
import { Update, UpdateType } from '../models/update';

const CHANNEL_URL = environment.BaseApiUrl + "/channels";

const credentialsOption = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
    'accept': 'application/json',
  }), credentials: 'same-origin',
  withCredentials: true
};

@Injectable({
  providedIn: 'root'
})

export class MessageService {

  private newMessageSource = new Subject<Message>();
  newMessage$ = this.newMessageSource.asObservable();

  private modifyMessageSource = new Subject<Message>();
  modifyMessage$ = this.modifyMessageSource.asObservable();

  private deleteMessageSource = new Subject<Message>();
  deleteMessage$ = this.deleteMessageSource.asObservable();

  constructor(private http: HttpClient) { }


  newUpdate(update: Update) {
    switch (update.type) {
      case UpdateType.NEW:
        this.newMessageSource.next(update.message);
        break;
      case UpdateType.DELETE:
        this.deleteMessageSource.next(update.message);
        break;
      case UpdateType.MODIFY:
        this.modifyMessageSource.next(update.message);
        break;
    }

  }
  getMessages(channelID: number): Observable<Message[]> {
    return this.http.get<Message[]>(CHANNEL_URL + "/" + channelID + "/messages", credentialsOption);
  }

  connectToChannel(channelID: number): Observable<Message[]> {
    return this.http.get<Message[]>(CHANNEL_URL + "/" + channelID + "/connect", credentialsOption);
  }

  postMessage(message: NewMessage) {
    this.http.post<Message>(CHANNEL_URL + "/" + message.channelID + "/messages", message, credentialsOption).subscribe();
  }

  putMessage(message: Message) {
    this.http.put<Message>(CHANNEL_URL + "/" + message.channelID + "/messages/" + message.id, message, credentialsOption).subscribe();
  }

  deleteMessage(message: Message) {
    this.http.delete<Message>(CHANNEL_URL + "/" + message.channelID + "/messages/" + message.id, credentialsOption).subscribe();
  }

}
