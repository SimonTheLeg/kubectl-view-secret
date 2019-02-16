# View Secret

Simple golang tool to view the base64 decoded content of a kubernetes secret

## Installation

Grab a version from the release page and place it inside your $PATH.

## Usage

You can call the plugin directly using kubectl

```
kubectl view-secret secret-name [-n, --namepspace secret-namespace]
```

Alternatively you can also use this plugin as a standalone binary

```shell
kubectl-view-secret secret-name [-n, --namepspace secret-namespace]
```

If -n, --namespace is not set, it defaults to the current namespace.