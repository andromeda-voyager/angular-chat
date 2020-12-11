import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  @Input() password: string = "";
  isPasswordCorrect = true;
  hide = true;
  constructor() { }

  ngOnInit(): void {
  }

}
