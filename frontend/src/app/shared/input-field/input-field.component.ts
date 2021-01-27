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
  @Output() inputValueChange = new EventEmitter<string>();

  @Input() showHint = false;
  isFocused = false;
  constructor() { }

  ngOnInit(): void {
  }

}
