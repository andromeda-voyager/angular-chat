import { Component, OnInit, Input } from '@angular/core';
import { Message } from '../shared/models/message';
import { MessageService } from '../shared/services/message.service';

@Component({
  selector: 'app-messages',
  templateUrl: './messages.component.html',
  styleUrls: ['./messages.component.scss']
})
export class PostsComponent implements OnInit {
  @Input() channelID: number = 0;
  @Input() messages: Message[] = []

  constructor(private messageService: MessageService) {
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
