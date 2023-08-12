# SMOG (2)

SOM and _A Little Smalltalk_ (which has always fascinated me, since about 1983) are small, OOP languages, in a pretty _pure_ sense.

**Smog** is small Go implementation of SOM.

I'm going to try to advance thru the creation using a each step as a "build it from scratch" project. (we will see if I can actually do that).

This _smog_ is a 1.5v attempt. 1.0 was to tranliterate Java/C to golang. That was the first _throw-away_. This might be the first prototype or maybe just the second throw-away.


### Step 0

_What's the go version of hierachical classes?_ Well, because of the _composition_ way of doing things, one can nest _structs_. (but not interfaces.)
This is the first golang porject I'm thinking thru that is being done in the <smirk> _generics era_ (like baseball's _deadball/liveball_ era divide).


### LSP support for SOM

There is an LSP support for _vscode_ [effortless-language-servers](https://marketplace.visualstudio.com/items?itemName=MetaConcProject.effortless-language-servers)