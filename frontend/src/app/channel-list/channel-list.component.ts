import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Channel } from '../shared/models/channel';
import { Server } from '../shared/models/server';

@Component({
  selector: 'app-channel-list',
  templateUrl: './channel-list.component.html',
  styleUrls: ['./channel-list.component.scss']
})
export class ChannelListComponent implements OnInit {
  @Input() server!: Server;
  showDialog = false;
  @Output() openChannelDialog = new EventEmitter();
  @Input() selectedChannel!: Channel;
  @Output() selectedChannelChange = new EventEmitter<Channel>();
  constructor() { }

  ngOnInit(): void {
    if (this.server.channels.length > 0) {
      this.selectedChannel = this.server.channels[0];
    }
    console.log(this.server.name)
  }




}
