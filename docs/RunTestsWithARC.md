# How I experimented using ARC to run tests on clouds

Working on Windows 11 Home.

## Solution 1

For solution 1, we just run the tests on ARC created runner pods on k8s.  

### Setup Local Working Env

- gcloud
- kubectl
- helm

Install gcloud by downloading the [Google Cloud CLI](https://docs.cloud.google.com/sdk/docs/install-sdk).  

Install kubectl using gcloud:  

```powershell
gcloud components install kubectl
```

Install helm using winget:  

```powershell
winget install Helm.Helm
```

### Set up Cloud Env

Set up clusters in cloud:  

```powershell
gcloud container clusters create alphabetbengali `
  --zone=asia-northeast1-a `
  --num-nodes=1 `
  --enable-autoscaling `
  --min-nodes=1 `
  --max-nodes=3 `
  --disk-size=30 `
  --disk-type=pd-standard
```

Connect kubectl to my cluster:  

```powershell
gcloud container clusters get-credentials alphabetbengali --zone=asia-northeast1-a --project=project-b8ee474a-9624-49ce-b08
```

Install ARC from helm:  

```powershell
helm install arc `
  --namespace arc-systems `
  --create-namespace `
  oci://ghcr.io/actions/actions-runner-controller-charts/gha-runner-scale-set-controller
```

Set the PAT from GitHub to my ARC:  
(It should have `repo` and `workflow` permissions)  

```powershell
helm install go-test `
  --namespace arc-runners `
  --set githubConfigUrl="https://github.com/CriticalHarupi/AlphabetBengali" `
  --set githubConfigSecret.github_token="MY_PAT" `
  --set runnerScaleSetName="go-test" `
  oci://ghcr.io/actions/actions-runner-controller-charts/gha-runner-scale-set
```

### Set up yml for Workflow

As we named the scale set go-test, we use this as label for `runs-on` in ymls for GitHub Actions workflows.  

Also, because we are using the ARC images to run the tests, we need to install dynamically for each run.  

```yml
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version-file: backend/go.mod
```

This worked.  
But it is not very efficient, as each run you will need to set up go, and that costs time and resources on cloud. So I am doing a solution 2.  

## Solution 2

For solution 2, Almost everything is same, just that I will try out this instead of a step to set up go:  

```yml
    runs-on: go-test
    container:
      image: golang:1.26.2
```

