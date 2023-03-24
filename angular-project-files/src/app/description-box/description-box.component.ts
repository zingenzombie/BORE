import { Component } from '@angular/core';

@Component({
  selector: 'app-description-box',
  templateUrl: './description-box.component.html',
  styleUrls: ['./description-box.component.css']
})
export class DescriptionBoxComponent {
  showDescription = false;

  toggle() {
    this.showDescription = !this.showDescription;
  }
}
