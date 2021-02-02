import { Component, OnInit, Input } from '@angular/core';
import { Channel } from '../shared/models/channel';
import { Message } from '../shared/models/message';

@Component({
  selector: 'app-messages',
  templateUrl: './messages.component.html',
  styleUrls: ['./messages.component.scss']
})
export class PostsComponent implements OnInit {
  socket = new WebSocket('ws://localhost:8080/ws');
  @Input() channel!: Channel;

  constructor() {
  }


  ngOnInit(): void {
    // Connection opened
    this.socket.addEventListener('open', () => {
    });

    // Listen for messages
    this.socket.addEventListener('message', (event) => {
      let message = JSON.parse(event.data)
      console.log(message);
      if(!this.channel.messages) {
        this.channel.messages = []
      }
      // var m = {text:"hello there"}
      // this.socket.send(JSON.stringify(m));
      this.channel.messages.push(message);
    });


  }


}
