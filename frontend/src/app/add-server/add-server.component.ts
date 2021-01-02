import { Component, OnInit, Input } from '@angular/core';
import { ChatService } from '../shared/chat.service';
import { Server } from '../shared/server';

@Component({
  selector: 'app-add-server',
  templateUrl: './add-server.component.html',
  styleUrls: ['./add-server.component.scss']
})
export class AddServerComponent implements OnInit {
  @Input() server: Server = { name: "", description: "", serverImageUrl: "" }
  file: File = null!;
  serverImageUrl: string = "assets/default-avatar.jpg"
  showRequired = false;
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
    if (this.server.name.length > 3) {
      this.chatService.createServer(this.file, this.server);
    }
    else {
      this.showRequired = true;
    }
  }

}
