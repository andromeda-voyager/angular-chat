import { Component, OnChanges, OnInit, SimpleChanges, Input } from '@angular/core';
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
  showTime = false;
  selectedMessage: Message | undefined
  constructor(private messageService: MessageService) {
  }

  ngOnInit(): void {
    this.messageService.newMessage$.subscribe(message => {
      this.messages.push(message);
    });

    this.messageService.modifyMessage$.subscribe(message => {
      this.replaceMessage(message);
    });

    this.messageService.deleteMessage$.subscribe(message => {
      this.removeMessage(message);
    });
  }

  ngOnChanges(changes: SimpleChanges) {
    this.messageService.connectToChannel(this.channel.id).subscribe(messages => {
      if (messages.length > 0) console.log(messages[0].timePosted);
      this.messages = messages;
      if (messages.length > 0) {
      }
    });
  }

  sendMessage() {
    this.messageService.postMessage(
      { channelID: this.channel.id, text: this.messageText, media: "" }
    )
    this.messageText = "";
  }

  getDate(message: Message) {
    return new Date(message.timePosted).toDateString()
  }

  getTime(message: Message) {
    return new Date(message.timePosted).toLocaleTimeString()
  }

  replaceMessage(message: Message) {
    let index = this.messages.findIndex(function (m) {
      return m.id == message.id
    });

    if (index >= 0) {
      this.messages[index] = message;
    }
  }

  removeMessage(message: Message) {
    let index = this.messages.findIndex(function (m) {
      return m.id == message.id
    });

    if (index >= 0) {
      this.messages.splice(index, 1);
    }
  }

  deleteMessageOnClick(message: Message) {
    this.messageService.deleteMessage(message);
    this.selectedMessage = undefined;
  }

  editMessageOnClick(message: Message) {
    message.isEditable = true;
    this.selectedMessage = undefined;
  }

  modifyMessage(event: any, message: Message) {
    message.text = event.target.textContent;
    this.messageService.putMessage(message);
    message.isEditable = false;
  }

}
