---
title: What is Infra?
position: 1
---

# What Is Infra?

## Overview

Infra is a tool that manages access to Kubernetes clusters. Access is granted or revoked via an API or CLI, and in the background, Infra takes care of provisioning users & groups with the right permissions no matter where the cluster is hosted (AWS, Google Cloud, etc). Users are automatically distributed short-lived credentials that expire after a short period of time. For larger teams, Infra integrates with identity providers like Okta to automatically give access via existing accounts.

## Example (Kubernetes)

```bash
# Log in as henry@acme.com via Okta
$ infra login infra.acme.dev
  ... logging in with Okta
  ... logged in as henry@acme.com

# Discover what you can access via `infra list`
# In the example below, 3 Kubernetes clusters are connected to Infra.
# Infra has already synchronized the kubeconfig file so the
# user can use their tool of choice right away (i.e. kubectl)

$ infra list
  NAME                           ACCESS
  production                     view
  production.web                 edit
  staging                        edit
  development                    edit

# You can switch clusters using your existing tools
# Infra also includes a CLI command for switching clusters
$ infra use production.web
Switched to context "production.web".

$ kubectl get pods
NAME                  READY   STATUS    RESTARTS   AGE
web-d6797786d-cz5jh   1/1     Running   0          8h
web-d6797786d-wzxns   1/1     Running   0          8h
web-d6797786d-zjbvl   1/1     Running   0          8h
```


## Use Cases

### Cloud-agnostic IAM

Infra works anywhere and doesn't depend on any existing identity & access management system such as AWS IAM, Google IAM or Azure AD. It can be completely self-hosted, including behind existing VPNs or proxies.

### Automatic onboarding & offboarding


### Just-in-time access

### Access requests


## How it works

### Short-lived tokens

### Granting Access

### Connectors

### The API server

## What's next

Get up and running with the [Quickstart](./quickstart.md) guide or read about the [how Infra works](./how-infra-works.md).
