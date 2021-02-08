import { Injectable } from '@angular/core';
import { Observable, Observer, Subject } from 'rxjs';
import { Message, NewMessage } from '../models/message';

@Injectable({
  providedIn: 'root'
})
export class MessageService {

  private newMessageSource = new Subject<Message>();
  newMessage$ = this.newMessageSource.asObservable();

  private modifyMessageSource = new Subject<Message>();
  modifyMessage$ = this.modifyMessageSource.asObservable();

  private deleteMessageSource = new Subject<Message>();
  deleteMessage$ = this.deleteMessageSource.asObservable();

  constructor() { }
}
