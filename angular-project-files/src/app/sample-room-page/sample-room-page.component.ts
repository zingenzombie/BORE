import { Component } from '@angular/core';

@Component({
  selector: 'app-sample-room-page',
  templateUrl: './sample-room-page.component.html',
  styleUrls: ['./sample-room-page.component.css']
})
export class SampleRoomPageComponent {
  submitted: boolean = false;
  password!: string;
  incorrect: boolean = false;
  correct: boolean = false;
  tempStr!: string;
  tempObj!: Object;
  onSubmit(pass: string) {
    this.submitted = true;
    this.tempStr = JSON.stringify(pass);
    const tempObject = JSON.parse(this.tempStr);
    this.tempStr = tempObject.pass;
    if (this.tempStr === 'password') {
      this.correct = true;
      this.incorrect = false;
    } else {
      this.incorrect = true;
      this.correct = false;
    }
  }
}
