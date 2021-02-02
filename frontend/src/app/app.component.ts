import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { LoginService } from './shared/services/login.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'chat';

  constructor(private loginService: LoginService) {

  }

  isLoggedIn(): boolean {
    return this.loginService.isLoggedIn();
  }

  logout() {
    this.loginService.logout();
  }
}
