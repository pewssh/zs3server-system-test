# System Tests

A black/grey box suite that tests the functionality of the 0Chain network as an end user via the CLI tools.

## Running tests

The tests require a full 0Chain network to be deployed and running in a healthy state.

Now open the system_test project in [GoLand](https://www.jetbrains.com/go/),
you should now be able to run any of the `cli_tests` in debug.

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
