require 'webrick'
require 'net/http'

# A web server that has a dependency on the outside world: in particular, it makes a call to example.com and proxies
# the response.
class WebServer < WEBrick::HTTPServlet::AbstractServlet
  def initialize(server, *options)
    super(server, options)
    @handlers = Handlers.new
  end

  def do_GET(request, response)
    status, type, body = @handlers.handle(request.path)
    response.status = status
    response['Content-Type'] = type
    response.body = body
  end
end

# The core implementation of the web server. It has a dependency on the outside world: in particular, it makes a call
# to example.com and proxies the response.
class Handlers
  def handle(path)
    uri = URI("http://www.example.com")
    response = Net::HTTP.get_response(uri)
    [response.code.to_i, response['Content-Type'], response.body]
  end
end

# Start the server on port 8000. This code only runs if called directly from the CLI, but not if required from another
# file.
if __FILE__ == $0
  server = WEBrick::HTTPServer.new :Port => 8000
  server.mount '/', WebServer
  trap 'INT' do server.shutdown end
  server.start
end
