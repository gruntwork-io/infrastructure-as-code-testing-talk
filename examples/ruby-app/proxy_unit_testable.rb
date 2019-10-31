require 'webrick'
require 'net/http'

# A web server that proxies a URL
class WebServer < WEBrick::HTTPServlet::AbstractServlet
  def do_GET(request, response)
    web_service = Proxy.new("http://www.example.org")
    handlers = Handlers.new(web_service)

    status_code, content_type, body = handlers.handle(request.path)

    response.status = status_code
    response['Content-Type'] = content_type
    response.body = body
  end
end

# The core implementation of the web server. It takes in dependencies, such as the proxy class, as inputs, so those can
# be replaced at test time with mocks ("dependency injection") to make unit testing easier.
class Handlers
  def initialize(proxy)
    @proxy = proxy
  end

  def handle(path)
    @proxy.proxy
  end
end

# This class proxies a given URL
class Proxy
  def initialize(url)
    @uri = URI(url)
  end

  def proxy
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
