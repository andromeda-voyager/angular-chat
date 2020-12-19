import { Component, OnInit, Input } from '@angular/core';
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
  constructor(private chatService: ChatService) { }

  ngOnInit(): void {
  }

  login() {
    this.chatService.login({ email: this.email, password: this.password }).subscribe(user => {
      console.log(user.username)
      console.log("password correct");
    }, error => this.handleError(error));
  }

  handleError(error: any) {
    console.log(error);
  }

}
