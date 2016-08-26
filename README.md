# go-libreofficekit [![](https://godoc.org/github.com/docsbox/go-libreofficekit?status.svg)](https://godoc.org/github.com/docsbox/go-libreofficekit)
Cgo bindings to LibreOfficeKit

# Install
```bash 
# Latest version of LibreOffice (5.2) is required
$ add-apt-repository ppa:libreoffice/ppa 
$ sudo apt-get update
$ apt-get install libreoffice libreofficekit-dev
$ go get github.com/docsbox/go-libreofficekit
```

# Usage

```go
package main

import "github.com/docsbox/go-libreofficekit"

func main() {
    office, _ := libreofficekit.NewOffice("/path/to/libreoffice")
    
    document, _ := office.LoadDocument("kittens.docx")
    document.SaveAs("kittens.pdf", "pdf", "skipImages")

    document.Close()
    office.Close()
}

```
