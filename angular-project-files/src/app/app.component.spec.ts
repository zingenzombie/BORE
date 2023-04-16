import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { AppComponent } from './app.component';
import { UserService } from './services/user.service';
import { HttpClientModule } from '@angular/common/http';
import { DescriptionBoxComponent } from './description-box/description-box.component';
import { By } from '@angular/platform-browser';
import { Router, RouterOutlet } from '@angular/router';
import { HomePageComponent } from './home-page/home-page.component';
import { SampleRoomPageComponent } from './sample-room-page/sample-room-page.component';

describe('AppComponent', () => {
  let component: AppComponent;
  let fixture: ComponentFixture<AppComponent>;
  let userService: UserService;
  let router: Router;


  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [AppComponent, DescriptionBoxComponent,],
      imports: [HttpClientModule,
        RouterTestingModule.withRoutes([
          { path: 'home', component: HomePageComponent },
          { path: 'room-page-1', component: SampleRoomPageComponent },
          { path: 'room-page-2', component: SampleRoomPageComponent },
          { path: 'room-page-3', component: SampleRoomPageComponent }
        ]),
        RouterOutlet],
      providers: [UserService]
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AppComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    router = TestBed.inject(Router);
    spyOn(router, 'navigateByUrl');
    userService = TestBed.inject(UserService);
  });

  it('should create the app', () => {
    expect(component).toBeTruthy();
  });

  it('should display the title and subtitle', () => {
    component.title = 'Test Title';
    component.subTitle = 'Test Subtitle';
    fixture.detectChanges();
    const titleElement = fixture.nativeElement.querySelector('.title');
    const subTitleElement = fixture.nativeElement.querySelector('.subTitle');
    expect(titleElement.textContent).toContain('Test Title');
    expect(subTitleElement.textContent).toContain('Test Subtitle');
  });

  it('should display the username when the user is logged in', () => {
    spyOn(userService, 'getUsername').and.returnValue('Test User');
    fixture.detectChanges();
    const usernameElement = fixture.nativeElement.querySelector('p');
    expect(usernameElement.textContent).toContain('Hello Test User!');
  });

  it('should not display the username when the user is not logged in', () => {
    spyOn(userService, 'getUsername').and.returnValue('');
    fixture.detectChanges();
    const usernameElement = fixture.nativeElement.querySelector('p');
    expect(usernameElement).toBeNull();
  });

  it('should contain a link to Home page', () => {
    const homeLink = fixture.debugElement.query(By.css('a[href="/home"]'));
    expect(homeLink).toBeTruthy();
  });

  it('should display the navigation links', () => {
    const homeLink = fixture.debugElement.query(By.css('a[href="/home"]')).nativeElement;
    const room1Link = fixture.debugElement.query(By.css('a[href="/room-page-1"]')).nativeElement;
    const room2Link = fixture.debugElement.query(By.css('a[href="/room-page-2"]')).nativeElement;
    const room3Link = fixture.debugElement.query(By.css('a[href="/room-page-3"]')).nativeElement;

    expect(homeLink.textContent).toContain('Home');
    expect(room1Link.textContent).toContain('Private Room: Circle');
    expect(room2Link.textContent).toContain('Private Room: Square');
    expect(room3Link.textContent).toContain('Private Room: Triangle');
  });

  it('should contain a link to Circle Room page', () => {
    const room1Link = fixture.debugElement.query(By.css('a[href="/room-page-1"]'));
    expect(room1Link).toBeTruthy();
  });

  it('should contain a link to Square Room page', () => {
    const room2Link = fixture.debugElement.query(By.css('a[href="/room-page-2"]'));
    expect(room2Link).toBeTruthy();
  });

  it('should contain a link to Triangle Room page', () => {
    const room3Link = fixture.debugElement.query(By.css('a[href="/room-page-3"]'));
    expect(room3Link).toBeTruthy();
  });


});