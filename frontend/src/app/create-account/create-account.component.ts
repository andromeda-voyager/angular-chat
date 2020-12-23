import { Component, OnInit, Input } from '@angular/core';
import { ChatService } from '../shared/chat.service';
import { User } from '../shared/user';

@Component({
  selector: 'app-create-account',
  templateUrl: './create-account.component.html',
  styleUrls: ['./create-account.component.scss']
})
export class CreateAccountComponent implements OnInit {

  @Input() user: User = { password: "", email: "", name: "", username: "" }
  hide = false;
  file!: File;
  avatarURL: string = "assets/default-avatar.jpg"
  showRequired = false;
  constructor(private chatService: ChatService) { }

  ngOnInit(): void {
  }

  createAccount() {
    if (this.userFieldsValid()) {
      this.chatService.createAccount(this.file, this.user);
    } else this.showRequired = true;
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

  userFieldsValid() {
    return (this.user.email.length > 4 &&
      this.user.name.length > 0 &&
      this.user.username.length > 0 &&
      this.user.password.length > 7)
  }


}
