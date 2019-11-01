require_relative "../examples/ruby-app/web_service_with_dependency_testable"
require "test/unit"
require 'webrick'
require 'net/http'

# A unit test for the web server in examples/ruby-app/web_service_with_dependencies_testable.rb
class TestWebServer < Test::Unit::TestCase
  def initialize(test_method_name)
    super(test_method_name)
    mock_web_service = MockWebService.new([200, 'text/html', 'mock example.com'])
    @handlers = Handlers.new(mock_web_service)
  end

  def test_unit_web_service
    status, type, body = @handlers.handle("/")
    assert_equal(200, status)
    assert_equal('text/html', type)
    assert_equal('mock example.com', body)
  end
end

# A mock implementation of the WebService class. Returns a mock response.
class MockWebService < WebService
  def initialize(response)
    super("http://www.mock-url-not-used.com")
    @response = response
  end

  def fetch_data
    @response
  end
end
