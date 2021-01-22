import { Component, OnInit, Input, ViewChild, ElementRef } from '@angular/core';
import { ChatService } from '../shared/chat.service';
import { Server, Post } from '../shared/server';
import { User } from '../shared/user';
import { LoginService } from '../shared/login.service';

declare var MediaRecorder: any;
@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.scss']
})
export class ChatComponent implements OnInit {

  user!: User;
  servers: Server[] = []
  dialogOpen: boolean = false;
  currentServer!: Server;
  currentChannel!: Server;
  image: string = "";
  showAdminPanel = false;
  @Input() postText: string = "";

  constructor(private chatService: ChatService, private loginService: LoginService) { }

  ngOnInit(): void {
    let loginData = this.loginService.getLoginResponse();
    console.log(loginData);
    this.user = loginData.user;
    if (loginData.servers) {
      this.servers = loginData.servers;
      this.currentServer = this.servers[0];
    }
  }

  closeDialog() {
    this.dialogOpen = false;

    console.log("closing dialogs")
  }

  changeServer(index: number) {
    this.currentServer = this.servers[index];
  }

  onNewServer(server: Server) {
    // if (this.servers == null) {
    //   this.account.connections = [];
    // }
    this.servers.push(server);
  }

  addServer() {
    console.log("opening dialogs")
    this.dialogOpen = true;
  }

  post() {
    // this.chatService.post({ serverID: this.currentConnection.serverID, text: this.postText, media: "" });
    this.postText = "";
    //  this.socket.send(JSON.stringify(message));
  }

  openAdminPanel() {
    this.showAdminPanel = true;
  }

  closeAdminPanel() {
    this.showAdminPanel = false;
  }
}
