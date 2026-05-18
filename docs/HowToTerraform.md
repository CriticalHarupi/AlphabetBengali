# How I manage cloud infrastructure with Terraform

Based on the [ARC doc](./RunTestsWithARC.md).  

Terraform is an IaC (Infrastructure as Code) tool.
Instead of running `gcloud` commands manually, you write `.tf` files describing what cloud infrastructure you want, and Terraform manages the resources to make it happen.

It covers the infrastructure layer: GKE cluster, Artifact Registry, IAM, workload identity, etc.
It does not cover what runs inside the cluster — that is still done with `helm` and `kubectl`.

## Setup Local Working Env

Install Terraform using winget:

```powershell
winget install Hashicorp.Terraform
```

Authenticate Terraform with GCP using application default credentials:

```powershell
gcloud auth application-default login
```

## Terraform Files

All files are in the `terraform/` directory.

`main.tf` — configures the Google provider and which version to use.  
`variables.tf` — declares variables (project ID, region, zone, GitHub repo).  
`cluster.tf` — GKE cluster and node pool with autoscaling.  
`artifact_registry.tf` — Docker repository for storing the runner image.  
`iam.tf` — service account, workload identity pool and provider, IAM bindings.  
`outputs.tf` — prints useful values after apply (workload identity provider, service account email, registry URL).  

## Apply

Initialize Terraform (downloads the Google provider plugin):

```powershell
cd terraform
terraform init
```

Preview what will be created without touching anything:

```powershell
terraform plan
```

Create all the resources:

```powershell
terraform apply
```

## Workload Identity Pool Soft-Delete

I created these before and they can be imported, rather than created again:

```powershell
terraform import google_iam_workload_identity_pool.github projects/project-b8ee474a-9624-49ce-b08/locations/global/workloadIdentityPools/github-pool
terraform import google_iam_workload_identity_pool_provider.github projects/project-b8ee474a-9624-49ce-b08/locations/global/workloadIdentityPools/github-pool/providers/github-provider
```

I tried to delete them, but GCP soft deleted them and I cannot create one with same name before they are really gone. So I just import them to Terraform.  

Then run `terraform apply` again.

## After Apply

These steps are still done manually after `terraform apply`:

Connect kubectl to the new cluster:

```powershell
gcloud container clusters get-credentials alphabetbengali --zone=asia-northeast1-a --project=project-b8ee474a-9624-49ce-b08
```

Install ARC controller:

```powershell
helm install arc `
  --namespace arc-systems `
  --create-namespace `
  oci://ghcr.io/actions/actions-runner-controller-charts/gha-runner-scale-set-controller
```

Trigger `build-runner.yml` from GitHub Actions to push the runner image to Artifact Registry.

Install the runner scale set:

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

## Destroy

```powershell
terraform destroy
```

This deletes the cluster, Artifact Registry, service account, and IAM resources.
The workload identity pool will be soft-deleted for 30 days (see note above).  
And of course the ARC too, as it's on the cluster.  
