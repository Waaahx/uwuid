# uwuid
uwuid is a small Go package that generates cute unique identifiers composed of random kawaii strings. Perfect for adding a playful touch to your IDs!

Installation
Use go get to add the package to your project:
```bash
go get github.com/Waaahx/uwuid
```

How to use : 
```go
package main

import (
    "fmt"
    "github.com/Waaahx/uwuid"
)

func main() {
    id := uwuid.GenerateUUID()
    fmt.Println("Generated uwuid:", id)
}
```

License
This project is licensed under the MIT License
