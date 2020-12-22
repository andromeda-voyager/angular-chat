import { Component, OnInit, Input } from '@angular/core';
import { ChatService } from '../shared/chat.service';
import { User } from '../shared/user';

@Component({
  selector: 'app-create-account',
  templateUrl: './create-account.component.html',
  styleUrls: ['./create-account.component.scss']
})
export class CreateAccountComponent implements OnInit {

  @Input() user:User = {password: "", email: "" , name:"", username:""}
  hide = false;
  file!: File;
  avatarURL: string = ""
  constructor(private chatService: ChatService) { }

  ngOnInit(): void {
  }

  createAccount() {
    this.chatService.createAccount(this.file, this.user);
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
  //  this.chatService.uploadImage(this.file);
  }


}
