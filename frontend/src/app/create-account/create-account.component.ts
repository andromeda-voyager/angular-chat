import { Component, OnInit, Input } from '@angular/core';
import { Router } from '@angular/router';
import { Account } from '../shared/models/user';
import { AccountService } from '../shared/services/account.service';
import { LoginService } from '../shared/services/login.service';

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
  constructor(private accountService: AccountService, private loginService: LoginService, private router: Router) { }

  ngOnInit(): void {
  }

  createAccountOnClick() {
    if(this.isValidCodeLength()) {
      console.log(this.account)
      this.accountService.createAccount(this.file, this.account).subscribe(user => {
        console.log(user);
        this.loginService.login(user);
        this.router.navigate(['chat']);
      })
    } else this.showRequired = true;
  }

  nextOnClick() {
    if (this.isAccountValid()) {
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

  isAccountValid() {
    return (this.isEmailValid() &&
      this.account.username.length > 0 &&
      this.account.password.length > 7 &&
      this.isValidImage())
  }

  // does a simple check for email syntax (not complete)
  isEmailValid() {
    return emailRegex.test(this.account.email)
  }

  isValidImage() {
    if (this.file) {
      return this.file.size/1000000 <= 1; 
    }
    return true; // null is valid since default image will be used on the server
  }

  isValidCodeLength() {
    return this.account.code.length > 4;
  }


}
