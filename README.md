# git-sync
This is git-based simple file-sync tool for multiple directories.  
without `git-sync`, you have to visit all dirs by using `cd` and `git` commands.
with `git-sync`, all you have to do is typing `git-sync`.


## Install
```bash
$ go install github.com/koooyooo/git-sync@latest
```

## Config
```bash
$ mkdir ~/.git-sync
$ vim ~/.git-sync/config.yaml
```

```yaml: config.yaml
dirs:
  - name: tech-docs
    path: ~/projects/tech-docs
  - name: daily-memos
    path: ~/work/daily-memos
```

## Usage

synchronize with auto generated messages with `update at: YYYY-MM-DD HH:MI:SS` style.
```bash
$ git-sync
```

synchronize with custom messages with editor like `git commit` without `-m` option.  
`-c` stands for **Custom** message for updated repositories.
```bash
$ git-sync -c
```