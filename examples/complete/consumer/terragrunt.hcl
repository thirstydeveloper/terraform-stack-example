dependency "fixture" {
  config_path = "../fixture"
}

dependency "test" {
  config_path = "../"
}

inputs = {
  id = dependency.fixture.outputs.id
}
