# Introduction

## Angular What and Why
* orginal version called AngularJS
* cuurent version Angular
* invented by Misko Hevery
* Angular 2
    * standards based
    * performant

## Hello Angular
* installs all code scaffold
* installs npm package
* can generate components and run the application

## Versions of Angular
* never Angular 3 becuase of the router
* release major version every 6 months, spring or fall
* 18 months of LTS support then you have to update
* ng update

# Benefits & Features of Angular

## Universal Benefits
* Reduction of cost
* Standards Complicance
    * work on more browers
    * modules
    * es6
    * accessible
    * i18n
* Open license
* Popular


## Subjective Benefits
* router,http,forms,rxjs
* uniformity
* backed by Google
* typescript

## Basic Features
* PWA - web apps that can be installed on the phone
* Lazy loading 
* Forms
* RXJS
* Router
* Animations

## Advanced Features
* Server Side Rendering  - search engines server create html & css
* Mobile
* Angular Language
* ngUpgrade - help migrate from AngularJS to Angular

# Angular Architecture

## One way data flow
* change detection begins in the parent and cascades through children
* the state is not allowed to update until change detection completes
* big performance boost

## Directives
* extend your element

## Zone.js & Change Detection
* state changes - HTTP ,Events
* once state change (First Name) cascading change (full name) - rerender 
* zone.js creates a wrapper around all async opearations in the brrowser
* angular subscribes to zone when any operation completes
* zone.js waits for direct and cascading changes then notifies angular to rerender

## Rendering Targets
*   angular can go on any device by changing the rendering engine
* for browser its
    * browser-platform
    * browser-platform-dynamic

# Tooling
* javascript fatigue amount of time it takes before you can start writing actual code
* there is modules, bundling, zone.js

## Server -Side Rendering
* smaller download size
* just gives you html
* SEO
* Full- Pre-render - create html for each view take html and load on cdn
* Dynamic Pre-render - request page built, sent to browser, 
* builds angular in a hidden div, replays events to keep state then swaps 

## Mobile & Native Frameworks
* mobile, ionic,
* desktop - electron

## Testing Tools
* Karma was built by angular
* TestBed - test components with template
* mockBackend - mock http server

## AOT Compiler
* radically improve performaces
* complies templates on the server the time saver

# Tips,Tricks & Gotachs

## Gotchas
* typescript decorators
* pure pipes run when input changes
* impure pipes run on every change detection
* cryptic error messages
* ng build
* delivery size

## Tips
* follow the style guide on angular
* do sorting & filtering in the component
* ngrx
* webpack
* dont touch the Dom directly
* understand package.json

# Angular Present & Futre
* make it smooth to upgrade to AngularJS

## vs other frameworks
* politcics
* personel
* skill sets
* typed backedn

## Future
* Angular Elements - turns your code in to  web component to be used in any frameworks
* 