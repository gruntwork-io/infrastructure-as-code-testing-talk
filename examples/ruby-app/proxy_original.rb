require 'webrick'
require 'net/http'

# A web server that proxies a URL
class WebServer < WEBrick::HTTPServlet::AbstractServlet
  def do_GET(request, response)
    handlers = Handlers.new
    status_code, content_type, body = handlers.handle(request.path)

    response.status = status_code
    response['Content-Type'] = content_type
    response.body = body
  end
end

# The core implementation of the web server.
class Handlers
  def handle(path)
    uri = URI("http://www.example.org")
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
