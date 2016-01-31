## Lorepixum Web Service

This tiny web service returns images` stubs with the given sizes. In fact it provides http interface for [lorempixum](https://github.com/vitkarpov/lorempixum) package.

Try it online:

```
$ curl https://pure-river-38402.herokuapp.com/jpeg/100/100 > foo.jpeg
```

I use it to get images into static html markups during development

## HTTP API

```
localhost/{format}/{x}/{y}
```

### format

Could be jpeg, png or gif for animated images

### x and y

Width and height of the image. Should be positive integer, otherwise web server returns a bad request error.
