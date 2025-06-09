Primeiro iniciamos um modulo em go com:

```go
go mod init ...
```

Em seguida vamos instalar o chi com o comando:

```go

go get github.com/go-chi/chi/v5
```

O primeiro código que vamos fazer, quando executado, não gera nada no console.

```go
package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	router := chi.NewRouter()

	router.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)

		w.Write([]byte(`{message:"Hello World"}`))
	})

  

	http.ListenAndServe(":3000", router)
}

```