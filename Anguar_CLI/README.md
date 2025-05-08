# Course Overview

# Generating a New Angular Application
* 
## Angular CLI JSON file
* chnage the prefix of components with the angular.json prefix flag
* strict modes in template will start to error in your html

## Configuting the CLI
* can use ng config to write/read angular.json

# Generating Code From Blueprints

## Blueprints
* place classes, interfaces, enum in the models folder
* place pipes in the shared folder

## Building and Serving

runtime.js - webpack runtime
main.js - app code
polyfills.js - platform polffills
styles.s - styles
vendor.js- angular and other vendors

## Building Target and Environment
* cache busting important when you want to replace images

## Adding new capabilities
* leverage ng add for angular's new features

## Code Coverage
* what lines of code we are actually exercising in our tests

## Workspaces
* angular has workspaces

## Generating Libraries
npm ng g library
cd dist/my-lib
npm publish

## Creating a Logger Library
* make the lib
* ng build my-logger --prod
* make sure the all the classes are in the module
* make sure the files are exported from public-api.ts
* and the library to the app.module