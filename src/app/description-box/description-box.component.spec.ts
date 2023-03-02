import { DebugElement } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DescriptionBoxComponent } from './description-box.component';

describe('DescriptionBoxComponent', () => {
  let component: DescriptionBoxComponent;
  let fixture: ComponentFixture<DescriptionBoxComponent>;
  let de: DebugElement;

  beforeEach(async () => {
    TestBed.configureTestingModule({
      declarations: [DescriptionBoxComponent]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DescriptionBoxComponent);
    // test component directly
    component = fixture.componentInstance;
    // test rendered html
    de = fixture.debugElement;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should toggle the description boolean', () => {
     expect(component.showDescription).toBeFalsy();
     component.toggle();
     expect(component.showDescription).toBeTruthy();
  });
})
