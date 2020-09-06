go_library(
    name = "pkg",
    srcs = ["pkg.go"],
)

go_test(
    name = "test",
    srcs = ["pkg_test.go"],
    deps = [":pkg"]
)

go_library(
    name = "pkg_test_external",
    srcs = ["pkg.go", "pkg_test.go"],
)

go_test(
    name = "test_external",
    srcs = ["pkg_external_test.go", "pkg_test.go"],
    external = True,
    deps = [":pkg"],
)

go_test(
    name = "test_external2",
    srcs = ["pkg_external_test.go"],
    external = True,
    deps = [":pkg_test_external"],
)
