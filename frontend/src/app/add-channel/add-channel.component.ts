import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ChatService } from '../shared/services/chat.service';
import { Channel, ChannelPermissions } from '../shared/models/channel';

@Component({
  selector: 'app-add-channel',
  templateUrl: './add-channel.component.html',
  styleUrls: ['./add-channel.component.scss']
})

export class AddChannelComponent implements OnInit {
  @Input() serverID!: number;
  @Input() name = "";
  @Input() channelPermissions: ChannelPermissions[] = [];
  @Output() newChannel = new EventEmitter<Channel>();
  showRequired = false;
  @Input() isSecret = false;
  showHint = false;
  constructor(private chatService: ChatService) { }

  ngOnInit(): void {
  }

  createChannel() {
    this.chatService.createChannel({ serverID: this.serverID, name: this.name, channelPermissions: this.channelPermissions })
  }

  checkValue(event: any) {
    console.log(event)
  }

}
