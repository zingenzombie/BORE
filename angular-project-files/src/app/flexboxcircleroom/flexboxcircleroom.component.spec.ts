import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FlexboxcircleroomComponent } from './flexboxcircleroom.component';

describe('FlexboxcircleroomComponent', () => {
  let component: FlexboxcircleroomComponent;
  let fixture: ComponentFixture<FlexboxcircleroomComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FlexboxcircleroomComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FlexboxcircleroomComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
