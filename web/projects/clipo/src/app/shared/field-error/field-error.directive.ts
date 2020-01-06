import {Directive, DoCheck, ElementRef, Input} from '@angular/core';
import {FormControl, ValidationErrors} from '@angular/forms';

@Directive({
  selector: '[appFieldError]'
})
export class FieldErrorDirective implements DoCheck {

  @Input()
  control: FormControl;

  constructor(
    private el: ElementRef
  ) {
  }

  ngDoCheck() {
    this.executeLogic();
  }

  private executeLogic(): void {
    this.el.nativeElement.innerText = this.errorMessage;
  }

  get errorMessage() {
    for (const propertyName in this.control.errors) {
      if (this.control.errors.hasOwnProperty(propertyName)) {
        return this.getValidatorErrorMessage(propertyName, this.control.errors[propertyName]);
      }
    }
    return null;
  }

  get showError() {
    return this.control.enabled && this.control.invalid && this.control.touched;
  }

  getValidatorErrorMessage(propertyName: string, error: ValidationErrors) {
    const config = {
      required: 'This field is required',
      email: 'Must be a valid email address',
      invalidPassword: 'Password must be at least 6 characters long, and contain a number.',
      minLength: `Minimum length ${error.requiredLength}`,
      maxLength: `Maximum length ${error.maxLength}`,
    };

    // @ts-ignore
    return config[propertyName];
  }

}

