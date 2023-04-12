import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FlexboxtriangleroomComponent } from './flexboxtriangleroom.component';

describe('FlexboxtriangleroomComponent', () => {
  let component: FlexboxtriangleroomComponent;
  let fixture: ComponentFixture<FlexboxtriangleroomComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FlexboxtriangleroomComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FlexboxtriangleroomComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
