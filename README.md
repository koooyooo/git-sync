# git-sync

<img src="image/image.png">

This is git-based simple file-sync tool for multiple directories.  
This is equivalent to applying following commands for each directory.

- `cd ${TARGET_DIR}`
- `git pull`
- `git add .`
- `git commit -m "updated at: ${TIME}"`
- `git push`

without `git-sync`, you have to visit all dirs by using `cd` and `git` commands.  
with `git-sync`, all you have to do is typing `git-sync`.


## Install
Just using `go install`
```bash
$ go install github.com/koooyooo/git-sync@latest
```
or `git clone` and install by `make`
```bash
$ git clone https://github.com/koooyooo/git-sync
$ cd git-sync
$ make install
```

## Config
`git-sync` requires `config.yaml` located in `~/.git-sync` directory.
```bash
$ mkdir ~/.git-sync
$ vim ~/.git-sync/config.yaml
```
sample contents is the followings.
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

synchronize with messages with editor like `git commit` without `-m` option.  
`-e` stands for **Edit** message for updated repositories.
```bash
$ git-sync -e
```