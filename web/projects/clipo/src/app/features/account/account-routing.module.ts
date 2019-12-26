import {RouterModule, Routes} from '@angular/router';
import {AuthGuardService} from '../../core/auth/auth-guard.service';
import {NgModule} from '@angular/core';
import {ProfileComponent} from './profile/profile.component';
import {ChangePasswordComponent} from './change-password/change-password.component';

const routes: Routes = [
  {
    path: '',
    component: ProfileComponent,
    canActivate: [AuthGuardService],
    children: [
      {
        path: '',
        redirectTo: 'profile',
        pathMatch: 'full'
      },
      {
        path: 'profile',
        component: ProfileComponent,
        data: {title: 'app.account.menu.profile'}
      },
      {
        path: 'change-password',
        component: ChangePasswordComponent,
        data: {title: 'app.account.menu.change-password'}
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AccountRoutingModule {
}
