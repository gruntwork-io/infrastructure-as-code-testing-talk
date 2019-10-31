const http = require('http');
const https = require('https');

/**
 * The entrypoint for this Lambda function. This will be called by API Gateway. Proxies a URL specified via an
 * environment variable.
 *
 * @param event
 * @param context
 * @param callback
 * @returns {Promise<void>}
 */
exports.handler = async (event, context, callback) => {
  console.log('Received an event:', JSON.stringify(event, null, 2));

  const urlToProxy = process.env.URL_TO_PROXY;
  if (urlToProxy) {
    const response = await httpGet(urlToProxy);
    callback(null, {statusCode: response.statusCode, headers: response.headers, body: response.body});
  } else {
    callback(null, {statusCode: 500, body: "Required environment variable URL_TO_PROXY not configured."});
  }
};

/**
 * Make an HTTP GET request to the given URL and return a Promise that contains the response, which will be of type
 * http.IncomingMessage, along with an extra body field that contains the response body.
 * @param url
 * @returns {Promise<any>}
 */
function httpGet(url) {
  return new Promise((resolve, reject) => {
    const httpLib = url.startsWith("https") ? https : http;
    httpLib.get(url, (response) => {
      let body = '';
      response.on('data', chunk => body += chunk);
      response.on('end', () => resolve(Object.assign({}, response, { body })));
    }).on('error', reject);
  });
}
