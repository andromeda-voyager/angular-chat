import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ChatComponent } from './chat/chat.component';
import { PostsComponent } from './messages/messages.component';
import { LoginComponent } from './login/login.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button'
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { IsFocusedDirective } from './shared/is-focused.directive';
import { HttpClientModule } from '@angular/common/http';
import { CreateAccountComponent } from './create-account/create-account.component';
import { AddServerComponent } from './add-server/add-server.component';
import { AdminComponent } from './admin/admin.component';
import { AddChannelComponent } from './add-channel/add-channel.component';
import { InputFieldComponent } from './shared/input-field/input-field.component';
import { ServerPanelComponent } from './server-panel/server-panel.component';

@NgModule({
  declarations: [
    AppComponent,
    ChatComponent,
    PostsComponent,
    LoginComponent,
    IsFocusedDirective,
    CreateAccountComponent,
    AddServerComponent,
    AdminComponent,
    AddChannelComponent,
    InputFieldComponent,
    ServerPanelComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    BrowserAnimationsModule,
    MatFormFieldModule,
    MatButtonModule,
    MatIconModule,
    MatInputModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
