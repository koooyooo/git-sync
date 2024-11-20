# git-sync
`git sync` is auto sync tool for one or many repositories.

## config
```bash
$ mkdir ~/.git-sync
$ touch ~/.git-sync/config.yaml
```

```yaml: config.yaml
dirs:
  - name: docs
    path: ~/docs
```

```bash
$ git-sync
```