import {AfterViewInit, Component, ContentChild, ElementRef} from '@angular/core';
import {FormArrayName} from '@angular/forms';
import {FormValidationContainer} from './form-validation-container';

@Component({
  // tslint:disable:component-selector
  selector: '[formArrayContainer], form-array-container',
  template: `
    <ng-content></ng-content>
    <ng-container #errorsContainer></ng-container>
  `
})
export class FormArrayContainerComponent extends FormValidationContainer implements AfterViewInit {

  // tslint:disable-next-line:variable-name
  @ContentChild(FormArrayName) _formControl: FormArrayName;

  // tslint:disable-next-line:variable-name
  @ContentChild(FormArrayName, {read: ElementRef}) _el: ElementRef;


  get formControl() {
    return this._formControl.control;
  }

  get formControlName(): string {
    if (this._formControl.name && typeof this._formControl.name === 'string') {
      return this._formControl.name;
    } else {
      throw new Error('expected _formControl.name an instance of string');
    }
  }

  get el(): ElementRef {
    return this._el;
  }
}
