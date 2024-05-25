<h1 align="center">uthoctl</h1>


uthoctl is the command line client to access the [utho-api](https://utho.com/api-docs) by utlizing the utho-go SDK [utho-sdk](https://github.com/uthoplatforms/utho-go).

- [Installing `uthoctl`](#installing-uthoctl)
  - [Downloading a Release from GitHub](#downloading-a-release-from-github)
  
- [Examples](#examples)


## Installing `uthoctl`

### Downloading a Release from GitHub

Visit the [Releases
page](https://github.com/uthoplatforms/utho-cli/releases) for the
[`uthoctl` GitHub project](https://github.com/uthoplatforms/utho-cli), and find the
appropriate archive for your operating system and architecture.
Download the archive from your browser or copy its URL and
retrieve it to your home directory with `wget` or `curl`.

### CLI Version
where `<version>` is the full semantic version, e.g., `0.1.5`.

You can get latest release from [Latest Releases
page](https://github.com/uthoplatforms/utho-cli/releases/latest)

### Installation on Linux

```bash
curl -LO https://github.com/uthoplatforms/utho-cli/releases/download/v<version>/uthoctl_<version>_linux_amd64.tar.gz
tar xf uthoctl_<version>_linux_amd64.tar.gz
sudo mv uthoctl /usr/local/bin
```

### Installation on MacOS

For x86 based Macs (intel cpu):

```bash
curl -LO https://github.com/uthoplatforms/utho-cli/releases/download/v<version>/uthoctl_<version>_darwin_amd64.tar.gz
tar xf uthoctl_<version>_darwin_amd64.tar.gz
sudo mv uthoctl /usr/local/bin
```

For Apple Silicon (M1) based Macs:

```bash
curl -LO https://github.com/uthoplatforms/utho-cli/releases/download/v<version>/uthoctl_<version>_darwin_arm64.tar.gz
tar xf uthoctl_<version>_darwin_arm64.tar.gz
sudo mv uthoctl /usr/local/bin
```

### Installation on Windows

```bash
curl -LO https://github.com/uthoplatforms/utho-cli/releases/download/v<version>/uthoctl_<version>_windows_amd64.tar.gz
```
You should be able to double-click the zip archive to extract the `uthoctl` executable.

Windows users can follow [How to: Add Tool Locations to the PATH Environment Variable](https://msdn.microsoft.com/en-us/library/office/ee537574(v=office.14).aspx) in order to add `uthoctl` to their `PATH`.


## Authenticating with Utho

To use `uthoctl`, you need to authenticate with Utho by providing a Personal Access Token (PAT) is the only method of
authenticating with the API. You can manage your tokens
at the Utho Control Panel [Applications Page](https://console.utho.com/switch/api).

```
uthoctl auth
```

You will be prompted to enter the Utho access token that you generated in the Utho control panel.

## Examples

`uthoctl` is able to interact with your Utho resources. Use `uthoctl --help` to get help about the CLI command. Below are a few common usage examples.

* Create new Compute Instances on your account:
```
uthoctl instance create <instance-name> --dcslug <location-slug> --image <image-name> --planid <plane-id> --billingcycle <billing cycle>
```

* List all Compute Instances on your account:
```
uthoctl instance list
```

* Add new domain to your account:
```
uthoctl domain <domain-name>
```

* Get information about your account:
```
uthoctl account get
```

* Add new firewall to your account:
```
uthoctl firewall <firewall-name>
```

* Add new loadbalancer to your account:
```
uthoctl loadbalancer <loadbalancer-name> --dcslug <location-slug> --type <loadbalancer-type>
```
