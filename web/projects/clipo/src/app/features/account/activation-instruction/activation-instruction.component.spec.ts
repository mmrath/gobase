import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ActivationInstructionComponent } from './activation-message.component';

describe('ActivationMessageComponent', () => {
  let component: ActivationInstructionComponent;
  let fixture: ComponentFixture<ActivationInstructionComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ActivationInstructionComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ActivationInstructionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
