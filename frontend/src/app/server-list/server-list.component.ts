import { Component, OnInit,Input, Output, EventEmitter } from '@angular/core';
import { Server } from '../shared/models/server';
import { Dialog } from '../chat/chat.component'

@Component({
  selector: 'app-server-list',
  templateUrl: './server-list.component.html',
  styleUrls: ['./server-list.component.scss']
})
export class ServerListComponent implements OnInit {

  constructor() { }
  @Input() servers: Server[] = [];
  Dialog = Dialog;
  @Output() openChannelDialog = new EventEmitter<Dialog>();
  @Input() selectedServer!: Server;
  @Output() selectedServerChange = new EventEmitter<Server>();
  ngOnInit(): void {
  }


  selectServerOnClick(server: Server) {
    this.selectedServerChange.emit(server);
  }

}
