require 'webrick'

# A basic web server that responds with "Hello, World!" to all requests
class WebServer < WEBrick::HTTPServlet::AbstractServlet
  def do_GET(request, response)
    response.status = 200
    response['Content-Type'] = 'text/plain'
    response.body = 'Hello, World!'
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
