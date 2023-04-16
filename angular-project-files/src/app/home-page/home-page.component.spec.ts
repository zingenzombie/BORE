import { ComponentFixture, TestBed } from '@angular/core/testing';
import { UsernameComponent } from '../username/username.component';
import { HomePageComponent } from './home-page.component';
import { By } from '@angular/platform-browser';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { UserService } from '../services/user.service';
import { FormsModule } from '@angular/forms';

describe('HomePageComponent', () => {
  let component: HomePageComponent;
  let fixture: ComponentFixture<HomePageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HomePageComponent, UsernameComponent, ],
      providers: [UserService],
      imports: [HttpClientTestingModule, FormsModule]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HomePageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should render a form with an upload button', () => {
    const formElement = fixture.nativeElement.querySelector('form');
    const fileInputElement = fixture.nativeElement.querySelector('input[type=file]');
    const submitButtonElement = fixture.nativeElement.querySelector('input[type=submit]');
    
    expect(formElement).toBeTruthy();
    expect(fileInputElement).toBeTruthy();
    expect(submitButtonElement).toBeTruthy();
    expect(submitButtonElement.value).toEqual('upload');
  });

  it('should contain the form element', () => {
    const formElement = fixture.debugElement.query(By.css('form'));
    expect(formElement).toBeTruthy();
  });
  
  it('the form element should be set to the correct encoding type', () => {
    const formElement = fixture.nativeElement.querySelector('#testForm');
    expect(formElement.getAttribute('enctype')).toEqual('multipart/form-data');
  });

  it('the form element should be set to the correct action URL', () => {
    const formElement = fixture.nativeElement.querySelector('#testForm');
  expect(formElement.getAttribute('action')).toEqual('http://localhost:3621/upload');
  });

  it('the form element should be set to the correct method', () => {
    const formElement = fixture.nativeElement.querySelector('#testForm');
    expect(formElement.getAttribute('method')).toEqual('post');
  });

  it('should contain the input text element', () => {
    const inputTextElement = fixture.debugElement.query(By.css('.text'));
    expect(inputTextElement).toBeTruthy();
  });
  
  it('input element type should be "file"', () => {
    const fileInputElement = fixture.nativeElement.querySelector('input[type=file]');
    expect(fileInputElement).toBeTruthy();
  });

  it('should contain the submit button element', () => {
    const submitButtonElement = fixture.debugElement.query(By.css('.button'));
    expect(submitButtonElement).toBeTruthy();
  });

  it('the submit button element should have the text "upload"', () => {
    const submitButtonElement = fixture.nativeElement.querySelector('input[type=submit]');
    expect(submitButtonElement.value).toEqual('upload');
  });
  
  it('should increase the file upload counter when a file is uploaded', () => {
    const formElement = fixture.nativeElement.querySelector('#testForm');
    const inputElement = fixture.nativeElement.querySelector('input[type=file]');
    spyOn(component, 'onUpload').and.callThrough();
  
    // 1st file
    inputElement.dispatchEvent(new Event('change'));
    fixture.detectChanges();
    formElement.dispatchEvent(new Event('submit'));
    fixture.detectChanges();
    expect(component.numFilesUploaded).toEqual(1);
  
    // 2nd file
    inputElement.dispatchEvent(new Event('change'));
    fixture.detectChanges();
    formElement.dispatchEvent(new Event('submit'));
    fixture.detectChanges();
    expect(component.numFilesUploaded).toEqual(2);
  
    // 3rd file
    inputElement.dispatchEvent(new Event('change'));
    fixture.detectChanges();
    formElement.dispatchEvent(new Event('submit'));
    fixture.detectChanges();
    expect(component.numFilesUploaded).toEqual(3);
  });
  
});
