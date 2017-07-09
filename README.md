# mini
Skeleton application, based [Gorilla] Go web toolkit.

I'm just try to organize the routes and services (inspired by service container from [Vertebrae])

### How to run
```
$ glide instal
$ go build main.go
$ ./main
```

### Tips
Middleware, how we deal with Middleware by specific routes using [Negroni], here the [link-article-middleware]

[Gorilla]: <http://www.gorillatoolkit.org/>
[Vertebrae]: <https://github.com/EwanValentine/Vertebrae>
[Negroni]: <https://github.com/urfave/negroni>
[link-article-middleware]: <https://www.calhoun.io/route-specific-middleware/>
