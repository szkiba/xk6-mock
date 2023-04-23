// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

/**
 * A [k6](https://go.k6.io/k6) extension enables mocking HTTP(S) servers during test development.
 * The design of the library was inspired by [Express.js](https://expressjs.com/).
 * If you have already known Express.js framework, using this library should be very simple.
 */

/**
 * Optional flags for `mock` function.
 * 
 * Passing this options object as first (or last) parameter to `mock` function (or to `Application` constructor) may
 * change its behavior.
 * 
 * @example
 * mock("https://example.com", callback, { sync:true });
*/
export interface MockOptions {
  /**
   * True value indicates synchronous mode operation. You should use it for synchronous k6 http API.
   */
  sync: boolean

  /**
   * True value indicaes that given mock definition should be ignored.
   */
  skip: boolean
}

/**
 * Create URL mock definition.
 * 
 * This function will create a new Express.js like web application and pass it to the provided
 * callback function for programming. After mock programming is done, new mock HTTP server will be start.
 * Whan you use http API from mock module, all matching URLs will be directed to this mock server.
 * 
 * You can create as many mock definitions (server) as you want.
 * 
 * You can disable the given mock definition quickly by passing options parameter with `skip` set to true.
 * ```JavaScript
 * mock(
 *   'https://example.com',
 *   app => {
 *     app.get('/', (req, res) => {
 *       res.json({ greeting: 'Hello World!' })
 *     })
 *   },
 *   { skip: true }
 * )
 * ```
 * 
 * Alternatively you can put `.skip` after function name:
 * ```JavaScript
 * mock.skip('https://example.com', app => {
 *  app.get('/', (req, res) => {
 *    res.json({ greeting: 'Hello World!' })
 *  })
 * })
 * ```
 * 
 * @param target the URL or URL prefix to be mocked
 * @param callback function to for defining route definitions for mock server
 * @param options optional flags (`sync`, `skip`)
 */
export function mock(target: String, callback: (app: Application) => void, options?: MockOptions): void;

/**
 * Deactivate URL mocking.
 * 
 * This function will remove mock definition associated to given URL and stop the related HTTP server.
 * 
 * @param target the URL or URL prefix of mock definition to be remove
 */
export function unmock(target: String): void;

// muxpress ------------------------------------------------------------------------

/**
 * Middleware defines middleware and request handler callback function.
 *
 * @param req the request object
 * @param res the response object
 * @param next calling from middleware enables processing next middleware
 */
export type Middleware = (req: Request, res: Response, next: () => void) => void;

/**
 * An application object represents a web application.
 * 
 * The following example starts a server and listens for connections on port 3000.
 * The application responds with a JSON object for requests to the root URL.
 * All other routes are answered with a 404 not found message.
 *
 * In this example, the name of the constructor is `Application`, but the name you use is up to you.
 * 
 * @example
 * const app = new Application()
 *
 * app.get('/', (req, res) => {
 *   res.json({message:"Hello World!"})
 * })
 *
 * app.listen(3000)
 */
export class Application {
  /**
   * Creates a new application instance.
   */
  constructor();

  /**
   * Routes HTTP GET requests to the specified path with the specified middleware functions.
   *
   * You can provide multiple middleware functions.
   *
   * @param path The path for which the middleware function is invoked (string or path pattern)
   * @param middleware Middleware functions
   */
  get(path: string, ...middleware: Middleware[]): void;

  /**
   * Routes HTTP HEAD requests to the specified path with the specified middleware functions.
   *
   * You can provide multiple middleware functions.
   *
   * @param path The path for which the middleware function is invoked (string or path pattern)
   * @param middleware Middleware functions
   */
  head(path: string, ...middleware: Middleware[]): void;

  /**
   * Routes HTTP POST requests to the specified path with the specified middleware functions.
   *
   * You can provide multiple middleware functions.
   *
   * @param path The path for which the middleware function is invoked (string or path pattern)
   * @param middleware Middleware functions
   */
  post(path: string, ...middleware: Middleware[]): void;

  /**
   * Routes HTTP PUT requests to the specified path with the specified middleware functions.
   *
   * You can provide multiple middleware functions.
   *
   * @param path The path for which the middleware function is invoked (string or path pattern)
   * @param middleware Middleware functions
   */
  put(path: string, ...middleware: Middleware[]): void;

