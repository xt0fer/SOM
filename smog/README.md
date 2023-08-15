# SMOG (2)

SOM and _A Little Smalltalk_ (which has always fascinated me, since about 1983) are small, OOP languages, in a pretty _pure_ sense.

**Smog** is small Go implementation of SOM.

I'm going to try to advance thru the creation using a each step as a "build it from scratch" project. (we will see if I can actually do that).

This _smog_ is a 1.5v attempt. 1.0 was to tranliterate Java/C to golang. That was the first _throw-away_. This might be the first prototype or maybe just the second throw-away.


### Step 0

- [x] _What's the go version of hierachical classes?_ Well, because of the _composition_ way of doing things, one can nest _structs_. (but not interfaces.)
- [x] This is the first golang porject I'm thinking thru that is being done in the <smirk> _generics era_ (like baseball's _deadball/liveball_ era divide).

### Step 1

- [x] Build the Universe structs and methods.
- [x] Create enough Data Model (ObjectSystem) of the running SOMObject(s) to support `initializeObjectSystem() *Object`.
- [x] At this phase, not all the object/class/etc methods for each struct are built out.

### Step 2

- [ ] Add tests to Step 1 work
- [ ] Build a first draft of the bytecode interpreter.
  - [x] build out `Frame`
  - [ ] build out `Interpreter`
- [ ] Add the ability to load primitives into the ObjectSystem.
- [ ] Build out `Class` methods
- [ ] Build out `Method` methods

...

### Step x

- Start work on `Compiler`
- Draft Lexer, Parser...


### LSP support for SOM

There is an LSP support for _vscode_ [effortless-language-servers](https://marketplace.visualstudio.com/items?itemName=MetaConcProject.effortless-language-servers)

### Why did you organize the SOM code differently?

See #4 in[Structuring Applications in Go](https://www.gobeyond.dev/structuring-applications/).
We all need to be golang coders like Ben Johnson.
