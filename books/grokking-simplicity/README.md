## Chapter 10 - First-cass functions part 1

Code smell: *Implicit argument in function name*

	1. There are very similar function implementations
	1. The names of the functions indicate the difference in the implementation

Refactoring: *Express implicit argument*

1. Identify the implicit argument in the function
2. Add explicit argument
3. Use the new argument in body in place of hard-coded value
4. Update the calling code

Refactoring: *Replace body with callback*

1. Identify the before, body, and after sections
2. Extract the whole thing into a function
3. Extract the body section into a function passed as an argument to that function



## Chapter 11 - First-class functions part 2

### Refactoring copy-on-write for arrays

1. Identify before, body, and after

   ```javascript
   function arraySet(array, idx, value) {
     var copy = array.slice(); // before
     copy[idx] = value; // <<-- body
     return copy; // after
   }
   
   function push(array, elem) {
     var copy = array.slice(); // before
     copy.push(elem) // <<-- body
     return copy; // after
   }
   
   function drop_last(array, elem) {
     var copy = array.slice(); // before
     copy.pop() // <<-- body
     return copy; // after
   }
   
   function drop_first(array, elem) {
     var copy = array.slice(); // before
     copy.shift() // <<-- body
     return copy; // after
   }
   ```

2. Extract function

   Before

   ```javascript
   function arraySet(array, idx, value) {
     var copy = array.slice(); // before
     copy[idx] = value; // <<-- body
     return copy; // after
   }
   ```

   After

   ```javascript
   function arraySet(array, idx, value) {
     return withArrayCopy(array);
   }
   
   function withArrayCopy(array) {
     var copy = array.slice(); 
     copy[idx] = value; // <<-- body
     return copy; 
   }
   ```

3. Extract callback

   Before

   ```javascript
   function arraySet(array, idx, value) {
     return withArrayCopy(array);
   }
   
   function withArrayCopy(array) {
     var copy = array.slice(); 
     copy[idx] = value; // <<-- body
     return copy; 
   }
   ```

   After

   ```javascript
   function arraySet(array, idx, value) {
     return withArrayCopy(
       array,
       (copy) => copy[idx] = value
     );
   }
   
   function withArrayCopy(array, modify) {
     var copy = array.slice(); 
     modify(copy);
     return copy; 
   }
   ```



Benefits:

1. Standardized discipline (copy-on-write)
2. Applied discipline to new operations
3. Optimized sequences of modifications

```javascript
var a1 = drop_first(array)
var a2 = push(a1,10)
var a3 = push(a2,11)
var a4 = arraySet(a3, 0, 42)
```

```javascript
var a4 = withArrayCopy(array, (copy) => {
  copy.shift();
  copy.push(10);
  copy.push(11);
  copy[0] = 42;
})
```

### Exercises

#### Implement `arraySet`, `push`, `drop_last`, and `drop_first` using `withArrayCopy`

```Javascript
function arraySet(array, idx, value) {
  return withArrayCopy(array, (copy) => copy[idx] = value)
}

function push(array, elem) {
  return withArrayCopy(array, (copy) => copy.push(elem))
}

function drop_last(array, elem) {
  return withArrayCopy(array, (copy) => copy.pop())
}

function drop_first(array, elem) {
  return withArrayCopy(array, (copy) => copy.shift())
}
```

#### Implement copy-on-write discipline for objects

```javascript
function objectSet(object, key, value) {
  var copy = Object.assign({}, object); // before
  copy[key] = value; // <<< -- body
  return copy // after
}

function objectDelete(object, key) {
  var copy = Object.assign({}, object); // before
	delete object[key]; // <<< -- body
  return copy // after
}
```

```javascript
function objectSet(object, key, value) {
  return withObjectCopy(object, (copy) => copy[key] = value);
}

function objectDelete(object, key) {
  return withObjectCopy(object, (copy) => delete copy[key]);
}

function withObjectCopy(object, modify) {
  var copy = Object.assign({}, object);
  modify(object);
  return copy;
}
```

##### `withLogging`

```javascript
try {
  sendEmail();
} catch( e ) {
  logToSnapErrors(e)
}

tryCatch(sendEmail, logToSnapErrors);

function tryCatch(tryFunc, catchFunc) {
  try {
    return tryFunc();
  } catch (e) {
    return catchFunc(e);
  }
}
```

##### Refactor simple if statment- *replace body with callback* 

```javascript
if ( array.length === 0 ) {
  console.log("Array is empty");
}

function when( condition, consequence ) {
  if ( condition ) consequence();
}

when( array.length === 0, () => console.log("Array is empty") );
when( hasItem(cart, "shoes"), () => setPriceByName(cart, "shoes", 0) );
```

##### Refactor simple if/else statment- *replace body with callback*

```javascript
function IF( condition, whenTrue, whenFalse ) {
  condition ? whenTrue() : whenFalse();
}
```

#### Returning functions from functions

```javascript
try {
  saveUserData(data);
} catch (e) {
  logToSnapErrors(e);
}

try {
  fetchProduct(productId);
} catch (e) {
  logToSnapErrors(e);
}

// rename to make refactoring clearer

try {
  saveUserDataNoLogging(data);
} catch (e) {
  logToSnapErrors(e);
}

try {
  fetchProductNoLogging(productId);
} catch (e) {
  logToSnapErrors(e);
}

// wrap with logging
function saveUserDataWithLogging(data) {
  try {
    saveUserDataNoLogging(data);
  } catch (e) {
    logToSnapErrors(e);
  }  
}

function fetchProductWithLogging(productId) {
  try {
    fetchProductNoLogging(productId);
  } catch (e) {
    logToSnapErrors(e);
  }
}

// generalize
function(arg) {
  try { // before
    saveUserDataNoLogging(data); // body
  } catch (e) { // after
    logToSnapErrors(e);
  }
}

function(arg) {
  try {
    fetchDataNoLogging(data);
  } catch (e) {
    logToSnapErrors(e);
  }
}

// apply "replace body with callback", but instead of adding a function as
// a callback, we'll wrap the call and return it.
function wrapWithLogging(f) {
  return function (arg) {
    try {
      f(arg);
    } catch (e) {
      logToSnapErrors(e);
    }
  }
}

var saveUserDataWithLogging = wrapWithLogging(saveUserDataNoLogging);
```

##### Exercises

###### `catchAndIgnore`

```javascript
try {
  codeThatMightThrow();
} catch(e) {
  // ignore
}

function catchAndIgnore(f) {
  return function(arg0, arg1, arg2) {
    try {
      f(arg0, arg1, arg2);
    } catch (e) {
      // ignore
      return null;
    }
  }
}
```

###### `makeAdder`

```Javascript
var increment = makeAdder(1);
// increment(10) => 11
var plus10 = makeAdder(10);
// plus10(12) => 22

function makeAdder(arg) {
  return (value) => arg + value; 
}
```

