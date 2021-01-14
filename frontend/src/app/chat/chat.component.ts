import { Component, OnInit, Input, ViewChild, ElementRef } from '@angular/core';
import { ChatService } from '../shared/chat.service';
import { Connection, Server, Post } from '../shared/server';
import { Account } from '../shared/account';
import { LoginService } from '../shared/login.service';

declare var MediaRecorder: any;
@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.scss']
})
export class ChatComponent implements OnInit {

  account!: Account;
  dialogOpen: boolean = false;
  currentConnection!: Connection;
  image: string = "";
  showAdminPanel = false;
  @Input() postText: string = "";

  constructor(private chatService: ChatService, private loginService: LoginService) { }

  ngOnInit(): void {
    this.account = this.loginService.getUserData();
    if (this.account) {
      if (this.account.connections && this.account.connections.length > 0) {
        this.currentConnection = this.account.connections[0];
      } else this.account.connections = [];
    }

  }

  closeDialog() {
    this.dialogOpen = false;

    console.log("closing dialogs")
  }

  changeConnection(index: number) {
    this.currentConnection = this.account.connections[index];
  }

  onNewConnection(connection: Connection) {
    if (this.account.connections == null) {
      this.account.connections = [];
    }
    this.account.connections.push(connection);
  }

  addServer() {
    console.log("opening dialogs")
    this.dialogOpen = true;
  }

  post() {
    this.chatService.post({ serverID: this.currentConnection.server.id, text: this.postText, mediaURL: "" });
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
