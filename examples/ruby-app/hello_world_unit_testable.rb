require 'webrick'

# A basic web server that responds with "Hello, World!" to all requests
class WebServer < WEBrick::HTTPServlet::AbstractServlet
  def do_GET(request, response)
    handlers = Handlers.new
    status_code, content_type, body = handlers.handle(request.path)

    response.status = status_code
    response['Content-Type'] = content_type
    response.body = body
  end
end

# The core implementation of the web server. This class that takes in simple values (e.g., request path) and returns
# simple values (arrays, strings), so it's easy to unit test.
class Handlers
  def handle(path)
    [200, 'text/plain', 'Hello, World']
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
