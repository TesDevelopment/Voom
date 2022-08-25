# Voom
A small "stack" language (similar to luac).

## "Quick" Docs

### Stacks
Stacks are the storage systems in voom. They are premade and can not be created/deleted.
These stacks store different data types.

#### String stack
The string stack stores all strings pushed by the user.

#### Int stack
The integer stack stores all ints passed by the user.

#### Global stack
The global stack stores functions the user fetches from storage.

### Globals
Globals are pre-made functions that can be fetched and pushed to the global stack.

#### printS
printS outputs the string located at the passed location on the string stack.

#### printI
printInoutputs the string located at the passed location on the int stack.

### Opcodes
Opcodes (or tokens) are like the building blocks of all "voom" code.
The best description of them would be **instructions** that the language understands.
Opcodes take 0->2 arguments. The first argument will usually index one of the stacks (See above)
the second arguemnt (if any) varies.

#### PUSHSTRING
pushS pushes a string onto the string stack.

```
pushS foo
```

#### PUSHINT
pushI pushes an int onto the int stack.

```
pushI 10
```

#### getglobal
getglobal retrieves a global function (see above) and is currently the only way to aquire a
function.

```
getglobal prints
```

#### call
call allows you to call/use any functions on the function stack.

```
call 0 5
     ^
     Function stack index
       ^
       String/Int stack index
```

#### deferS
deferS clears the String stack

#### deferI
deferI clears the Int stack
