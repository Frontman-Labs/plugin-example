# Example Plugins

This repository contains example plugins for use with Frontman, an API Gateway written in Go.

## Building a Plugin

To build the plugins, run the following command:

```bash
make all
```

This will build multiple plugin shared objects that can be loaded by Frontman.

## Using a Plugin

To use a plugin, add its file path to the Frontman configuration file, under the plugin.order field. For example:

```yaml
plugin:
  enabled: true
  order:
    - "/path/to/plugin.so"
```

This will cause Frontman to load the plugin and call its methods during request processing.

## Contributing
Contributions to this repository are welcome. If you have an example plugin you'd like to add, please submit a pull request with the plugin code and a README file explaining its use.