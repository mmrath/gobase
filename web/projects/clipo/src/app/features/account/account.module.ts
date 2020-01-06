import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {ChangePasswordComponent} from './change-password/change-password.component';
import {ProfileComponent} from './profile/profile.component';
import {ResetPasswordInitComponent} from './reset-password-init/reset-password-init.component';
import {ResetPasswordComponent} from './reset-password/reset-password.component';
import {SharedModule} from '../../shared/shared.module';
import {AccountRoutingModule} from './account-routing.module';
import {LoginComponent} from './login/login.component';
import {RegisterComponent} from './register/register.component';
import {ActivationInstructionComponent} from './activation-instruction/activation-instruction.component';



@NgModule({
  declarations: [
    ChangePasswordComponent,
    ProfileComponent,
    ResetPasswordInitComponent,
    ResetPasswordComponent,
    LoginComponent,
    RegisterComponent,
    ActivationInstructionComponent,
  ],
  imports: [
    CommonModule,
    SharedModule,
    AccountRoutingModule
  ]
})
export class AccountModule {
}
