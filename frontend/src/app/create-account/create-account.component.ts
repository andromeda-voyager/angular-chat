import { Component, OnInit, Input } from '@angular/core';
import { ChatService } from '../shared/chat.service';
import { User } from '../shared/user';

const emailRegex = /.+@.+\..+/;

@Component({
  selector: 'app-create-account',
  templateUrl: './create-account.component.html',
  styleUrls: ['./create-account.component.scss']
})
export class CreateAccountComponent implements OnInit {

  @Input() user: User = { password: "", email: "", name: "", username: "", code: "" }
  hide = false;
  file: File = null!;
  avatarURL: string = "assets/default-avatar.jpg"
  showRequired = false;
  showFirstCard = true;
  constructor(private chatService: ChatService) { }

  ngOnInit(): void {
  }

  createAccount() {
      this.chatService.createAccount(this.file, this.user);
  }

  next() {
    if (this.userFieldsValid()) {
      this.chatService.sendVerificationCode(this.user.email);
      this.showFirstCard = false;
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
    return (this.isValidEmail() &&
      this.user.name.length > 0 &&
      this.user.username.length > 0 &&
      this.user.password.length > 7)
  }

  // does a simple check for email syntax (not complete)
  isValidEmail() {
    return emailRegex.test(this.user.email)
  }

  isValidCodeLength() {
    return this.user.code.length > 4;
  }


}
