import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SampleRoomPageComponent } from './sample-room-page.component';

describe('SampleRoomPageComponent', () => {
  let component: SampleRoomPageComponent;
  let fixture: ComponentFixture<SampleRoomPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SampleRoomPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SampleRoomPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
