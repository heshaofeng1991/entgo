### Golang entgo 说明。
- entgo       go orm框架(mysql)  <br>

#### install
```bash
go get -u -v github.com:heshaofeng1991/entgo.git
go install github.com:heshaofeng1991/entgo.git@latest
```

#### schema有更新的时候
```bash
make all
```

#### notes
1. 当有schema变化的时候，注意看一下makefile，大家可以自行去更新就好。