require_relative "../examples/ruby-app/hello_world_unit_testable"
require "test/unit"
require 'webrick'
require 'net/http'

# A unit test for the web server in examples/ruby-app/hello_world_unit_testable.rb
class TestWebServer < Test::Unit::TestCase
  def initialize(test_method_name)
    super(test_method_name)
    @handlers = Handlers.new
  end

  def test_unit_hello
    status, type, body = @handlers.handle("/")
    assert_equal(200, status)
    assert_equal('text/plain', type)
    assert_equal('Hello, World', body)
  end
end
