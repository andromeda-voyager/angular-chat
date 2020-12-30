import { Component, OnInit, Input, ViewChild, ElementRef } from '@angular/core';
declare var MediaRecorder: any;
@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.scss']
})
export class ChatComponent implements OnInit {

  dialogOpen: boolean = false;

  isCameraOpen: boolean = false;
  videoStream: MediaStream = new MediaStream();
  videoFeedStream: MediaStream = new MediaStream();
  mediaSource: MediaSource = new MediaSource();
  image: string = "";
  //videoFeedStream = URL.createObjectURL(this.mediaSource);
  constructor() { }

  ngOnInit(): void {


  }

  closeDialog() {
      this.dialogOpen =false;
    
    console.log("closing dialogs")
  }

  addServer() {
    console.log("opening dialogs")
    this.dialogOpen = true;
  }

  takePicture() {
   var track = this.videoStream.getVideoTracks()[0];
    this.videoFeedStream.addTrack(track);
    //this.image = new ImageCapture(this.track);
  }

  openCamera() {
    this.isCameraOpen = true;
    navigator.mediaDevices.getUserMedia({ video: true, audio: false })
      .then((stream) => {
        this.videoStream = stream;
        //this.video.play();
      })
      .catch(function (err) {
        console.log("An error occurred: " + err);
      });
  }

}
