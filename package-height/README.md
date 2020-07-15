# Conflicting package names

When introducing Please to a large project, the following error ocurred:

```
    //internal/helm:_test#lib
Error building target //internal/helm:_test#lib: exit status 1
internal/helm/envresolver_integration_test.go, line 21, column 2: internal compiler error: conflicting package heights 44 and 12 for path "testing"

Please file a bug report including a short program that triggers the error.
https://golang.org/issue/new
```

Based on examining the code, the culprit was a custom `testing` package that was imported alongside the `testing` package from the standard library.

When I tried to reproduce this issue, I found that packages with the same name can potentially conflict and simply overwrite each other.
(My guess is that if the official `testing` package gets imported in a transitive dependency and then higher in the chain there is another testing package, you get the package height error above.)
In this repo, the custom `testing` package completely overwrites the stdlib one:

```
❯ plz test
Build stopped after 90ms. 1 target failed:
    //lib:_test#lib
Error building target //lib:_test#lib: exit status 1
lib/lib_test.go, line 9, column 17: undefined: "testing".T
```

The solution in this case is to follow Go's pkgdir convention: place archives under `pkg/OS_ARCH/path/to/pkg.a`.

See it in action:

```
❯ plz test --profile good
//lib:test 1 test run in 0s; 1 passed [cached]
1 test target and 1 test run in 0s; 1 passed. Total time 70ms
```

I think by default `go_library` should output archives to the official location. Currently it leaves them in the package dir and symlinks potential `.a` files during compilation.
I'm pretty sure there was a use case for that, but it causes issues with conflicting package names.
