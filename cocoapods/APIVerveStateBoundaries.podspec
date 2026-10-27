Pod::Spec.new do |s|
  s.name             = 'APIVerveStateBoundaries'
  s.version          = '1.2.3'
  s.summary          = 'State Boundaries provides GeoJSON data and Census identifiers for US states using a two-letter postal code.'
  s.description      = <<-'DESC'
State Boundaries provides GeoJSON data and Census identifiers for US states using a two-letter postal code. Responses include the state name, FIPS code, and GEOID; paid plans add border polygon coordinates and land area measurements.

This iOS/macOS SDK provides a convenient way to access the State Boundaries from APIVerve.
Supports both callback-based and async/await patterns.
                       DESC

  s.homepage         = 'https://apiverve.com/marketplace/stateboundaries?utm_source=cocoapods&utm_medium=referral&utm_content=homepage'
  s.license          = { :type => 'MIT', :file => 'LICENSE' }
  s.author           = { 'APIVerve' => 'hello@apiverve.com' }
  # The SDK repo holds every language; the pod lives under cocoapods/ with its own tag.
  s.source           = { :git => 'https://github.com/apiverve/stateboundaries-api.git', :tag => "cocoapods/v#{s.version}" }

  s.ios.deployment_target = '15.0'
  s.osx.deployment_target = '12.0'

  s.swift_version = '5.0'

  s.source_files = 'cocoapods/Sources/APIVerveStateBoundaries/**/*.swift'

  s.frameworks = 'Foundation'

  # APIResponse, APIVerveError and the request plumbing, shared by every APIVerve pod.
  s.dependency 'APIVerveCore', '~> 1.2'
end
