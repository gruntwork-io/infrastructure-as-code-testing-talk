/**
 * The entrypoint for this Lambda function. This will be called by API Gateway. Returns a simple "Hello, World"
 * JSON response to represent a typical backend web service.
 *
 * @param event
 * @param context
 * @param callback
 * @returns {Promise<void>}
 */
exports.handler = (event, context, callback) => {
  console.log('Received an event:', JSON.stringify(event, null, 2));
  const response = {
    statusCode: 200,
    body: JSON.stringify({text: "Hello, World!"}),
    headers: {"Content-Type": "application/json"}
  };
  callback(null, response);
};
