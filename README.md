# GLinkChecker

Simple app for periodic connection checking.

## Description

GLinkChecker has been written as a tool for WiFi breaks detection in home network. It's designed to periodicaly open TCP connection to router (gateway) administration interface on port 80. Check interval is user configurable (10 seconds by default). If connection cannot be established application logs error and changes check interval to 2s. After connection is back check interval is also restored to 'normal'.

## Build and usage info

Download and build application:

```shell
git clone https://github.com/markamdev/glinkchecker
cd glinkchecker
go build .
```

Run application. By default it tries to periodically connect gateway IP (in most cases it's a home router) on port 80.

```shell
./glinkchecker
```

If link checking should be performed using other remote address (or port) it can be specified using `--address` and `--port` flags:

```shell
./glinkchecker --address --port
```

Checking interval can be changed using `--interval` flag. Default interval is 30s, lowest allowed is 10s. Checking interval is set to 5s if connection break occured.

## License

Code is published under [MIT License](https://opensource.org/licenses/MIT) as it seems to be the most permissive license. If for some reason you need to have this code published with other license (ex. to reuse the code in your project) please contact [author](#author-/-contact) directly.

## Author / contact

If you need to contact me feel free to send an email to:  [markamdev.84#dontwantSPAM#gmail.com](maitlo:)
