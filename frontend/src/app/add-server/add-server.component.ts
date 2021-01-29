import { Component, EventEmitter, OnInit, Input, Output } from '@angular/core';
import { ChatService } from '../shared/services/chat.service';
import { Server, Invite } from '../shared/models/server';

@Component({
  selector: 'app-add-server',
  templateUrl: './add-server.component.html',
  styleUrls: ['./add-server.component.scss']
})
export class AddServerComponent implements OnInit {
  @Input() server: Server = { id: 0, name: "", description: "", image: "", roles:[], channels:[]}
  @Input() invite: Invite = { code: "" }
  @Output() newServer = new EventEmitter<Server>();
  file: File = null!;
  serverImageUrl: string = "assets/default-avatar.jpg"
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
    if (this.server.name.length > 3) {
      console.log(this.server.name);
      this.chatService.createServer(this.file, this.server).subscribe(connection => {
        this.newServer.emit(connection);
      })
    }
    else {
      this.showRequired = true;
    }
  }

}
