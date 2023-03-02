import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CircleRoomComponent } from './circle-room.component';

describe('CircleRoomComponent', () => {
  let component: CircleRoomComponent;
  let fixture: ComponentFixture<CircleRoomComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CircleRoomComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CircleRoomComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
