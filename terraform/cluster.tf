resource "google_container_cluster" "main" {
  name                     = "alphabetbengali"
  location                 = var.zone
  remove_default_node_pool = true
  initial_node_count       = 1
  deletion_protection      = false
}

resource "google_container_node_pool" "main" {
  name     = "default-pool"
  location = var.zone
  cluster  = google_container_cluster.main.name

  initial_node_count = 1

  autoscaling {
    min_node_count = 1
    max_node_count = 3
  }

  node_config {
    disk_size_gb = 30
    disk_type    = "pd-standard"
  }
}
