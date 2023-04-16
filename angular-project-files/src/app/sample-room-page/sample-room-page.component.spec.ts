import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SampleRoomPageComponent } from './sample-room-page.component';
import { FormsModule } from '@angular/forms';

describe('SampleRoomPageComponent', () => {
  let component: SampleRoomPageComponent;
  let fixture: ComponentFixture<SampleRoomPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FormsModule],
      declarations: [SampleRoomPageComponent],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SampleRoomPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should set submitted to true on form submit', () => {
    component.onSubmit('{"pass": "password"}');
    expect(component.submitted).toBeTrue();
  });
  
  it('should set incorrect to true and correct to false if password is incorrect', () => {
    component.onSubmit('{"pass": "wrongpassword"}');
    expect(component.correct).toBeFalse();
    expect(component.incorrect).toBeTrue();
  });
});
