import { Component, OnInit, Input } from '@angular/core';
import { Router } from '@angular/router';
import { ChatService } from '../shared/chat.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  @Input() password: string = "";
  @Input() email: string = "";
  isPasswordCorrect = true;
  hide = true;
  loginFailed = false;
  constructor(private chatService: ChatService, private router: Router) { }

  ngOnInit(): void {
  }

  login() {
    this.chatService.login({ email: this.email, password: this.password }).subscribe(user => {
      console.log(user)
      console.log(":")
      console.log("password correct");
      this.chatService.addUserData(user)
      this.router.navigate(['chat']);
    }, error => this.handleError(error));
  }

  handleError(error: any) {
    this.loginFailed = true;
    console.log(error);
  }

}
