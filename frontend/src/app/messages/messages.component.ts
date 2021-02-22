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
  selectedMessage!: Message;
  top = 0;
  left = 0;
  showMessageMenu = false;
  constructor(private messageService: MessageService) {
  }

  ngOnInit(): void {
    this.messageService.newMessage$.subscribe(message => {
      console.log("new message");
      this.messages.push(message);
    });

    this.messageService.modifyMessage$.subscribe(message => {
      this.modifyMessage(message);
    });

    this.messageService.deleteMessage$.subscribe(message => {
      this.deleteMessage(message);
    });
  }

  ngOnChanges(changes: SimpleChanges) {
    this.messageService.connectToChannel(this.channel.id).subscribe(messages => {
     if(messages.length > 0) console.log(messages[0].timePosted);
      this.messages = messages;
      if(messages.length > 0) {
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

  getTime(message:Message) {
    return new Date(message.timePosted).toLocaleTimeString()
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

  deleteMessageOnClick() {
    this.messageService.deleteMessage(this.selectedMessage);
    this.showMessageMenu=false;
  }

  openMessageMenuOnClick(event: MouseEvent, message: Message) {
    // this.top = event.pageY -5;
    // this.left = event.pageX -5;
    this.showMessageMenu = true;
    this.selectedMessage = message;
  }


}
