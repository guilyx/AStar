---
name: Bug report
about: Create a report to help us improve
title: ''
labels: bug
assignees: ''

---

Thank you for creating the issue!

- [ ] Yes, I'm using a binary release within 2 latest major releases. Only such installations are supported.
- [ ] Yes, I've searched similar issues on GitHub and didn't find any.
- [ ] Yes, I've included all information below (version, config, etc).

Please include the following information:

<details><summary>Version of golangci-lint</summary>

```console
$ golangci-lint --version
# paste output here
```

</details>

<details><summary>Config file</summary>

```console
$ cat .golangci.yml
# paste output here
```

</details>

<details><summary>Go environment</summary>

```console
$ go version && go env
# paste output here
```

</details>

<details><summary>Verbose output of running</summary>

```console
$ golangci-lint cache clean
$ golangci-lint run -v
# paste output here
```

</details>
