import { Component, OnInit, Input,ViewChild, ElementRef} from '@angular/core';

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.scss']
})
export class ChatComponent implements OnInit {

  isCameraOpen: boolean = false;
  videoStream:MediaStream = new MediaStream();
  
  constructor() { }

  ngOnInit(): void { }

  openCamera() {
    this.isCameraOpen = true;
    navigator.mediaDevices.getUserMedia({ video: true, audio: false })
    .then((stream) => {
        this.videoStream = stream;
       //this.video.play();
    })
    .catch(function(err) {
        console.log("An error occurred: " + err);
    });
  }
  
}
