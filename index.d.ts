/**
 * A [k6](https://go.k6.io/k6) extension for mocking HTTP(s) servers during test development.
 * The design of the library was inspired by [Express](https://expressjs.com/). If you already know Express framework, using this library should be very simple.
 *
 * ## Features
 *
 * - Start mock HTTP server inside of a k6 process
 * - Familiar, Express like mock route definitions
 * - Almost transparent for test scripts: change import from `k6/http` to `k6/x/mock/http`
 *
 * ## http module
 *
 * The `k6/x/mock/http` module is a drop-in replacement for `k6/http` module. During the k6 test development simply replace `k6/http` imports with `k6/x/mock/http`. Each URL's host and port part will be automatically replaced with the mock server's host and port value. Other parts of the request are remain untouched.
 *
 * ## mock module
 *
 * The `k6/x/mock` module's default export is an Express like default mock server with the usual HTTP method functions (get, post, ..) and use function for defining middlewares.
 *
 * You can add route definitions both in [Init and VU stages](https://k6.io/docs/using-k6/test-life-cycle/#init-and-vu-stages).
 *
 */

/**
 * NextFunc defines callback function's `next` parameter function.
 * Calling from middleware enables processing next middleware.
 */
export type NextFunc = () => void;

/**
 * The `req` object represents the HTTP request and has properties for the request query string, parameters, body, HTTP headers, and so on.
 *
 * In this documentation and by convention, the object is always referred to as `req` (and the HTTP response is `res`) but its actual name is determined by the parameters to the callback function in which you’re working.
 *
 * For example:
 *
 * ```js
 * mock.get("/user/{id}", function (req, res) {
 *   res.send("user " + req.params.id);
 * });
 * ```
 *
 * But you could just as well have:
 *
 * ```js
 * mock.get("/user/{id}", function (request, response) {
 *   response.send("user " + request.params.id);
 * });
 * ```
 *
 * The `req` object contains a `Request` field which is a golang's [http.Request](https://golang.org/pkg/net/http/#Request) instance.
 */
export interface Request {
  /**
   * Contains key-value pairs of data submitted in the request body.
   * By default, it is undefined, and is populated when the request
   * Content-Type is `application/json`.
   */
  body: Record<string, any> | undefined;

  /**
   * This property is an object that contains cookies sent by the request.
   */
  cookies: Record<string, string>;

  /**
   * Contains a string corresponding to the HTTP method of the request: GET, POST, PUT, and so on.
   */
  method: string;

  /**
   * This property is an object containing properties mapped to the named route parameters.
   * For example, if you have the route /user/{name}, then the “name” property is available as req.params.name.
   * This object defaults to {}.
   */
  params: Record<string, string>;

  /**
   * Contains the path part of the request URL.
   */
  path: string;

  /**
   * Contains the request protocol string: either http or (for TLS requests) https.
   */
  protocol: string;

  /**
   * This property is an object containing a property for each query string parameter in the route.
   *
   * For example:
   *
   * ```js
   * // GET /search?q=tobi+ferret
   * console.dir(req.query.q);
   * // => 'tobi ferret'
   *
   * // GET /shoes?color=blue&color=black&color=red
   * console.dir(req.query.color);
   * // => ['blue', 'black', 'red']
   * ```
   */
  query: Record<string, any>;

  /**
   * Returns the specified HTTP request header field (case-insensitive match).
   *
   * @param field the header field name
   * @returns the header field value.
   */
  get: (field: string) => string;

  /**
   * Returns the specified HTTP request header field (case-insensitive match).
   *
   * @param field the header field name
   * @returns the header field value.
   */
  header: (field: string) => string;
}

/**
 * The `res` object represents the HTTP response that a server sends when it gets an HTTP request.
 *
 * In this documentation and by convention, the object is always referred to as `res` (and the HTTP request is `req`) but its actual name is determined by the parameters to the callback function in which you’re working.
 *
 * For example:
 *
 * ```js
 * mock.get("/user/{id}", function (req, res) {
 *   res.send("user " + req.params.id);
 * });
 * ```
 *
 * But you could just as well have:
 *
 * ```js
 * mock.get("/user/{id}", function (request, response) {
 *   response.send("user " + request.params.id);
 * });
 * ```
 *
 * The `res` object contains a `Writer` field which is a golang's [http.ResponseWriter](https://golang.org/pkg/net/http/#ResponseWriter) instance.
 */
export interface Response {
  /**
   * Appends the specified value to the HTTP response header field. If the header is not already set, it creates the header with the specified value.
   *
   * @param field the header field name
   * @param value the value to append
   * @returns The instance for fluent/chaining API
   */
  append: (field: string, value: string) => Response;

  /**
   * Sends a JSON response. This method sends a response (with the correct content-type) that is the parameter converted to a JSON string.
   *
   * @param body the object to send
   * @returns The instance for fluent/chaining API
   */
  json: (body: Record<string, any>) => Response;

  /**
   * Sends a plain text response. This method sends a response (with the correct content-type) that is the string formatting result.
   *
   * @param format go format string
   * @param v format values (if any)
   * @returns The instance for fluent/chaining API
   */
  text: (format: string, v?: any[]) => Response;

