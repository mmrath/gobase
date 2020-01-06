import {AfterViewInit, Component, ContentChild, ElementRef} from '@angular/core';
import {ControlContainer, FormControlName} from '@angular/forms';
import {FormValidationContainer} from './form-validation-container';

@Component({
  // tslint:disable:component-selector
  selector: '[formFieldContainer], form-field-container',
  template: `
      <ng-content></ng-content>
      <ng-container #errorsContainer></ng-container>
  `
})
export class FormFieldContainerComponent extends FormValidationContainer implements AfterViewInit {

  // tslint:disable-next-line:variable-name
  @ContentChild(FormControlName) _formControl: FormControlName;

  // tslint:disable-next-line:variable-name
  @ContentChild(FormControlName, {read: ElementRef}) _input: ElementRef;

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

  get el(): ElementRef<any> {
    return this._input;
  }
}
