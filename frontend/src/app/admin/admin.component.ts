import { Component, Input, OnInit } from '@angular/core';
import { ChatService } from '../shared/services/chat.service';
import { Server } from '../shared/models/server';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.scss']
})
export class AdminComponent implements OnInit {

  @Input() server!: Server;
  constructor(private chatService:ChatService) { }

  ngOnInit(): void {
  }
  
  deleteServer() {
    this.chatService.deleteServer(this.server).subscribe();
  }

}
