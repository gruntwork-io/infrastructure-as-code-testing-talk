require_relative "../examples/ruby-app/proxy_unit_testable"
require "test/unit"
require 'webrick'
require 'net/http'

# A unit test for the web server in examples/ruby-app/proxy_unit_testable.rb
class TestWebServer < Test::Unit::TestCase
  def initialize(test_method_name)
    super(test_method_name)
    mock_proxy = MockProxy.new([200, 'text/html', 'mock example.org'])
    @handlers = Handlers.new(mock_proxy)
  end

  def test_proxy
    status_code, content_type, body = @handlers.handle("/")
    assert_equal(200, status_code)
    assert_equal('text/html', content_type)
    assert_equal('mock example.org', body)
  end
end

# A mock implementation of the Proxy class. Returns a fixed response.
class MockProxy
  def initialize(response)
    @response = response
  end

  def proxy
    @response
  end
end
