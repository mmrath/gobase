import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ChangePasswordComponent } from './change-password/change-password.component';
import { ProfileComponent } from './profile/profile.component';
import { ResetPasswordInitComponent } from './reset-password-init/reset-password-init.component';
import { ResetPasswordComponent } from './reset-password/reset-password.component';
import { SignUpComponent } from './sign-up/sign-up.component';



@NgModule({
  declarations: [ChangePasswordComponent, ProfileComponent, ResetPasswordInitComponent, ResetPasswordComponent, SignUpComponent],
  imports: [
    CommonModule
  ]
})
export class AccountModule { }
