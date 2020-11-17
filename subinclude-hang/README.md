# Subrepo tools

Please hangs when a package is subincluded in the root package and in subpackages as well.

```
❯ plz query deps -v debug --show_all_output //cmd
09:08:54.901   DEBUG: Found repo root at /Users/mark/Projects/sagikazarmark/please-issues/subinclude-hang
09:08:54.901   DEBUG: Attempting to read config from /etc/please/plzconfig...
09:08:54.901   DEBUG: Attempting to read config from /Users/mark/.config/please/plzconfig...
09:08:54.901   DEBUG: Read config from /Users/mark/.config/please/plzconfig
09:08:54.901   DEBUG: Attempting to read config from /Users/mark/Projects/sagikazarmark/please-issues/subinclude-hang/.plzconfig...
09:08:54.901   DEBUG: Read config from /Users/mark/Projects/sagikazarmark/please-issues/subinclude-hang/.plzconfig
09:08:54.901   DEBUG: Attempting to read config from /Users/mark/Projects/sagikazarmark/please-issues/subinclude-hang/.plzconfig_darwin_amd64...
09:08:54.901   DEBUG: Attempting to read config from /Users/mark/Projects/sagikazarmark/please-issues/subinclude-hang/.plzconfig.local...
09:08:54.902   DEBUG: Loading built-in build rules...
09:08:54.920   DEBUG: Parser initialised
09:08:54.920   DEBUG: Original target scan complete
09:08:54.921   DEBUG: Registering subrepo pleasings2 in package //:all
09:09:04.903  NOTICE: Build running for 10s, 1 / 3 tasks done, 0 workers busy
09:09:14.904  NOTICE: Build running for 20s, 1 / 3 tasks done, 0 workers busy
09:09:24.902  NOTICE: Build running for 30s, 1 / 3 tasks done, 0 workers busy
09:09:34.905  NOTICE: Build running for 40s, 1 / 3 tasks done, 0 workers busy
```
