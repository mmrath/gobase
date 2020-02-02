import {RouterModule, Routes} from '@angular/router';
import {AuthGuardService} from '../../core/auth/auth-guard.service';
import {NgModule} from '@angular/core';
import {ProfileComponent} from './profile/profile.component';
import {ChangePasswordComponent} from './change-password/change-password.component';
import {LoginComponent} from './login/login.component';
import {RegisterComponent} from './register/register.component';
import {ActivationInstructionComponent} from './activation-instruction/activation-instruction.component';
import {ActivateComponent} from './activate/activate.component';

const routes: Routes = [
  {
    path: '',
    children: [
      {
        path: '',
        redirectTo: 'profile',
        pathMatch: 'full'
      },
      {
        path: 'profile',
        component: ProfileComponent,
        canActivate: [AuthGuardService],
        data: {title: 'app.account.menu.profile'}
      },
      {
        path: 'login',
        component: LoginComponent,
        data: {title: 'app.account.menu.login'}
      },
      {
        path: 'register',
        component: RegisterComponent,
        data: {title: 'app.account.menu.register'}
      },
      {
        path: 'change-password',
        component: ChangePasswordComponent,
        canActivate: [AuthGuardService],
        data: {title: 'app.account.menu.change-password'}
      },
      {
        path: 'activation-instruction',
        component: ActivationInstructionComponent,
        data: {title: 'app.account.menu.activation-message'}
      },
      {
        path: 'activate',
        component: ActivateComponent,
        data: {title: 'app.account.menu.activate'}
      },
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AccountRoutingModule {
}
