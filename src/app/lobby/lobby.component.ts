import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-lobby',
  templateUrl: './lobby.component.html',
  styleUrls: ['./lobby.component.scss']
})
export class LobbyComponent implements OnInit {

  rooms: string[] = [];
  constructor() { }

  ngOnInit(): void {
    this.rooms.push("room 1");
    this.rooms.push("room 2");
  }

}
