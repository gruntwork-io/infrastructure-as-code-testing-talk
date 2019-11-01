require 'webrick'
require 'net/http'

# A web server that has a dependency on the outside world: in particular, it makes a call to example.com and proxies
# the response. In this case, the part of the code that depends on the outside world is injectable, which makes unit
# testing easier.
class WebServer < WEBrick::HTTPServlet::AbstractServlet
  def initialize(server, *options)
    super(server, options)
    web_service = WebService.new("http://www.example.com")
    @handlers = Handlers.new(web_service)
  end

  def do_GET(request, response)
    status, type, body = @handlers.handle(request.path)
    response.status = status
    response['Content-Type'] = type
    response.body = body
  end
end

# The core implementation of the web server. It has a dependency on the outside world: in particular, it makes a call
# to example.com and proxies the response. In this case, the part of the code that depends on the outside world is
# injectable, which makes unit testing easier.
class Handlers
  def initialize(web_service)
    @web_service = web_service
  end

  def handle(path)
    @web_service.fetch_data
  end
end

# This class fetches data from an external web service
class WebService
  def initialize(url)
    @uri = URI(url)
  end

  def fetch_data
    response = Net::HTTP.get_response(@uri)
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