  /**
   * Sends a HTML text response. This method sends a response (with the correct content-type) that is the body string paramter.
   * @param body the string to send
   * @returns The instance for fluent/chaining API
   */
  html: (body: string) => Response;

  /**
   * Sends a binray response. This method sends a response (with the "application/octet-stream" content-type) that is the body paramter.
   *
   * @param body the data to send
   * @returns The instance for fluent/chaining API
   */
  binary: (body: string | number[] | ArrayBuffer) => Response;

  /**
   * Sends the HTTP response.
   *
   * When the parameter is a ArrayBuffer or number[], the method sets the Content-Type response header field to “application/octet-stream”.
   * When the parameter is a String, the method sets the Content-Type to “text/html”.
   * Otherwise the method sets the Content-Type to "application/json" and convert paramter to JSON representation before sending.
   *
   * @param body the data to send
   * @returns The instance for fluent/chaining API
   */
  send: (body: string | number[] | ArrayBuffer) => Response;

  /**
   * Sets the HTTP status for the response.
   *
   * @param code the satus code value
   * @returns The instance for fluent/chaining API
   */
  status: (code: number) => Response;

  /**
   * Sets the Content-Type HTTP header to the MIME type as from mime parameter.
   *
   * @params mime the content type
   * @returns The instance for fluent/chaining API
   */
  type: (mime: string) => Response;

  /**
   * Adds the header field to the Vary response header.
   *
   * @param header the header filed name
   * @returns The instance for fluent/chaining API
   */
  vary: (header: string) => Response;

  /**
   * Sets the response’s HTTP header field to value.
   *
   * @param field the header field name
   * @param value the value to set
   * @returns The instance for fluent/chaining API
   */
  set: (field: string, value: string) => Response;

  /**
   * Redirects to the URL, with specified status, a positive integer that corresponds to an HTTP status code.
   *
   * @param code the HTTP status code (301, 302, ...)
   * @param loc the location to redirect
   * @returns The instance for fluent/chaining API
   */
  redirect: (code: number, loc: string) => Response;
}

/**
 * CallbackFunc defines middleware and request handler callback function.
 *
 * @param req the request object
 * @param res the response object
 * @param next indicating the next middleware function
 */
export type CallbackFunc = (req: Request, res: Response, next: NextFunc) => void;

/**
 * A server object represents a web server.
 */
export declare class Server {
  /**
   * Creates a new server instance.
   */
  constructor();

  /**
   * This method is just like the get,head,post,... methods, except that it matches all HTTP methods (verbs).
   *
   * You can provide multiple callback functions that behave just like middleware.
   *
   * @param path The path for which the callback function is invoked (string or path pattern)
   * @param callback Callback functions (middlewares or request handler)
   * @returns The instance for fluent/chaining API
   */
  all(path: string, ...callback: CallbackFunc[]): Server;

  /**
   * Routes HTTP GET requests to the specified path with the specified callback functions.
   *
   * You can provide multiple callback functions that behave just like middleware.
   *
   * @param path The path for which the callback function is invoked (string or path pattern)
   * @param callback Callback functions (middlewares or request handler)
   * @returns The instance for fluent/chaining API
   */
  get(path: string, ...callback: CallbackFunc[]): Server;

  /**
   * Routes HTTP HEAD requests to the specified path with the specified callback functions.
   *
   * You can provide multiple callback functions that behave just like middleware.
   *
   * @param path The path for which the callback function is invoked (string or path pattern)
   * @param callback Callback functions (middlewares or request handler)
   * @returns The instance for fluent/chaining API
   */
  head(path: string, ...callback: CallbackFunc[]): Server;

  /**
   * Routes HTTP POST requests to the specified path with the specified callback functions.
   *
   * You can provide multiple callback functions that behave just like middleware.
   *
   * @param path The path for which the callback function is invoked (string or path pattern)
   * @param callback Callback functions (middlewares or request handler)
   * @returns The instance for fluent/chaining API
   */
  post(path: string, ...callback: CallbackFunc[]): Server;

  /**
   * Routes HTTP PUT requests to the specified path with the specified callback functions.
   *
   * You can provide multiple callback functions that behave just like middleware.
   *
   * @param path The path for which the callback function is invoked (string or path pattern)
   * @param callback Callback functions (middlewares or request handler)
   * @returns The instance for fluent/chaining API
   */
  put(path: string, ...callback: CallbackFunc[]): Server;

  /**
   * Routes HTTP PATCH requests to the specified path with the specified callback functions.
   *
   * You can provide multiple callback functions that behave just like middleware.
   *
   * @param path The path for which the callback function is invoked (string or path pattern)
   * @param callback Callback functions (middlewares or request handler)
   * @returns The instance for fluent/chaining API
   */
  patch(path: string, ...callback: CallbackFunc[]): Server;

  /**
   * Routes HTTP `DELETE` requests to the specified path with the specified callback functions.
   *
   * You can provide multiple callback functions that behave just like middleware.
   *
   * @param path The path for which the callback function is invoked (string or path pattern)
   * @param callback Callback functions (middlewares or request handler)
   * @returns The instance for fluent/chaining API
   */
  delete(path: string, ...callback: CallbackFunc[]): Server;

