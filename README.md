# ssh-buildpack
CF buildpack that enables to connect into container (warden) via SSH.

```bash
$ mkdir empty
$ cd empty
$ touch file   # at least one empty file need to be in the directory
$ cf push APP_NAME -b github.com/igm/ssh-buildpack
$ go get github.com/igm/ssh-buildpack/src/client
$ client APP_URL  # go get should produce binary
$ ssh foo@localhost -p 2222  
foo@localhost\'s password: # enter "bar" as password
vcap@xxxxxxxxx:~$ lsb_release -a    # to check the ubuntu version
No LSB modules are available.
istributor ID:	Ubuntu
Description:	Ubuntu 10.04.4 LTS
Release:	10.04
Codename:	lucid
vcap@xxxxxxxxx:~$
```

The sshd server is a binary compiled from https://github.com/jpillora/go-and-ssh/tree/master/sshd

