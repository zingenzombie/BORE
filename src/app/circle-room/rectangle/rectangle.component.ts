import { Component, Input } from '@angular/core';
@Component({
  selector: 'app-rectangle',
  templateUrl: './rectangle.component.html',
  styleUrls: ['./rectangle.component.css']
})
export class RectangleComponent {
  showRectangle = false;

  users = ['User 1', 'User 2', 'User 3'];
}
