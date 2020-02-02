import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {NotificationService} from '../../../core/notifications/notification.service';
import {AccountService} from '../account.service';

@Component({
  selector: 'app-activate',
  templateUrl: './activate.component.html',
  styleUrls: ['./activate.component.scss']
})
export class ActivateComponent implements OnInit {
  message: string;

  constructor(private route: ActivatedRoute,
              private router: Router,
              private notificationService: NotificationService,
              private accountService: AccountService) {
  }

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      const activationKey = params.key;
      if (activationKey) {
        this.accountService.activate(activationKey).subscribe(
          resp => {
            this.notificationService.success('Account activated successfully');
            this.router.navigate(['account', 'login']);
          }
        );
      } else {
        this.message = 'Activation key is missing. Did you click the right link?' +
         ' Please contact support if you think this could be an error.';
      }
    });
  }


}
