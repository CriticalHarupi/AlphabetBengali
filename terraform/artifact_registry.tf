resource "google_artifact_registry_repository" "runner" {
  location      = var.region
  repository_id = "alphabetbengali"
  format        = "DOCKER"
}
