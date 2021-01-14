import { Component, Input, OnInit } from '@angular/core';
import { ChatService } from '../shared/chat.service';
import { Connection } from '../shared/server';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.scss']
})
export class AdminComponent implements OnInit {

  @Input() connection!: Connection;
  constructor(private chatService:ChatService) { }

  ngOnInit(): void {
  }
  
  deleteServer() {
    this.chatService.deleteServer(this.connection.server.id).subscribe();
  }

}
