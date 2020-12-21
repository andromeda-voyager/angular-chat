import { Component, OnInit, Input } from '@angular/core';
import { ChatService } from '../shared/chat.service';

@Component({
  selector: 'app-create-account',
  templateUrl: './create-account.component.html',
  styleUrls: ['./create-account.component.scss']
})
export class CreateAccountComponent implements OnInit {

  @Input() password: string = "";
  @Input() email: string = "";
  @Input() name: string = "";
  @Input() username: string = "";
  hide = false;
  file!: File;
  avatarURL: string = ""
  constructor(private chatService: ChatService) { }

  ngOnInit(): void {
  }

  createAccount() {

  }

  onUpload() {

  }

  toggleHide() {
  }

  onChange(event: any) {
    this.file = event.target.files[0];
    var reader = new FileReader();
    reader.onload = (event: any) => {
      this.avatarURL = event.target.result;
    }
    reader.readAsDataURL(event.target.files[0]);
    this.chatService.upload(this.file);
  }


}
