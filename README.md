# Gorky 

An app to change your desktop wallpaper from different sources.
 
Currently, implemented for Unsplash only, more may come.

NOTE: the name is a work in progress


## Installation

```bash
go get github.com/guumaster/gorky
go install github.com/guumaster/gorky
```

## Usage

```
$> gorky 

// Output: 
2020/05/02 22:40:34 Downloading new image
2020/05/02 22:40:35 Changing background
2020/05/02 22:40:35 Saving background to /home/$HOME/.local/share/gorky/gorky_285610876.png
2020/05/02 22:40:36 Cleaning old backgrounds
```

## Install as a service

The service will change the wallpaper each 12 hours (it would be a parameter at some point)
```
gorky-service -service install 

systemctl start gorky
```

## TODO

* [ ] migrate to `cobra`
* [ ] export `service/runner` to separate package
* [ ] make the CLI usable
* [ ] allow to change unsplash collection by config
* [ ] check for config before service installation

* [ ] integrate more services (Pinterest, wa)
* [ ] add service parameters. ex: 
    - runAt("12pm")
    - or more "cron like"
