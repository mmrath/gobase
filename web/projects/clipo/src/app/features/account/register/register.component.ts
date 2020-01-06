import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {AccountService} from '../account.service';
import {Router} from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent implements OnInit {

  form: FormGroup;
  submitSuccess: boolean;

  constructor(private fb: FormBuilder,
              private accountService: AccountService,
              private router: Router,) {
    this.form = this.fb.group({
      firstName: ['', [Validators.required, Validators.maxLength(32)]],
      lastName: ['', [Validators.required, Validators.maxLength(32)]],
      email: ['', [Validators.required, Validators.email, Validators.maxLength(32)]],
      password: ['', [Validators.required, Validators.minLength(6), Validators.maxLength(20)]],
    });
    this.submitSuccess = false;
  }

  ngOnInit() {
  }

  register() {
    if (this.form.invalid) {
      return;
    }

    this.accountService.register(this.form.value)
    .subscribe(_ => {
      this.router.navigate(['account', 'activation-instruction'])
      .then(r => {
        // do any unsubscription here
      });
    });
  }
}
