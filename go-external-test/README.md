# Mixed external and "internal" tests

When an external test refers to something in the "internal" (in this case: it's in the the package, not in `_test`) package in a `_test.go` file,
Please is unable to build the external test package. (See the example in this repo)

For example:

```starlark
go_test(
    name = "test_external",
    srcs = ["pkg_external_test.go", "pkg_test.go"],
    external = True,
    deps = [":pkg"],
)
```

The above rule sources `pkg_test.go` which is in the `pkg` package, whereas `pkg_external_test.go` is in `pkg_test`.

The build fails with the following error:

```
❯ plz build --shell //pkg:test_external
Build stopped after 530ms. 1 target failed:
    //pkg:_test_external#lib
Error building target //pkg:_test_external#lib: exit status 1
pkg/pkg_external_test.go, line 12, column 14: undefined: pkg.CompareWith
pkg/pkg_test.go, line 1, column 9: package pkg; expected pkg_test
```

It complains about the package not being external.

As a workaround, I tried building a separate library from the library source and the necessary `_test.go` files:

```starlark
go_library(
    name = "pkg_test_external",
    srcs = ["pkg.go", "pkg_test.go"],
)

go_test(
    name = "test_external2",
    srcs = ["pkg_external_test.go"],
    external = True,
    deps = [":pkg_test_external"],
)
```

It fails with a different error, complaining about a missing import path:

```
❯ plz build --shell //pkg:test_external2
Build stopped after 560ms. 1 target failed:
    //pkg:_test_external2#lib
Error building target //pkg:_test_external2#lib: exit status 1
pkg/pkg_external_test.go, line 6, column 2: can't find import: "github.com/sagikazarmark/please-issues/go-external-test/pkg"
```

I traced the error back to the way import config is built and commented about it here: https://github.com/thought-machine/please/pull/1177#issuecomment-687793475


**Possible solutions**

The above workaround would be a great start for solving these issues.
Actually, in general, I think it would be great if `go_library` allowed specifying every aspect of package naming.

Another solution could be adding an `internal_srcs` attribute to `go_test` and accept test files from the same package there.
