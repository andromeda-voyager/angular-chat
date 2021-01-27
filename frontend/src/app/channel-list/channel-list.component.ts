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

  constructor() { }

  ngOnInit(): void {
    console.log(this.server.name)
  }




}
