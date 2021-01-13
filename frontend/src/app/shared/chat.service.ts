import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { NewServer, Server, Invite, Connection } from '../shared/server';

const formOptions = {
  headers: new HttpHeaders({
    'accept': 'application/json',
  }), credentials: 'same-origin',
  withCredentials: true
};

const createServerUrl = environment.BaseApiUrl + "/create-server";
const joinServerUrl = environment.BaseApiUrl + "/join-server";
const getPostsURL = environment.BaseApiUrl + "/posts";

@Injectable({
  providedIn: 'root'
})

export class ChatService {
  constructor(private http: HttpClient) { }

  createServer(file: File, server: NewServer): Observable<Connection> {
    const formData = new FormData();
    if (file) {
      formData.append("image", file, file.name);
    }
    formData.append("server", JSON.stringify(server));

    return this.http.post<Connection>(createServerUrl, formData, formOptions);
  }

  joinServer(invite: Invite): Observable<Server> {
    return this.http.post<Server>(joinServerUrl, invite);
  }

  getPosts(serverID: number): Observable<Server> {
    return this.http.get<Server>(getPostsURL + "?serverID=" + serverID, formOptions);
  }
}
