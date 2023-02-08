import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CircleroomComponent } from './circleroom.component';

describe('CircleroomComponent', () => {
  let component: CircleroomComponent;
  let fixture: ComponentFixture<CircleroomComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CircleroomComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CircleroomComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
