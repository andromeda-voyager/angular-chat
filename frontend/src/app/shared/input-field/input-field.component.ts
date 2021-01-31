import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

@Component({
  selector: 'app-input-field',
  templateUrl: './input-field.component.html',
  styleUrls: ['./input-field.component.scss']
})
export class InputFieldComponent implements OnInit {

  @Input() label = "";
  @Input() hint = "";
  @Input() inputValue: string = "";
  @Input() isPassword: boolean = false;
  @Input() isDark = false;

  @Output() inputValueChange = new EventEmitter<string>();
  @Input() showHint = false;
  isFocused = false;
  type = this.isPassword ? "password" : "text"

  constructor() {
  }

  ngOnInit(): void {
  }


}
