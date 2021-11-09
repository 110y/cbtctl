# cbtctl

A CLI tool for GCP Cloud Build Triggers

## Usage

### Run a build trigger with substitutions

```
$ cbtctl run --project <GCP Project Name> --trigger <Trigger Name> --substitutions "_VAR=val"
```

### On Cloud Build

By making `cbtctl` container connect to the [network provided by Cloud Build](https://cloud.google.com/build/docs/build-config-file-schema#network), we can use `cbtctl` even on the Cloud Build.

#### Example `cloudbuild.yaml`:

```yaml
steps:
    - name: gcr.io/cloud-builders/docker:19.03.9
      args:
          - run
          - --rm
          - --network
          - cloudbuild
          - 110y/cbtctl:0.0.1
          - run
          - --project=<GCP Project Name>
          - --trigger=<Trigger Name>
          - --substitutions="VAR=val"
```

## Motivation

While `gcloud` have a `builds triggers` subcommand which controls the build triggers, not all build triggers features are supported by `gcloud`. For example, we can run a build trigger by executing `gcloud alpha builds triggers run`, however, we can not use the `substitutions` feature for the build trigger with such `gcloud` command. That's why I made `cbtctl`. `cbtctl` aims to support build triggers features which are not supported by `gcloud` at this time.
