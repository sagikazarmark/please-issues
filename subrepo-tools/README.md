# Subrepo tools

I'm a huge fan of reusing tools across projects and making project tooling as standard as possible.

This is why I created my version of the pleasings repo: https://github.com/sagikazarmark/mypleasings

The repo comes with two kinds of things:

- Build rules for things like Helm charts
- A bunch of rules for downloading binaries

In case of Helm, there is a `helm_lint` rule that lints the Helm chart.
For that, it needs a Helm binary, which is downloaded in the same repo, using the `helm_binary` rule.

Relevant parts:

```starlark
CONFIG.setdefault("HELM_TOOL", "//tools/k8s:helm")

def helm_lint(name:str, visibility:list=None):
    """Lints a Helm chart.

    Args:
      name (str): Name of the rule.
      visibility (list): Visibility of the rule.
    """

    return gentest(
        name = name,
        test_cmd = f'"$(location {CONFIG.HELM_TOOL})" lint "$PKG_DIR" 2>&1',
        data = glob(["**"], hidden = True) + [CONFIG.HELM_TOOL],
        visibility = visibility,
        no_test_output = True,
    )
```

And the download target in `tools/k8s`:

```starlark
CONFIG.setdefault("HELM_VERSION", "3.3.1")

helm_binary(
    name = "helm",
    version = CONFIG.HELM_VERSION,
    visibility = ["PUBLIC"],
)
```

When running tests inside my pleasings clone, it works perfectly.

When I try to use it as a subrepo (as shown in this example), it complains:

```
❯ plz build
Build stopped after 360ms. 1 target failed:
    //tools/k8s:helm
//deploy/charts/hello-world:lint depends on //tools/k8s:helm, but the directory tools/k8s doesn't exist
```

I need to manually configure `HELM_TOOL` in my repo in order to make it work.

My aim is to create a rule, that can go with the "default" rule in the subrepo, but also gives the user a chance to override it.
I think this is somewher in line with your goals, to move tools out of the please distribution.

Is there a way I can do that?
