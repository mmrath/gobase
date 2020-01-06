import {Injectable} from '@angular/core';
import {RegisterAccountRequest} from './model';
import {HttpClient} from '@angular/common/http';
import {NotificationService} from '../../core/core.module';
import {catchError, tap} from "rxjs/operators";


@Injectable({
  providedIn: 'root'
})
export class AccountService {

  constructor(
    private httpClient: HttpClient,
    private notificationService: NotificationService) {
  }

  register(data: RegisterAccountRequest) {
    const url = '/api/account/register';
    return this.httpClient
    .post(url, data)
    .pipe(
      tap(res => console.log(res)),
      catchError(this.notificationService.handleError)
    );
  }
}
