import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ChatService } from '../shared/services/chat.service';
import { Channel, Override } from '../shared/models/channel';
import { Server } from '../shared/models/server';

@Component({
  selector: 'app-add-channel',
  templateUrl: './add-channel.component.html',
  styleUrls: ['./add-channel.component.scss']
})

export class AddChannelComponent implements OnInit {
  @Input() server!: Server;
  name = "";
  @Output() newChannel = new EventEmitter<Channel>();
  showRequired = false;
  @Input() isSecret = false;
  showHint = false;
  @Input() overrides: Override[] = [];
  constructor(private chatService: ChatService) { }

  ngOnInit(): void {
    for(let role of this.server.roles) {
      console.log(role.name);
      this.overrides.push({roleID: role.id, permissions:0})
    }
  }

  addChannel() {
    console.log(this.name);
    this.chatService.addChannel(
      { name: this.name, id: 0, serverID: this.server.id, overrides:this.overrides }
    ).subscribe(channel => {
      this.newChannel.emit(channel);
    })
  }

  checkValue(event: any) {
    console.log(event)
  }

}