  /**
   * Routes HTTP OPTIONS requests to the specified path with the specified callback functions.
   *
   * You can provide multiple callback functions that behave just like middleware.
   *
   * @param path The path for which the callback function is invoked (string or path pattern)
   * @param callback Callback functions (middlewares or request handler)
   * @returns The instance for fluent/chaining API
   */
  options(path: string, ...callback: CallbackFunc[]): Server;

  /**
   * Uses the specified middleware function or functions.
   *
   * @param path The path for which the callback function is invoked (string or path pattern)
   * @param callback Callback functions (middlewares or request handler)
   * @returns The instance for fluent/chaining API
   */
  use(path: string, ...callback: CallbackFunc[]): Server;

  /**
   * Starts the server.
   *
   * @param port TCP port number, if 0 or missing then random unused port will be allocated
   * @param host host name or IP address for listening on, default 127.0.0.1
   * @returns The instance for fluent/chaining API
   */
  start(port?: number, host?: string): Server;
}

/**
 * A module object represents module default object.
 */
export declare class Module {
  /**
   * This method is just like the get,head,post,... methods, except that it matches all HTTP methods (verbs).
   *
   * You can provide multiple callback functions that behave just like middleware.
   *
   * @param path The path for which the callback function is invoked (string or path pattern)
   * @param callback Callback functions (middlewares or request handler)
   * @returns The instance for fluent/chaining API
   */
  all(path: string, ...callback: CallbackFunc[]): Server;

  /**
   * Routes HTTP GET requests to the specified path with the specified callback functions.
   *
   * You can provide multiple callback functions that behave just like middleware.
   *
   * @param path The path for which the callback function is invoked (string or path pattern)
   * @param callback Callback functions (middlewares or request handler)
   * @returns The instance for fluent/chaining API
   */
  get(path: string, ...callback: CallbackFunc[]): Server;

  /**
   * Routes HTTP HEAD requests to the specified path with the specified callback functions.
   *
   * You can provide multiple callback functions that behave just like middleware.
   *
   * @param path The path for which the callback function is invoked (string or path pattern)
   * @param callback Callback functions (middlewares or request handler)
   * @returns The instance for fluent/chaining API
   */
  head(path: string, ...callback: CallbackFunc[]): Server;

  /**
   * Routes HTTP POST requests to the specified path with the specified callback functions.
   *
   * You can provide multiple callback functions that behave just like middleware.
   *
   * @param path The path for which the callback function is invoked (string or path pattern)
   * @param callback Callback functions (middlewares or request handler)
   * @returns The instance for fluent/chaining API
   */
  post(path: string, ...callback: CallbackFunc[]): Server;

  /**
   * Routes HTTP PUT requests to the specified path with the specified callback functions.
   *
   * You can provide multiple callback functions that behave just like middleware.
   *
   * @param path The path for which the callback function is invoked (string or path pattern)
   * @param callback Callback functions (middlewares or request handler)
   * @returns The instance for fluent/chaining API
   */
  put(path: string, ...callback: CallbackFunc[]): Server;

  /**
   * Routes HTTP PATCH requests to the specified path with the specified callback functions.
   *
   * You can provide multiple callback functions that behave just like middleware.
   *
   * @param path The path for which the callback function is invoked (string or path pattern)
   * @param callback Callback functions (middlewares or request handler)
   * @returns The instance for fluent/chaining API
   */
  patch(path: string, ...callback: CallbackFunc[]): Server;

  /**
   * Routes HTTP `DELETE` requests to the specified path with the specified callback functions.
   *
   * You can provide multiple callback functions that behave just like middleware.
   *
   * @param path The path for which the callback function is invoked (string or path pattern)
   * @param callback Callback functions (middlewares or request handler)
   * @returns The instance for fluent/chaining API
   */
  delete(path: string, ...callback: CallbackFunc[]): Server;

  /**
   * Routes HTTP OPTIONS requests to the specified path with the specified callback functions.
   *
   * You can provide multiple callback functions that behave just like middleware.
   *
   * @param path The path for which the callback function is invoked (string or path pattern)
   * @param callback Callback functions (middlewares or request handler)
   * @returns The instance for fluent/chaining API
   */
  options(path: string, ...callback: CallbackFunc[]): Server;

  /**
   * Uses the specified middleware function or functions.
   *
   * @param path The path for which the callback function is invoked (string or path pattern)
   * @param callback Callback functions (middlewares or request handler)
   * @returns The instance for fluent/chaining API
   */
  use(path: string, ...callback: CallbackFunc[]): Server;

  /**
   * Resolve URL to mock URL, replace host and port with mock server's host and por value.
   *
   * Only external URLs are resolved, URLs are already pointing to mock server will remain untouched.
   *
   * @param url original URL string
   * @returns resolved URL, points to mock server
   */
  resolve(url: string): string;
}

/**
 * default export
 */
declare const mock: Module;

/**
 * default export
 */
export default mock;
