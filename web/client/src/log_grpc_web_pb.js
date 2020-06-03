/**
 * @fileoverview gRPC-Web generated client stub for pb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.pb = require('./log_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.pb.logServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.pb.logServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pb.DHCPLogsRequest,
 *   !proto.pb.DHCPLogsResponse>}
 */
const methodDescriptor_logService_GetDHCPLogs = new grpc.web.MethodDescriptor(
  '/pb.logService/GetDHCPLogs',
  grpc.web.MethodType.UNARY,
  proto.pb.DHCPLogsRequest,
  proto.pb.DHCPLogsResponse,
  /**
   * @param {!proto.pb.DHCPLogsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pb.DHCPLogsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.pb.DHCPLogsRequest,
 *   !proto.pb.DHCPLogsResponse>}
 */
const methodInfo_logService_GetDHCPLogs = new grpc.web.AbstractClientBase.MethodInfo(
  proto.pb.DHCPLogsResponse,
  /**
   * @param {!proto.pb.DHCPLogsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pb.DHCPLogsResponse.deserializeBinary
);


/**
 * @param {!proto.pb.DHCPLogsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.pb.DHCPLogsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pb.DHCPLogsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pb.logServiceClient.prototype.getDHCPLogs =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pb.logService/GetDHCPLogs',
      request,
      metadata || {},
      methodDescriptor_logService_GetDHCPLogs,
      callback);
};


/**
 * @param {!proto.pb.DHCPLogsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pb.DHCPLogsResponse>}
 *     A native promise that resolves to the response
 */
proto.pb.logServicePromiseClient.prototype.getDHCPLogs =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pb.logService/GetDHCPLogs',
      request,
      metadata || {},
      methodDescriptor_logService_GetDHCPLogs);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pb.SwitchLogsRequest,
 *   !proto.pb.SwitchLogsResponse>}
 */
const methodDescriptor_logService_GetSwitchLogs = new grpc.web.MethodDescriptor(
  '/pb.logService/GetSwitchLogs',
  grpc.web.MethodType.UNARY,
  proto.pb.SwitchLogsRequest,
  proto.pb.SwitchLogsResponse,
  /**
   * @param {!proto.pb.SwitchLogsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pb.SwitchLogsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.pb.SwitchLogsRequest,
 *   !proto.pb.SwitchLogsResponse>}
 */
const methodInfo_logService_GetSwitchLogs = new grpc.web.AbstractClientBase.MethodInfo(
  proto.pb.SwitchLogsResponse,
  /**
   * @param {!proto.pb.SwitchLogsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pb.SwitchLogsResponse.deserializeBinary
);


/**
 * @param {!proto.pb.SwitchLogsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.pb.SwitchLogsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pb.SwitchLogsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pb.logServiceClient.prototype.getSwitchLogs =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pb.logService/GetSwitchLogs',
      request,
      metadata || {},
      methodDescriptor_logService_GetSwitchLogs,
      callback);
};


/**
 * @param {!proto.pb.SwitchLogsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pb.SwitchLogsResponse>}
 *     A native promise that resolves to the response
 */
proto.pb.logServicePromiseClient.prototype.getSwitchLogs =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pb.logService/GetSwitchLogs',
      request,
      metadata || {},
      methodDescriptor_logService_GetSwitchLogs);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pb.SimilarSwitchesRequest,
 *   !proto.pb.SimilarSwitchesResponse>}
 */
const methodDescriptor_logService_GetSimilarSwitches = new grpc.web.MethodDescriptor(
  '/pb.logService/GetSimilarSwitches',
  grpc.web.MethodType.UNARY,
  proto.pb.SimilarSwitchesRequest,
  proto.pb.SimilarSwitchesResponse,
  /**
   * @param {!proto.pb.SimilarSwitchesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pb.SimilarSwitchesResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.pb.SimilarSwitchesRequest,
 *   !proto.pb.SimilarSwitchesResponse>}
 */
const methodInfo_logService_GetSimilarSwitches = new grpc.web.AbstractClientBase.MethodInfo(
  proto.pb.SimilarSwitchesResponse,
  /**
   * @param {!proto.pb.SimilarSwitchesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pb.SimilarSwitchesResponse.deserializeBinary
);


/**
 * @param {!proto.pb.SimilarSwitchesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.pb.SimilarSwitchesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pb.SimilarSwitchesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pb.logServiceClient.prototype.getSimilarSwitches =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pb.logService/GetSimilarSwitches',
      request,
      metadata || {},
      methodDescriptor_logService_GetSimilarSwitches,
      callback);
};


/**
 * @param {!proto.pb.SimilarSwitchesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pb.SimilarSwitchesResponse>}
 *     A native promise that resolves to the response
 */
proto.pb.logServicePromiseClient.prototype.getSimilarSwitches =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pb.logService/GetSimilarSwitches',
      request,
      metadata || {},
      methodDescriptor_logService_GetSimilarSwitches);
};


module.exports = proto.pb;

