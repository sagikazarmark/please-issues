_out = f"pkg/{CONFIG.GOOS}_{CONFIG.GOARCH}/{CONFIG.GO_IMPORT_PATH}/lib/testing.a" if CONFIG.GOOD == "True" else None

go_library(
    name = "testing",
    srcs = ["testing.go"],
    out = _out,
    visibility = ["PUBLIC"],
)
