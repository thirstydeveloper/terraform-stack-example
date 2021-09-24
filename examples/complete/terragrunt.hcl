dependency "fixture" {
  config_path = "./fixture"
}

inputs = {
  id = dependency.fixture.outputs.id
}

terraform {
  source = "../../src/example"
}
