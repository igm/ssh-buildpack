# ssh-buildpack
CF buildpack that enables to connect into container (warden) via SSH.

```bash
$ mkdir empty
$ cd empty
$ cf push APP_NAME -b github.com/igm/ssh-buildpack
$ go get github.com/igm/ssh-buildpack/src/client
$ ./client APP_URL
$ ssh foo@localhost -p 2222  # pass is: bar
```
