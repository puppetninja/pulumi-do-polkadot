# Devspace

It is possible to test the chart on minikube.

First, build the custom images:

```
devspace build -t dev --skip-push
```

Create `test_values.yaml` based on the default chart values in charts/polkadot/values.yaml

Then start minikube and run:

```
helm install test -f test_values.yaml  charts/polkadot --namespace test --create-namespace
```
