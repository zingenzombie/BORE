import { DebugElement } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule } from '@angular/forms';
import { User } from './user';

import { UsernameComponent } from './username.component';

// describe('UsernameComponent', () => {
//   let component: UsernameComponent;
//   let fixture: ComponentFixture<UsernameComponent>;
//   let de: DebugElement;

//   beforeEach(async () => {
//     TestBed.configureTestingModule({
//       imports: [FormsModule],
//       declarations: [UsernameComponent]
//     })
//     .compileComponents();
//   });

//   beforeEach(() => {
//     fixture = TestBed.createComponent(UsernameComponent);
//     // test component directly
//     component = fixture.componentInstance;
//     // test rendered html
//     de = fixture.debugElement;

//     fixture.detectChanges();
//   });
  
//   // username component is created
//   it('should create', () => {
//     expect(component).toBeTruthy();
//   });

//   it('should update submitted status to accepted', () => {
//     expect(component.submitted).toBeFalsy();
//     component.onSubmit(username: s);
//     expect(component.submitted).toBeTruthy();
//   });

  
// });
