import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Channel } from '../shared/models/channel';
import { Server } from '../shared/models/server';
import { Dialog } from '../chat/chat.component'

@Component({
  selector: 'app-server-panel',
  templateUrl: './server-panel.component.html',
  styleUrls: ['./server-panel.component.scss']
})
export class ServerPanelComponent implements OnInit {
  @Input() server!: Server;
  showDialog = false;
  Dialog = Dialog;

  @Output() openChannelDialog = new EventEmitter<Dialog>();
  @Input() selectedChannel!: Channel;
  @Output() selectedChannelChange = new EventEmitter<Channel>();
  constructor() { }

  ngOnInit(): void {
    if (this.server.channels.length > 0) {
      this.selectedChannel = this.server.channels[0];
    }
    console.log(this.server.name)
  }

  selectChannelOnClick(channel: Channel) {
    this.selectedChannelChange.emit(channel);
  }




}
