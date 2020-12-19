import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ChatComponent } from './chat/chat.component';
import { LobbyComponent } from './lobby/lobby.component';
import { LoginComponent } from './login/login.component';

const routes: Routes = [
  { path: 'chat', component: ChatComponent },
  { path: 'lobby', component: LobbyComponent },
  { path: 'login', component: LoginComponent },
  { path: '', redirectTo: '/lobby', pathMatch:'full'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
