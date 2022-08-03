go build cmd/app.go
vi halcommand.sh

```
echo 'started the shell command'
curl https://google.com
echo 'finished the shell command'
```
chmod +x halcommand.sh

copy to /usr/local/bin/halcommand.sh (or move somewhere PATH variable recognizes , do `echo $PATH` and see)

./app