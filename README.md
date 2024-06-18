# System Tests

A black/grey box suite that tests the functionality of the 0Chain network as an end user via the CLI tools.

# Zs3server Testing using Warp

## Run individual tests against local 0chain network

Prerequisites: [Go](https://go.dev/doc/install) 

For developing new system tests for code still in developer branches, tests can be run against a locally running chain. Typically, for a 0chain change you will have a PR for several modules that need to work together. For example, 0chain, blobber, GoSDK, zboxcli,  zwalletcli, minio and zs3server.



For zboxcli, zwalletcli, mc , and zs3server you need to first build the executable and copy into local system test directory. For Example:

```
cd zboxcli
make install
cp ./zbox ../zs3server-tests/tests/cli_tests


cd ../zwalletcli
make zwallet
cp ./zwallet ../zs3server-tests/tests/cli_tests


zs3server
cd ../zs3server
go mod tidy
go build .
cp ./minio ../zs3server-tests/tests/cli_tests
```


## Running the test cases


Make sure you have the correct system test branch. Now you need to edit system_test/tests/cli_tests/config/zbox_config.yaml Edit the line block_worker: https://dev.0chain.net/dns to the appropriate setting for you, something like

```block_worker: http://192.168.1.100:9091```

Now open the system_test project in GoLand, you should now be able to run any of the cli_tests in debug.

You can run tests against a remote chain if you have already deployed elsewhere eg. dev.0chain.net

## Handling test failures
The test suite/pipeline should pass when ran against a healthy network.
If some tests fail, it is likely that a code issue has been introduced.
Try running the same tests against another network to rule out environmental issues.
If the failure persists, and you believe this to be a false positive, [contact the system tests team](https://0chain.slack.com/archives/C02AV6MKT36).

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.


## License
[MIT](https://choosealicense.com/licenses/mit/)


running the test 


Prerequistes
File can be changed to run custom test cases 
1. allocation.yaml 
    -- change data, parity and lock while creating an allocation

2. mc_hosts.yaml
    -- change server, port, access_key, secret_key, concurrenct, secondary_serer, secondary_port

2. hosts.yaml
    -- change server, port, access_key, secret_key, concurrenct
```
cd tests/cli_tests
go test


<!-- for zs3server test -->

cd zs3server_tests
go test

```