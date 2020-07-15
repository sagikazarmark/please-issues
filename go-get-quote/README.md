# Please Go Get Quote issue

This repo reproduces a bug introduced in [thought-machine/please#856](https://github.com/thought-machine/please/pull/856).

```bash
# Latest go_get rule
plz run :example

# An earlier go_get rule from wollemi
plz run --profile good :example
```
