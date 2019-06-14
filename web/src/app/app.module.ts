import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';

import {SharedModule} from '@app/shared';
import {CoreModule} from '@app/core';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {SettingsModule} from "@app/settings";

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    AppRoutingModule,

    // core & shared
    CoreModule,
    SharedModule,

    SettingsModule,


  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {
}
