## Lorepixum Web Service

This tiny web service returns images` stubs with the given sizes. In fact it provides http interface for [lorempixum](github.com/vitkarpov/lorempixum) package.

Try it:

```
$ curl lorempixum.ru/jpeg/100/100 > foo.jpeg
```

## HTTP API

```
localhost/{format}/{x}/{y}
```

### format

Could be jpeg, png or gif for animated images

### x and y

Width and height of the image. Should be positive integer, otherwise web server returns a bad request error.
