import { Component, OnInit, Input, ViewChild, ElementRef } from '@angular/core';
import { ChatService } from '../shared/chat.service';
import { Connection, Server } from '../shared/server';
import { Account } from '../shared/account';
import { LoginService } from '../shared/login.service';

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
  connections: Connection[] = []
  image: string = "";
 
  //videoFeedStream = URL.createObjectURL(this.mediaSource);
  constructor(private chatService: ChatService, private loginService: LoginService) { }

  ngOnInit(): void {
    let account = this.loginService.getUserData()
    this.connections = account.connections;
    console.log(this.connections[0].server.imageURL)
    
  }

  closeDialog() {
    this.dialogOpen = false;

    console.log("closing dialogs")
  }

  onNewConnection(connection: Connection) {
    if (this.connections == null) {
      this.connections = [];
    }
    this.connections.push(connection);
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
