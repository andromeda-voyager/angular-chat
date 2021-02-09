import { Component, OnChanges, SimpleChanges, Input } from '@angular/core';
import { Channel } from '../shared/models/channel';
import { Message } from '../shared/models/message';
import { MessageService } from '../shared/services/message.service';

@Component({
  selector: 'app-messages',
  templateUrl: './messages.component.html',
  styleUrls: ['./messages.component.scss']
})
export class PostsComponent implements OnChanges {
  @Input() channel!: Channel;
  @Input() messages: Message[] = []
  @Input() messageText: string = "";

  constructor(private messageService: MessageService) {
    // let m: NewMessage = { channelID: 0, text: "hello", media: "none", id: 0, timePosted: new Date() }
    // this.messages.push(m);
  }

  ngOnChanges(changes: SimpleChanges) {
    console.log("something changed");
    this.messageService.getMessages(this.channel.id).subscribe(messages => {
      this.messages = messages;
    });
  }

  sendMessage() {
    this.messageService.postMessage(
      { channelID: this.channel.id, text: "hello", media: "none" }
    )
  }

  modifyMessage(message: Message) {
    let index = this.messages.findIndex(function (m) {
      return m.id == message.id
    });

    if (index >= 0) {
      this.messages[index] = message;
    }
  }

  deleteMessage(message: Message) {
    let index = this.messages.findIndex(function (m) {
      return m.id == message.id
    });

    if (index >= 0) {
      this.messages.splice(index, 1);
    }
  }

  ngOnInit(): void {
    this.messageService.newMessage$.subscribe(message => {
      this.messages.push(message);
    });

    this.messageService.modifyMessage$.subscribe(message => {
      this.modifyMessage(message);
    });

    this.messageService.deleteMessage$.subscribe(message => {
      this.deleteMessage(message);
    });
  }


}
