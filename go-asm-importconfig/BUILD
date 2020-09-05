http_archive(
    name = "pleasegomod",
    urls = ["https://github.com/sagikazarmark/please-go-modules/releases/download/v0.0.3/gogetgen_darwin_amd64.tar.gz"],
)

go_binary(
    name = "main",
    srcs = glob(["*.go"], exclude = ["*_test.go"]),
    visibility = ["PUBLIC"],
    deps = ["//third_party/go/golang.org/x/sys/unix"],
)
