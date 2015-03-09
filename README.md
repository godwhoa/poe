A little lib for phonetic key generation<br>
or you could say port of hastebin's key generator


P.S:Ignore the bad package name :p<br>
example:
```
package main

import (
	"fmt"
	"github.com/godwhoa/poe"
)
func main(){
	//Generate a key of length 10
	fmt.Println(poe.KeyGen(10))	
}

```