import { DebugElement } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule } from '@angular/forms';
import { User } from './user';
import {HttpClientTestingModule, HttpTestingController} from '@angular/common/http/testing'
import { UserService } from '../services/user.service';
import { UsernameComponent } from './username.component';

describe('UsernameService', () => {
  let service: UserService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [FormsModule, HttpClientTestingModule],
      providers: [UserService],
      declarations: [UsernameComponent]
    });
    service = TestBed.inject(UserService);
    httpMock = TestBed.inject(HttpTestingController);
    
  });

  afterEach(() => {
    httpMock.verify();
  })

  it ('should be created', () => {
    expect(service).toBeTruthy();
  })
  
  it('should set and get the username', () => {
    const username = 'John Doe';
    service.setUsername(username);
    expect(service.getUsername()).toEqual(username);
  });

  it('should post the username', () => {
    const username = 'Jane Doe';
    service.postUsername(username);
    const req = httpMock.expectOne('http://10.136.165.182:3621/setName');
    expect(req.request.method).toBe('POST');
    req.flush({});
  });
});

describe('UsernameComponent', () => {
  let component: UsernameComponent;
  let fixture: ComponentFixture<UsernameComponent>;
  let userService: UserService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UsernameComponent ],
      imports: [ HttpClientTestingModule, FormsModule ],
      providers: [ UserService ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(UsernameComponent);
    component = fixture.componentInstance;
    userService = TestBed.inject(UserService);
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should add a new user to the user list', () => {
    component.addUser('John');
    expect(component.userList).toContain('John');
  });

  it('should set the submitted flag to true', () => {
    component.onSubmit('John');
    expect(component.submitted).toBeTrue();
  });
  
  it('should call the postUsername method of the UserService', () => {
    spyOn(userService, 'postUsername');
    component.onSubmit('John');
    expect(userService.postUsername).toHaveBeenCalledWith('John');
  });
});

