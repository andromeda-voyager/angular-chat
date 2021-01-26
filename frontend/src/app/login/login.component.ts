import { Component, OnInit, Input } from '@angular/core';
import { Router } from '@angular/router';
import { LoginService } from '../shared/services/login.service';

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
  constructor(private loginService: LoginService, private router: Router) {
  }

  ngOnInit(): void {
    this.loginService.loginWithCookie().subscribe(user => {

      this.router.navigate(['chat']);

    }, error => { });
  }

  login() {
    this.loginService.login({ email: this.email, password: this.password }).subscribe(user => {
      console.log(user)
      console.log(":")
      console.log("password correct");
      this.router.navigate(['chat']);
    }, error => this.handleError(error));
  }

  handleError(error: any) {
    this.loginFailed = true;
    console.log(error);
  }

}
