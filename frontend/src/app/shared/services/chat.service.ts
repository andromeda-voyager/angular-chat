import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { Server, Invite, ServerRequest } from '../models/server';
import { NewChannel, Channel } from '../models/channel';
import { NewPost } from '../models/post';

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

const createServerUrl = environment.BaseApiUrl + "/create-server";
const joinServerUrl = environment.BaseApiUrl + "/join-server";
const getPostsURL = environment.BaseApiUrl + "/posts";
const postURL = environment.BaseApiUrl + "/post";
const deleteServerURL = environment.BaseApiUrl + "/delete-server";
const createChannelURL = environment.BaseApiUrl + "/create-channel";

@Injectable({
  providedIn: 'root'
})

export class ChatService {
  constructor(private http: HttpClient) { }

  createServer(file: File, server: Server): Observable<Server> {
    const formData = new FormData();
    if (file) {
      formData.append("image", file, file.name);
    }
    formData.append("server", JSON.stringify(server));

    return this.http.post<Server>(createServerUrl, formData, formOptions);
  }

  createChannel(serverRequest: ServerRequest): Observable<Channel> {
    return this.http.post<Channel>(createChannelURL, serverRequest, formOptions);
  }

  deleteServer(server: Server): Observable<Server> {
    return this.http.post<Server>(deleteServerURL, server, jsonOptions);
  }

  joinServer(invite: Invite): Observable<Server> {
    return this.http.post<Server>(joinServerUrl, invite);
  }

  getPosts(serverID: number): Observable<Server> {
    return this.http.get<Server>(getPostsURL + "?serverID=" + serverID, jsonOptions);
  }

  post(post: NewPost) {
    this.http.post(postURL, post, jsonOptions).subscribe();
  }
}
