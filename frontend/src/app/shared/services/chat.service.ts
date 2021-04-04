import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import { environment } from 'src/environments/environment';
import { Server, NewServer, Invite } from '../models/server';
import { Update, UpdateEvent } from '../models/update';
import { Channel } from '../models/channel';
import { MessageService } from './message.service';

const formOptions = {
  headers: new HttpHeaders({
    'accept': 'application/json',
  }), credentials: 'same-origin',
  withCredentials: true
};

const jsonOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
    'accept': 'application/json',
  }), credentials: 'same-origin',
  withCredentials: true
};

const SERVER_URL = environment.BaseApiUrl + "/servers";
const JOIN_SERVER_URL = environment.BaseApiUrl + "/servers/join";
const CHANNEL_URL = environment.BaseApiUrl + "/channels";

@Injectable({
  providedIn: 'root'
})

export class ChatService {

  socket!: WebSocket;

  private updateSource = new Subject<Update>();
  update$ = this.updateSource.asObservable();

  constructor(private http: HttpClient, private messageService: MessageService) { }

  // return new Observable((observer: Observer<ConnectResponse>) => {

  connect() {
    this.socket = new WebSocket('ws://localhost:8080/ws');
    
    this.socket.addEventListener('message', (event) => {
      let update: Update = JSON.parse(event.data);
      switch (update.event) {
        case UpdateEvent.MESSAGE:
          this.messageService.newUpdate(update);
          break;
        case UpdateEvent.CHANNEL:
          break;
        case UpdateEvent.ROLE:
          break;
      }
    });

    this.socket.onclose = function (event) {
      console.log("socket closed")
    };

  }

  connectToServer(serverID: number): Observable<Server> {
    return this.http.get<Server>(SERVER_URL + "/" + serverID + "/connect", formOptions);
  }

  disconnect() {
    this.socket.close();
  }

  getServers(): Observable<Server[]> {
    return this.http.get<Server[]>(SERVER_URL, formOptions);
  }

  createServer(file: File, server: NewServer): Observable<Server> {
    const formData = new FormData();
    if (file) {
      formData.append("image", file, file.name);
    }
    formData.append("server", JSON.stringify(server));

    return this.http.post<Server>(SERVER_URL, formData, formOptions);
  }

  addChannel(channel: Channel): Observable<Channel> {
    return this.http.post<Channel>(CHANNEL_URL, channel, formOptions);
  }

  deleteServer(serverID: number): Observable<Server> {
    return this.http.delete<Server>(SERVER_URL + "/" + serverID, jsonOptions);
  }

  createInviteCode(serverID: number, invite: Invite): Observable<Invite> {
    console.log("creatin code");
    return this.http.post<Invite>(SERVER_URL + "/" + serverID + "/invite", invite, jsonOptions);
  }

  joinServer(invite: Invite): Observable<Server> {
    return this.http.post<Server>(JOIN_SERVER_URL, invite, jsonOptions);
  }

}
