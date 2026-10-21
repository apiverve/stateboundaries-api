# frozen_string_literal: true

Gem::Specification.new do |spec|
  spec.name          = "apiverve_stateboundaries"
  spec.version       = "1.2.2"
  spec.authors       = ["APIVerve"]
  spec.email         = ["hello@apiverve.com"]

  spec.summary       = "State Boundaries API - Ruby Client"
  spec.description   = "State Boundaries provides GeoJSON data and Census identifiers for US states using a two-letter postal code. Responses include the state name, FIPS code, and GEOID; paid plans add border polygon coordinates and land area measurements."
  spec.homepage      = "https://apiverve.com/marketplace/stateboundaries?utm_source=ruby&utm_medium=referral&utm_content=homepage"
  spec.license       = "MIT"
  spec.required_ruby_version = ">= 2.7.0"

  spec.metadata["homepage_uri"] = spec.homepage
  spec.metadata["source_code_uri"] = "https://github.com/apiverve/stateboundaries-API/tree/main/ruby"
  spec.metadata["changelog_uri"] = "https://apiverve.com/changelog"
  spec.metadata["documentation_uri"] = "https://docs.apiverve.com/ref/stateboundaries"

  spec.files = Dir.glob("{lib}/**/*") + %w[README.md LICENSE]
  spec.require_paths = ["lib"]

  spec.add_dependency "faraday", ">= 1.0", "< 3.0"
  spec.add_dependency "faraday-multipart", "~> 1.0"
  spec.add_dependency "json", "~> 2.0"
end
