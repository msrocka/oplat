# oplat

`oplat` is a command line tool to check for updated components in the openLCA target platform.

## Usage

Update the `repos.txt` file with the URLs of the repositories you want to check for updates.
Then, just run the `oplat` command:

```bash
oplat path/to/olca-app/olca-app/platform.target
```

It will then check the repositories and print the components that could be updated.
