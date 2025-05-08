# Introduction
* expressive 
* powerful data binding
* modular

## Anatomy of An Angular Application
* set of components, services used by the component
* you can specify component metadata
* have checklist

## Sample Application
* we can have resuable components such as a typeahead cpnt list-cpnt with detail cpnts

# Introduction to Component

## Bootstrap Our App Component
* bootstrap property in NgModule contains the startup component that index.html uses
* in the Sources pane our code is under the webpack folder  
* everything is case sensitive

## Checklist and Summary
* for properties in Components
    * appropriate data type
    * appropriate default value
    * camelCase
* for functions in Compones
    * camelCase

# Templates, Interpolation & Directives

## Introduction

## Building A Template
* import bootstrap css and font awesome css into styles
* we dont need jquery and popper

## Using a Component as a Directive
* Angular looks to the Module to find the selector for any Component

## Adding Logic with Directive
* structrual directive modify the dom based on logic

# Data Binding & Pipes
* propety binding - control the htl

## Handling Input with two-way bidning
* brackets for parent to child relationship
* parentheses to send and event which in this special case, updates the variable it it attacheed to listFilter
```html
[(ngModel)]="listFilter"
```

## Transforming Data with Pipes
parameters to a pipe
```html

{{ value | currency:'USD':'symbol':'1.2-2'}}
```

# More on Compoennts

## Defining Interfaces
* we can use interfaces as an important way to make sure components are created with a certain type
* extending a class with the implements keyword of that interface
* you should be able to implements with a class

## Lifecycle Hooks
* OnInit, to 
* OnChanges, perform action after change to input properties
* OnDestroy, clean up

# Building Nested Components


# Services and Dependecy Injection
* dependecny injection - recevies instances from external source rather than creating it itself

## Regirstering the service
* root injector, - service is available throughout the application
    * register with the Injectable Decorator of the service
* component injecter- service is avaialble only to that component and its children
    * use the componets provider proeprtis to register with a service   

## Injecting the Service
* class without a constructor use implicit constuctor

# Retrieving Data Using HTTP

# Navigation and Routing Basics


## Generating Code with the Angular CLI 


## Typing Routes to Actions
```html
<a [routerLink]="['/welcome',product.id]">
<a [routerLink]="['/products']">
```
* routing in intricate not encapsulated

# Navigation and Routing Additional Techniques

## Passing Paramerters to a Route
* to read emitters parameters as they change
```ts
this.route.paramMap.subscribe(
    params => console.log(params.get('id'))
)
```

## Protecting Routes with Guards
* CanActivate,CanDeactivate,Resolve,CanLoad
* Resolve - pre-fetch data before activating a route

canActivate
```ts
  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
    const id = Number(route.paramMap.get('id'));
    if (isNaN(id) || id < 1) {
      alert('Invalid product id');
      this.router.navigate(['/products']);
      return false;
    }
    return true;
  }
```

# Angular Modules

## Declarations Array
* every component,directive & pipe we make belongs ONLY to one angular module
* BrowserModule imports and export CommonModule

# Building Testing and Deploying with the CLI