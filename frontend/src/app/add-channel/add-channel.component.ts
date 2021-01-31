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

  addChannel() {
    console.log(this.name);
    this.chatService.addChannel({ serverID: this.server.id, Channel:{name: this.name, id:0, posts:[] }, Roles:this.server.roles}).subscribe(channel => {
      this.newChannel.emit(channel);
    })
  }

  checkValue(event: any) {
    console.log(event)
  }

}
