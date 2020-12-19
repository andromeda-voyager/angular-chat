import { Component, OnInit, Input } from '@angular/core';
import { Message } from '../shared/message';

@Component({
  selector: 'app-messenger',
  templateUrl: './messenger.component.html',
  styleUrls: ['./messenger.component.scss']
})
export class MessengerComponent implements OnInit {
// Create WebSocket connection.
socket = new WebSocket('ws://localhost:8080');
messages: Message[] = [];
@Input() inputText: string = "";
constructor() {
}


ngOnInit(): void {
  // Connection opened
  this.socket.addEventListener('open', () => {
  });

  // Listen for messages
  this.socket.addEventListener('message', (event) => {
    this.messages.push({ text: event.data, isIncomming: true });
  });

}

sendText() {
  console.log(this.inputText);
  let m = { text: this.inputText, isIncomming: false };
  this.messages.push(m)
  this.inputText = "";
  this.socket.send(m.text);
}

}
