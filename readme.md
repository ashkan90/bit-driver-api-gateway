## Test purpose
To test the gateway, you should run firstly;
```shell
sh run.sh
```

then start mock services to test the gateway.
```shell
cd mock/svc-1 && go run .
cd mock/svc-2 && go run .
```

_for now authentication middleware/strategy is bypassed manually. to see is it working you can uncomment lines in `proxy/middleware/check_auth.go` file_

#TODO
- [ ] Unit test
- [ ] Update readme file
