import { Component, OnInit, Input } from '@angular/core';
import { Message } from '../shared/models/message';
import { Post } from '../shared/models/post';

@Component({
  selector: 'app-posts',
  templateUrl: './posts.component.html',
  styleUrls: ['./posts.component.scss']
})
export class PostsComponent implements OnInit {
  // Create WebSocket connection.
//  socket = new WebSocket('ws://localhost:8080/ws');
  @Input() posts: Post[] = [];

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


}
