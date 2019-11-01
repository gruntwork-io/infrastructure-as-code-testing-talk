require_relative "../examples/ruby-app/web_service_with_dependency_testable"
require "test/unit"
require 'webrick'
require 'net/http'

# An integration test for the web server in examples/ruby-app/web_service_with_dependencies_testable.rb
class TestWebServer < Test::Unit::TestCase
  def initialize(test_method_name)
    super(test_method_name)
    web_service = WebService.new("http://www.example.com")
    @handlers = Handlers.new(web_service)
  end

  def test_integration_web_service
    do_integration_test('/', lambda { |response|
      assert_equal(200, response.code.to_i)
      assert_include(response['Content-Type'], 'text/html')
      assert_include(response.body, '<h1>Example Domain</h1>')
    })
  end

  def do_integration_test(path, check_response)
    port = 8000
    server = WEBrick::HTTPServer.new :Port => port
    server.mount '/', WebServer

    begin
      # Start the web server in a separate thread so it
      # doesn't block the test
      thread = Thread.new do
        server.start
      end

      # Make an HTTP request to the web server at the
      # specified path
      uri = URI("http://localhost:#{port}#{path}")
      response = Net::HTTP.get_response(uri)

      # Use the specified check_response lambda to validate
      # the response
      check_response.call(response)
    ensure
      # Shut the server and thread down at the end of the
      # test
      server.shutdown
      thread.join
    end
  end
end
