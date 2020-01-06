import {CommonModule} from '@angular/common';
import {ChangeDetectorRef, ErrorHandler, NgModule, Optional, SkipSelf} from '@angular/core';
import {HTTP_INTERCEPTORS, HttpClient, HttpClientModule} from '@angular/common/http';
import {RouterStateSerializer, StoreRouterConnectingModule} from '@ngrx/router-store';
import {StoreModule} from '@ngrx/store';
import {EffectsModule} from '@ngrx/effects';
import {StoreDevtoolsModule} from '@ngrx/store-devtools';
import {
  TranslateLoader,
  TranslateModule,
  TranslatePipe,
  TranslateService
} from '@ngx-translate/core';
import {TranslateHttpLoader} from '@ngx-translate/http-loader';

import {environment} from '../../environments/environment';

import {AppState, metaReducers, reducers, selectRouterState} from './core.state';
import {AuthEffects} from './auth/auth.effects';
import {selectAuth, selectIsAuthenticated} from './auth/auth.selectors';
import {authLogin, authLogout} from './auth/auth.actions';
import {AuthGuardService} from './auth/auth-guard.service';
import {TitleService} from './title/title.service';
import {ROUTE_ANIMATIONS_ELEMENTS, routeAnimations} from './animations/route.animations';
import {AnimationsService} from './animations/animations.service';
import {AppErrorHandler} from './error-handler/app-error-handler.service';
import {CustomSerializer} from './router/custom-serializer';
import {LocalStorageService} from './local-storage/local-storage.service';
import {HttpErrorInterceptor} from './http-interceptors/http-error.interceptor';
import {GoogleAnalyticsEffects} from './google-analytics/google-analytics.effects';
import {NotificationService} from './notifications/notification.service';
import {SettingsEffects} from './settings/settings.effects';
import {
  selectEffectiveTheme,
  selectSettingsLanguage,
  selectSettingsStickyHeader
} from './settings/settings.selectors';
import {ValidationErrorsModule} from './validation-error/validation-errors.module';
import {MESSAGES_PIPE_FACTORY_TOKEN, MESSAGES_PROVIDER} from "./validation-error/tokens";

export {
  TitleService,
  selectAuth,
  authLogin,
  authLogout,
  routeAnimations,
  AppState,
  LocalStorageService,
  selectIsAuthenticated,
  ROUTE_ANIMATIONS_ELEMENTS,
  AnimationsService,
  AuthGuardService,
  selectRouterState,
  NotificationService,
  selectEffectiveTheme,
  selectSettingsLanguage,
  selectSettingsStickyHeader
};

export function HttpLoaderFactory(http: HttpClient) {
  return new TranslateHttpLoader(
    http,
    `${environment.i18nPrefix}/assets/i18n/`,
    '.json'
  );
}

export function translatePipeFactoryCreator(translateService: TranslateService) {
  return (detector: ChangeDetectorRef) => new TranslatePipe(translateService, detector);
}

@NgModule({
  imports: [
    // angular
    CommonModule,
    HttpClientModule,

    // ngrx
    StoreModule.forRoot(reducers, {metaReducers}),
    StoreRouterConnectingModule.forRoot(),
    EffectsModule.forRoot([
      AuthEffects,
      SettingsEffects,
      GoogleAnalyticsEffects
    ]),
    environment.production
      ? []
      : StoreDevtoolsModule.instrument({
        name: 'Angular NgRx Material Starter'
      }),

    // 3rd party
    TranslateModule.forRoot({
      loader: {
        provide: TranslateLoader,
        useFactory: HttpLoaderFactory,
        deps: [HttpClient]
      }
    }),
    ValidationErrorsModule.forRoot()
  ],
  declarations: [],
  providers: [
    {provide: HTTP_INTERCEPTORS, useClass: HttpErrorInterceptor, multi: true},
    {provide: ErrorHandler, useClass: AppErrorHandler},
    {provide: RouterStateSerializer, useClass: CustomSerializer},
    {
      provide: MESSAGES_PIPE_FACTORY_TOKEN,
      useFactory: translatePipeFactoryCreator,
      deps: [TranslateService]
    },
    {
      provide: MESSAGES_PROVIDER,
      useExisting: TranslateService
    }
  ],
  exports: [TranslateModule]
})
export class CoreModule {
  constructor(
    @Optional()
    @SkipSelf()
      parentModule: CoreModule
  ) {
    if (parentModule) {
      throw new Error('CoreModule is already loaded. Import only in AppModule');
    }
  }
}
