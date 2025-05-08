# First Look At Ngrx
* inialize store
* im memory 
* runtine only not if users refreshes page or exits the application

image.png![1661531585256](image/README/1661531585256.png)
* you kinda of select a feature module, then objects within the feature module as the store
* then you use your reducers to setup the middle man that will interact with the store
* then set your actions with data to update the store 
* add the reducer to the store in ngModule
```ts
    StoreModule.forFeature('products', productReducer),
```

![1661622364262](image/README/1661622364262.png)

# Devtools and debugging
* get redux dev tools
* then use ng add
* maxAge stores up to x actions
* can import and export for debugging, and testing


# Strong Typing the State
* organize state into feature module
* you will want to work with state only as lazy loaded
![1661614963882](image/README/1661614963882.png)
* inital values are set in the reducer
* declare inital values as const

* using selectors decouples store from the components
* selector values are cached
![1661615534140](image/README/1661615534140.png)
* selector is a pure fn
* you can compose selectors
![1661615977077](image/README/1661615977077.png)

# Stronly Typing Action with Action Creators
* have a naming convention for your actions
* you can use actions for http requests if they succed or fail

# Working with Effects
* state may depend on external data
* an effect is a extension of a servie
![1661621621157](image/README/1661621621157.png)
* to install 
```ts
ng add @ngrx/effects
```
* ofType comes from ngrx/effects
![1661622153668](image/README/1661622153668.png)
![1661622220820](image/README/1661622220820.png)

* this now allow for async instead of subscription

# Performing Update Operations
* use immuateble array methods, so that the state is not changed

# Architectural Considerations

## Container Presentational Component Pattern
* ngrx takes logic out of components
![1661639675157](image/README/1661639675157.png)
* onChange means skipped Input that have not changed
![1661639814831](image/README/1661639814831.png)
* focus on getting rid of async pipe, and 

* use index.ts for more cleaner code

# Actions should capture events and not commands
* index.ts reduce amnt of exports

## State Module
* just state and no presentation componets

# Final Words
* @ngrx/entity hepls with CRUD operations
* @ngrx/schematics - to generate ngrx schematic
@* @ngrx/router-store - connection router to store
* @ngrx/data - abstracts away  @ngrx/entity, provide metadata and offer extension points as needed giving up comtrol

![1661650056496](image/README/1661650056496.png)