# Missing import config from ASM library

When a Go library contains Assembly code, but also depends on other libraries, its dependencies are missing from the import config.

You can reproduce it with:

```bash
./pleasew build :main
```

It fails with:

```
Error building target //:main: exit status 1
cannot find package golang.org/x/sys/internal/unsafeheader (using -importcfg)
/usr/local/go/pkg/tool/darwin_amd64/link: cannot open file : open : no such file or directory
```
