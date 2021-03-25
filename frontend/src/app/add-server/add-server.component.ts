import { Component, EventEmitter, OnInit, Input, Output } from '@angular/core';
import { ChatService } from '../shared/services/chat.service';
import { Server, NewServer, Invite } from '../shared/models/server';

@Component({
  selector: 'app-add-server',
  templateUrl: './add-server.component.html',
  styleUrls: ['./add-server.component.scss']
})
export class AddServerComponent implements OnInit {
  @Input() server: NewServer = {name: "", description: ""}
  @Input() invite: Invite = { code: "" }
  @Output() newServer = new EventEmitter<Server>();
  file: File = null!;
  serverImageUrl: string = "http://localhost:8080/static/images/default-avatar.jpg";
  showRequired = false;
  showCreateServer = false;
  showAddServer = false;

  constructor(private chatService: ChatService) { }

  ngOnInit(): void {
  }

  joinServer() {
    if (this.invite.code.length > 7) {
      this.chatService.joinServer(this.invite).subscribe(connection => {
      //  this.newConnection.emit(connection);
      })
    }
    else {
      this.showRequired = true;
    }

  }

  onChange(event: any) {
    this.file = event.target.files[0];
    var reader = new FileReader();
    reader.onload = (event: any) => {
      this.serverImageUrl = event.target.result;
    }
    reader.readAsDataURL(event.target.files[0]);
    //  this.chatService.uploadImage(this.file);
  }

  createServer() {
    if (this.server.name.length != 0) {
      console.log(this.server.name);
      this.chatService.createServer(this.file, this.server).subscribe(server => {
        this.newServer.emit(server);
      })
    }
    else {
      this.showRequired = true;
    }
  }

}
