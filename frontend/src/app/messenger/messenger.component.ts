import { Component, OnInit, Input } from '@angular/core';
import { Message } from '../shared/message';

@Component({
  selector: 'app-messenger',
  templateUrl: './messenger.component.html',
  styleUrls: ['./messenger.component.scss']
})
export class MessengerComponent implements OnInit {
  // Create WebSocket connection.
//  socket = new WebSocket('ws://localhost:8080/ws');
  messages: Message[] = [];
  @Input() inputText: string = "";
  constructor() {
  }


  ngOnInit(): void {
    // // Connection opened
    // this.socket.addEventListener('open', () => {
    // });

    // // Listen for messages
    // this.socket.addEventListener('message', (event) => {
    //   let message = JSON.parse(event.data)
    //   message.isIncomming = true;
    //   this.messages.push(message);
    // });

  }

  sendText() {
    console.log(this.inputText);
    let message = { text: this.inputText, isIncomming: false };
    this.messages.push(message)
    this.inputText = "";
  //  this.socket.send(JSON.stringify(message));
  }

}
