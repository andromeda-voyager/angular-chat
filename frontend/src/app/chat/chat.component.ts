import { Component, OnInit } from '@angular/core';
import { ChatService } from '../shared/services/chat.service';
import { Server } from '../shared/models/server';
import { User } from '../shared/models/user';
import { LoginService } from '../shared/services/login.service';
import { Router } from '@angular/router';
import { Channel } from '../shared/models/channel';

export enum Dialog {
  NONE = -1,
  ADMIN = 0,
  ADD_SERVER = 1,
  ADD_CHANNEL = 2
}

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.scss']
})

export class ChatComponent implements OnInit {
  Dialog = Dialog;
  user!: User;
  // servers = new Map<number, Server>()
  servers: Server[] = [];
  selectedServer!: Server;
  selectedChannel!: Channel;
  image: string = "";
  dialogOption: Dialog = Dialog.NONE;
  isMessagesActive = true;
  constructor(private chatService: ChatService, private loginService: LoginService, private router: Router) { }

  ngOnInit(): void {
    if(!this.loginService.isLoggedIn()) {
        this.router.navigate(['login']);
    }
    this.chatService.getServers().subscribe(servers => {
      this.servers = servers;
    });
    this.chatService.connect();
  }

  closeDialog() {
    this.dialogOption = Dialog.NONE;
  }

  selectServerOnClick(index: number) {
    console.log("index", index);
    this.chatService.connectToServer(this.servers[index].id).subscribe(server => {
      this.servers[index] = server; // more up to date version
      this.selectedServer = this.servers[index];
      console.log(server);
    })
  }

  onNewServer(server: Server) {
    this.servers.push(server);
    this.closeDialog();
  }

  onServerDeleted(server: Server) {
    this.servers.splice(this.servers.indexOf(server), 1);
    this.closeDialog();
    this.selectedServer = this.servers[0];
  }

  onNewChannel(channel: Channel) {
    this.selectedServer.channels.push(channel);
    this.closeDialog();
  }

  openDialog(option: Dialog) {
    this.dialogOption = option;
  }

}
