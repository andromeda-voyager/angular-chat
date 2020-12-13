import { Directive, ElementRef } from '@angular/core';

@Directive({
  selector: '[isFocused]'
})
export class IsFocusedDirective {

  constructor(private element: ElementRef) { }

  ngAfterViewInit() {
    this.element.nativeElement.focus();
  }

}
