import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ResetPasswordInitComponent } from './reset-password-init.component';

describe('ResetPasswordInitComponent', () => {
  let component: ResetPasswordInitComponent;
  let fixture: ComponentFixture<ResetPasswordInitComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ResetPasswordInitComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ResetPasswordInitComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
