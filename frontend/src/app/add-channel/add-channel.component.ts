import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ChatService } from '../shared/services/chat.service';
import { Channel, ChannelPermissions } from '../shared/models/channel';
import { Server } from '../shared/models/server';

@Component({
  selector: 'app-add-channel',
  templateUrl: './add-channel.component.html',
  styleUrls: ['./add-channel.component.scss']
})

export class AddChannelComponent implements OnInit {
  @Input() server!: Server;
  name = "";
  @Input() channelPermissions: ChannelPermissions[] = [];
  @Output() newChannel = new EventEmitter<Channel>();
  showRequired = false;
  @Input() isSecret = false;
  showHint = false;
  constructor(private chatService: ChatService) { }

  ngOnInit(): void {
  }

  createChannel() {
    console.log(this.name);
    this.chatService.createChannel({ serverID: this.server.id, name: this.name, channelPermissions: this.channelPermissions }).subscribe(channel => {
      this.newChannel.emit(channel);
    })
  }

  checkValue(event: any) {
    console.log(event)
  }

}
