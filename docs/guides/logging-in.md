---
title: Logging in
position: 1
---

# Logging In

## Install Infra CLI

{% tabs %}
{% tab label="macOS" %}
```
brew install infrahq/tap/infra
```
You may need to perform `brew link` if your symlinks are not working.
```
brew link infrahq/tap/infra
```
{% /tab %}
{% tab label="Windows" %}
```powershell
scoop bucket add infrahq https://github.com/infrahq/scoop.git
scoop install infra
```
{% /tab %}
{% tab label="Linux" %}
```
# Ubuntu & Debian
echo 'deb [trusted=yes] https://apt.fury.io/infrahq/ /' | sudo tee /etc/apt/sources.list.d/infrahq.list
sudo apt update
sudo apt install infra
```
```
# Fedora & Red Hat Enterprise Linux
sudo dnf config-manager --add-repo https://yum.fury.io/infrahq/
sudo dnf install infra
```
{% /tab %}
{% /tabs %}

## Login to Infra

```
infra login SERVER
```

> **Don't know your Infra server?**
> Ask your Infra team for the hostname that you should use to login.

## See what you can access

```
infra list
```

## Switch to the cluster context you want

```
infra use DESTINATION
```

## Use your preferred tool to run commands

```
# for example, you can run kubectl commands directly after the infra context is set
kubectl [command]
```
