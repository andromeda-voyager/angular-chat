import { Injectable } from '@angular/core';
import { Observable, Observer, Subject } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Message, NewMessage } from '../models/message';

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


  getMessages(channelID: number): Observable<Message[]> {
    return this.http.get<Message[]>(CHANNEL_URL + "/" + channelID + "/messages", credentialsOption);
  }

  connectToChannel(channelID: number): Observable<Message[]> {
    return this.http.get<Message[]>(CHANNEL_URL + "/" + channelID + "/connect", credentialsOption);
  }

  postMessage(message: NewMessage) {
    this.http.post<Message>(CHANNEL_URL + "/" + message.channelID + "/messages", message, credentialsOption).subscribe();
  }

}
