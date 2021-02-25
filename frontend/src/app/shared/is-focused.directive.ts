import { Directive, ElementRef, Input, OnChanges } from '@angular/core';

@Directive({
  selector: '[isFocused]'
})
export class IsFocusedDirective implements OnChanges{

  @Input('isFocused') isFocused:boolean = false;
  constructor(private element: ElementRef) { }

  //ngAfterViewInit() {
    ngOnChanges() {
      if(this.isFocused) {
        this.element.nativeElement.focus();
      }
  }

}
