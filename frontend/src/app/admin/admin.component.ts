import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ChatService } from '../shared/services/chat.service';
import { Role, Server } from '../shared/models/server';

enum Menu {
  ROLES = 1,
  MEMBERS = 2,
  GENERAL = 0
}

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.scss']
})

export class AdminComponent implements OnInit {

  Menu = Menu;
  @Input() server!: Server;
  @Output() serverDeleted = new EventEmitter<Server>();

  menuIndex: Menu = 0;
  selectedRole!: Role;
  constructor(private chatService:ChatService) { }

  ngOnInit(): void {
  }
  
  deleteServer() {
    this.chatService.deleteServer(this.server.id).subscribe();
    this.serverDeleted.emit(this.server);
  }

}
