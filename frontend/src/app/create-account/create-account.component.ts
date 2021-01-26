import { Component, OnInit, Input } from '@angular/core';
import { Router } from '@angular/router';
import { Account } from '../shared/models/user';
import { AccountService } from '../shared/services/account.service';

const emailRegex = /.+@.+\..+/; // basic email syntax (not a valid check)

@Component({
  selector: 'app-create-account',
  templateUrl: './create-account.component.html',
  styleUrls: ['./create-account.component.scss']
})
export class CreateAccountComponent implements OnInit {

  @Input() account: Account = { password: "", email: "", username: "", code:""}
  hide = false;
  file: File = null!;
  avatarURL: string = "assets/default-avatar.jpg"
  showRequired = false;
  showFirstCard = true;
  constructor(private accountService: AccountService, private router: Router) { }

  ngOnInit(): void {
  }

  createAccount() {
      this.accountService.createAccount(this.file, this.account).subscribe(user => {
        console.log(user);
        this.router.navigate(['chat']);
      })
  }

  next() {
    if (this.userFieldsValid()) {
      this.accountService.sendVerificationCode(this.account.email);
      this.showFirstCard = false;
    } else this.showRequired = true;
  }

  onUpload() {

  }

  toggleHide() {
  }

  onChange(event: any) {
    this.file = event.target.files[0];
    this.isValidImage()
    var reader = new FileReader();
    reader.onload = (event: any) => {
      this.avatarURL = event.target.result;
    }
  
    reader.readAsDataURL(event.target.files[0]);
  }

  userFieldsValid() {
    return (this.isValidEmail() &&
      this.account.username.length > 0 &&
      this.account.password.length > 7 &&
      this.isValidImage())

  }

  // does a simple check for email syntax (not complete)
  isValidEmail() {
    return emailRegex.test(this.account.email)
  }

  isValidImage() {
    if (this.file) {
      console.log(this.file.size/1000000)
      return this.file.size/1000000 <= 1; 
    }
    return true; // null is valid since default image will be used on the server
  }

  isValidCodeLength() {
    return this.account.code.length > 4;
  }


}
