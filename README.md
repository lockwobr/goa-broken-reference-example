### Goa example of broken Reference

# repro
Gerenate the swagger with `make swagger`. I have pretty printed it the swagger.json and you can see that ther FirstName and Lastname do not inherit from the media type. But if you make a fresh project and use the cli and use the same design it works.

If you do the flowing you get different swagger output.

```
    cd $GOPATH/src/github.com/lockwobr/
    mkdir goa-broken-reference-example-test
    cp -R goa-broken-reference-example/design/ goa-broken-reference-example-test/
    goagen swagger -d github.com/lockwobr/goa-broken-reference-example-test/design -o public

```