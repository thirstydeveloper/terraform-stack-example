output "id" {
  value = join("-", [var.fixture_namespace, var.fixture_environment])
}
