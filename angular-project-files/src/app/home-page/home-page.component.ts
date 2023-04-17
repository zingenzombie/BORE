import { Component } from '@angular/core';
import { UsernameComponent } from '../username/username.component';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})


export class HomePageComponent {
  numFilesUploaded = 0;
  fileList: any[] = [];

  onUpload(event: any) {
    this.numFilesUploaded++;
    const fileInput = event.target.querySelector('input[type="file"]');
    const uploadTime = new Date();
    const files = fileInput.files;
    for (let i = 0; i < files.length; i++) {
      const file = files[i];
      this.fileList.push({name: file.name, time: uploadTime});
    }
  }

  
}