  /**
   * Routes HTTP PATCH requests to the specified path with the specified middleware functions.
   *
   * You can provide multiple middleware functions.
   *
   * @param path The path for which the middleware function is invoked (string or path pattern)
   * @param middleware Middleware functions
   */
  patch(path: string, ...middleware: Middleware[]): void;

  /**
   * Routes HTTP `DELETE` requests to the specified path with the specified middleware functions.
   *
   * You can provide multiple middleware functions.
   *
   * @param path The path for which the middleware function is invoked (string or path pattern)
   * @param middleware Middleware functions
   */
  delete(path: string, ...middleware: Middleware[]): void;

  /**
   * Routes HTTP OPTIONS requests to the specified path with the specified middleware functions.
   *
   * You can provide multiple middleware functions.
   *
   * @param path The path for which the middleware function is invoked (string or path pattern)
   * @param middleware Middleware functions
   */
  options(path: string, ...middleware: Middleware[]): void;

  /**
   * Uses the specified middleware function or functions.
   *
   * @param path The path for which the middleware function is invoked (string or path pattern)
   * @param middleware Middleware functions
   */
  use(path: string, ...middleware: Middleware[]): void;

  /**
   * Mount static web content from given source directory.
   *
   * @param path The path where the source will be mounted on
   * @param docroot The source directory path
   */
  static(path: string, docroot: string): void;

  /**
   * Starts the server.
   *
   * @param port TCP port number, if 0 or missing then random unused port will be allocated
   * @param callback host name or IP address for listening on, default 127.0.0.1
   * @returns The instance for fluent/chaining API
   */
  listen(addr?: string, callback?: () => void): void;
}

/**
 * The `req` object represents the HTTP request and has properties for the request query string, parameters, body, HTTP headers, and so on.
 *
 * In this documentation and by convention, the object is always referred to as `req` (and the HTTP response is `res`) but its actual name is determined by the parameters to the callback function in which you’re working.
 *
 * @example
 * app.get("/user/:id", function (req, res) {
 *   res.send("user " + req.params.id);
 * });
 *
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
   * For example, if you have the route /user/:name, then the “name” property is available as req.params.name.
   * This object defaults to empty.
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
 * @example
 * app.get("/user/:id", function (req, res) {
 *   res.send("user " + req.params.id);
 * });
 */
export interface Response {
  /**
   * Appends the specified value to the HTTP response header field. If the header is not already set, it creates the header with the specified value.
   *
   * @param field the header field name
   * @param value the value to append
   */
  append: (field: string, value: string) => Response;

  /**
   * Sends a JSON response. This method sends a response (with the correct content-type) that is the parameter converted to a JSON string.
   *
   * @param body the object to send
   */
  json: (body: Record<string, any>) => Response;

  /**
   * Sends a plain text response. This method sends a response (with the correct content-type) that is the string formatting result.
   *
   * @param format go format string
   * @param v format values (if any)
   */
  text: (format: string, v?: any[]) => Response;

  /**
   * Sends a HTML text response. This method sends a response (with the correct content-type) that is the body string paramter.
   * @param body the string to send
   */
  html: (body: string) => Response;

  /**
   * Sends a binray response. This method sends a response (with the "application/octet-stream" content-type) that is the body paramter.
   *
   * @param body the data to send
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
   */
  send: (body: string | number[] | ArrayBuffer) => Response;

  /**
   * Sets the HTTP status for the response.
   *
   * @param code the satus code value
   */
  status: (code: number) => Response;

  /**
   * Sets the Content-Type HTTP header to the MIME type as from mime parameter.
   *
   * @params mime the content type
   */
  type: (mime: string) => Response;

  /**
   * Adds the header field to the Vary response header.
   *
   * @param header the header filed name
   */
  vary: (header: string) => Response;

  /**
   * Sets the response’s HTTP header field to value.
   *
   * @param field the header field name
   * @param value the value to set
   */
  set: (field: string, value: string) => Response;

  /**
   * Redirects to the URL, with specified status, a positive integer that corresponds to an HTTP status code.
   *
   * @param code the HTTP status code (301, 302, ...)
   * @param loc the location to redirect
   */
  redirect: (code: number, loc: string) => Response;
}
