# Manuka Honey UI

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 8.3.24.

## Development server

Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The app will automatically reload if you change any of the source files.

## Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

## Build

Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory. Use the `--prod` flag for a production build.

## Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).

## Running end-to-end tests

Run `ng e2e` to execute the end-to-end tests via [Protractor](http://www.protractortest.org/).

##GUI / UI testing
Login to Linux server and Go to folder 
/home/ubuntu/Honey_Jar_Final_UI/src/app

FOR UI, change IP address at:
	/home/ubuntu/Honey_Jar_Final_UI/src/app/shared/rest-api.service.ts

Type below command
ng serve --host 0.0.0.0
•	This Angular command will invoke UI. At end you will get below output
** Angular Live Development Server is listening on 0.0.0.0:4200, open your browser on http://localhost:4200/ **
ℹ ｢wdm｣: Compiled successfully.


Go to /home/ubuntu/MH-Project/MHasset-transfer, type command 
node server.js

Now open chrome browser and type say url
e.g 	http://13.233.82.179:4200/jardata
Angular UI will be open and then input JAR id in search tab. Search result will get JAR details from Blockchain server

