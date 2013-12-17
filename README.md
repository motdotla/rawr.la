# Rawr.la

rawr. hello there LA.

## Development

```
git clone https://github.com/scottmotte/rawr.la.git
cd rawr.la
go get github.com/codegangsta/martini
go get github.com/codegangsta/martini-contrib
go run server.go
```

## Production

```
git clone https://github.com/scottmotte/rawr.la.git
cd rawr.la
heroku create -b https://github.com/kr/heroku-buildpack-go.git 
git push heroku master
```
