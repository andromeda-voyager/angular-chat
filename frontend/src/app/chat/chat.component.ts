import { Component, OnInit, Input } from '@angular/core';
import { ChatService } from '../shared/services/chat.service';
import { Server } from '../shared/models/server';
import { LoginResponse, User } from '../shared/models/user';
import { LoginService } from '../shared/services/login.service';
import { Router } from '@angular/router';
import { Channel } from '../shared/models/channel';

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.scss']
})
export class ChatComponent implements OnInit {

  user!: User;
  servers: Server[] = []
  selectedServer!: Server;
  selectedChannel!: Channel;
  image: string = "";
  showDialog = false;
  dialogOption: number =0;
  @Input() postText: string = "";

  constructor(private chatService: ChatService, private loginService: LoginService, private router: Router) { }

  ngOnInit(): void {
    let data = sessionStorage.getItem("loginData");
    if (data != null) {
      let loginData: LoginResponse = JSON.parse(data);
      this.user = loginData.user;
      this.servers = loginData.servers;
      if (this.servers) {
        this.selectedServer = this.servers[0];
        if(this.selectedServer.channels) {
          this.selectedChannel = this.selectedServer.channels[0];
        }
      } 
    } else {
      this.router.navigate(['login']);
    }

  }

  closeDialog() {
    this.showDialog = false;

    console.log("closing dialogs")
  }

  changeServer(index: number) {
    this.selectedServer = this.servers[index];
  }

  onNewServer(server: Server) {
    // if (this.servers == null) {
    //   this.account.connections = [];
    // }
    this.servers.push(server);
    this.showDialog = false;
  }

  onNewChannel(channel: Channel) {
    this.selectedServer.channels.push(channel);
  }

 

  post() {
    // this.chatService.post({ serverID: this.currentConnection.serverID, text: this.postText, media: "" });
    this.postText = "";
    //  this.socket.send(JSON.stringify(message));
  }

  openDialog(option: number) {
    this.dialogOption = option;
    this.showDialog = true;  
  }

}
