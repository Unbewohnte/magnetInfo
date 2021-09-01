# magnet-info
## Parse a magnet link into a struct

# Installation 
```
go get github.com/Unbewohnte/magnet-info
```

# Usage
```
package main

import (
	"fmt"
	magnetinfo "magnet-info"
)

func main() {
	info, _ := magnetinfo.Parse("magnet:?xt=urn:btih:83918ea4bb488cefd3d8b8b8762597d32aebb4fa&tr=http%3A%2F%2Facademictorrents.com%2Fannounce.php&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969")

	fmt.Printf("%+v", info)

    // ...
}
```
