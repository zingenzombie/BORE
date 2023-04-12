import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FlextriangleroomComponent } from './flextriangleroom.component';

describe('FlextriangleroomComponent', () => {
  let component: FlextriangleroomComponent;
  let fixture: ComponentFixture<FlextriangleroomComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FlextriangleroomComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FlextriangleroomComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
