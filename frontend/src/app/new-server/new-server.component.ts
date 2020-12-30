import { Component, OnInit, Input } from '@angular/core';
import { ChatService } from '../shared/chat.service';
import { Server } from '../shared/server';

@Component({
  selector: 'app-new-server',
  templateUrl: './new-server.component.html',
  styleUrls: ['./new-server.component.scss']
})
export class NewServerComponent implements OnInit {
  @Input() server: Server = { name: "", description: "", serverImageUrl: "" }
  file: File = null!;
  serverImageUrl: string = "assets/default-server.jpg"
  constructor(private chatService: ChatService) { }

  ngOnInit(): void {
  }


  onChange(event: any) {
    this.file = event.target.files[0];
    var reader = new FileReader();
    reader.onload = (event: any) => {
      this.serverImageUrl = event.target.result;
    }
    reader.readAsDataURL(event.target.files[0]);
    //  this.chatService.uploadImage(this.file);
  }

  createServer() {
    this.chatService.createServer(this.file, this.server);

  }

}
