import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FlexboxsquareroomComponent } from './flexboxsquareroom.component';

describe('FlexboxsquareroomComponent', () => {
  let component: FlexboxsquareroomComponent;
  let fixture: ComponentFixture<FlexboxsquareroomComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FlexboxsquareroomComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FlexboxsquareroomComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
