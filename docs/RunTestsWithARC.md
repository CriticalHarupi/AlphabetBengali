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
  --create-namespace `
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

But with that, I need to install the ARC with some extra configs:  

```powershell
helm install go-test `
  --namespace arc-runners `
  --create-namespace `
  --set githubConfigUrl="https://github.com/CriticalHarupi/AlphabetBengali" `
  --set githubConfigSecret.github_token="MY_PAT" `
  --set runnerScaleSetName="go-test" `
  -f k8s/arc-values.yml `
  oci://ghcr.io/actions/actions-runner-controller-charts/gha-runner-scale-set
```

And `k8s/arc-values.yml` is like this:  

```yml
containerMode:
  type: kubernetes
  kubernetesModeWorkVolumeClaim:
    accessModes: ["ReadWriteOnce"]
    storageClassName: "standard"
    resources:
      requests:
        storage: 1Gi
  kubernetesModeServiceAccountName: "default"

template:
  spec:
    securityContext:
      fsGroup: 123
```

This tells the ARC to create a new pod when the ci workflow uses `container` keyword, and run that job in the specified image.  

## Solutions 3

Solution 2 is much faster then solution 1, but it still have some problems (or trade-offs, depends on app and ci task).  

For example, we still need to pull the golang image, although it will be cached in cluster node, but that could be updated and still need to be loaded by the pod for running our ci task.  

Another solution is to prepare an image that includes all things we need for our test. (ARC and go, for this project)  
And if we store that image on cloud, we can config ARC to use that image for our ci tasks.  

### Set up Cloud Env

First, we need an artifact repository on cloud to store our image.  

```powershell
gcloud artifacts repositories create alphabetbengali `
  --repository-format=docker `
  --location=asia-northeast1
```

Then, we build the image and register it.  
This is done by [build-runner.yml](../.github/workflows/build-runner.yml).  
This workflow will not be executed a lot (when docker file is updated).  

With this image stored in cloud, our ci task only needs to load it to run the tests.

And now we need to install the ARC like this:  

```powershell
helm install go-test `
  --namespace arc-runners `
  --create-namespace `
  --set githubConfigUrl="https://github.com/CriticalHarupi/AlphabetBengali" `
  --set githubConfigSecret.github_token="MY_PAT" `
  --set runnerScaleSetName="go-test" `
  -f k8s/arc-runner-image-values.yml `
  oci://ghcr.io/actions/actions-runner-controller-charts/gha-runner-scale-set
```

In which, the `arc-runner-image-values.yml` is like this:  
(It tells ARC what runner image it should use, that is, the one we stored in cloud)  

```yml
template:
  spec:
    containers:
      - name: runner
        image: asia-northeast1-docker.pkg.dev/project-b8ee474a-9624-49ce-b08/alphabetbengali/runner:latest
        command: ["/home/runner/run.sh"]
```

Thus, in the ci task yml, we don't need both setup-go part or the container part.  
Because the runner image had it prepared.  