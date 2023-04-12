import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SquareroomComponent } from './squareroom.component';

describe('SquareroomComponent', () => {
  let component: SquareroomComponent;
  let fixture: ComponentFixture<SquareroomComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SquareroomComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SquareroomComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
