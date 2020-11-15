/**
 * @fileoverview gRPC-Web generated client stub for pkg.host.hostpb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as pkg_host_hostpb_host_pb from '../../../pkg/host/hostpb/host_pb';


export class HostServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoGetAll = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_host_hostpb_host_pb.GetAllResponse,
    (request: pkg_host_hostpb_host_pb.GetAllRequest) => {
      return request.serializeBinary();
    },
    pkg_host_hostpb_host_pb.GetAllResponse.deserializeBinary
  );

  getAll(
    request: pkg_host_hostpb_host_pb.GetAllRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_host_hostpb_host_pb.GetAllResponse>;

  getAll(
    request: pkg_host_hostpb_host_pb.GetAllRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_host_hostpb_host_pb.GetAllResponse) => void): grpcWeb.ClientReadableStream<pkg_host_hostpb_host_pb.GetAllResponse>;

  getAll(
    request: pkg_host_hostpb_host_pb.GetAllRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_host_hostpb_host_pb.GetAllResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.host.hostpb.HostService/GetAll',
        request,
        metadata || {},
        this.methodInfoGetAll,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.host.hostpb.HostService/GetAll',
    request,
    metadata || {},
    this.methodInfoGetAll);
  }

  methodInfoGetByID = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_host_hostpb_host_pb.GetByIDResponse,
    (request: pkg_host_hostpb_host_pb.GetByIDRequest) => {
      return request.serializeBinary();
    },
    pkg_host_hostpb_host_pb.GetByIDResponse.deserializeBinary
  );

  getByID(
    request: pkg_host_hostpb_host_pb.GetByIDRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_host_hostpb_host_pb.GetByIDResponse>;

  getByID(
    request: pkg_host_hostpb_host_pb.GetByIDRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_host_hostpb_host_pb.GetByIDResponse) => void): grpcWeb.ClientReadableStream<pkg_host_hostpb_host_pb.GetByIDResponse>;

  getByID(
    request: pkg_host_hostpb_host_pb.GetByIDRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_host_hostpb_host_pb.GetByIDResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.host.hostpb.HostService/GetByID',
        request,
        metadata || {},
        this.methodInfoGetByID,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.host.hostpb.HostService/GetByID',
    request,
    metadata || {},
    this.methodInfoGetByID);
  }

  methodInfoDelete = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_host_hostpb_host_pb.DeleteResponse,
    (request: pkg_host_hostpb_host_pb.DeleteRequest) => {
      return request.serializeBinary();
    },
    pkg_host_hostpb_host_pb.DeleteResponse.deserializeBinary
  );

  delete(
    request: pkg_host_hostpb_host_pb.DeleteRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_host_hostpb_host_pb.DeleteResponse>;

  delete(
    request: pkg_host_hostpb_host_pb.DeleteRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_host_hostpb_host_pb.DeleteResponse) => void): grpcWeb.ClientReadableStream<pkg_host_hostpb_host_pb.DeleteResponse>;

  delete(
    request: pkg_host_hostpb_host_pb.DeleteRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_host_hostpb_host_pb.DeleteResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.host.hostpb.HostService/Delete',
        request,
        metadata || {},
        this.methodInfoDelete,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.host.hostpb.HostService/Delete',
    request,
    metadata || {},
    this.methodInfoDelete);
  }

  methodInfoStore = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_host_hostpb_host_pb.StoreResponse,
    (request: pkg_host_hostpb_host_pb.StoreRequest) => {
      return request.serializeBinary();
    },
    pkg_host_hostpb_host_pb.StoreResponse.deserializeBinary
  );

  store(
    request: pkg_host_hostpb_host_pb.StoreRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_host_hostpb_host_pb.StoreResponse>;

  store(
    request: pkg_host_hostpb_host_pb.StoreRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_host_hostpb_host_pb.StoreResponse) => void): grpcWeb.ClientReadableStream<pkg_host_hostpb_host_pb.StoreResponse>;

  store(
    request: pkg_host_hostpb_host_pb.StoreRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_host_hostpb_host_pb.StoreResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.host.hostpb.HostService/Store',
        request,
        metadata || {},
        this.methodInfoStore,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.host.hostpb.HostService/Store',
    request,
    metadata || {},
    this.methodInfoStore);
  }

  methodInfoUpdate = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_host_hostpb_host_pb.UpdateResponse,
    (request: pkg_host_hostpb_host_pb.UpdateRequest) => {
      return request.serializeBinary();
    },
    pkg_host_hostpb_host_pb.UpdateResponse.deserializeBinary
  );

  update(
    request: pkg_host_hostpb_host_pb.UpdateRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_host_hostpb_host_pb.UpdateResponse>;

  update(
    request: pkg_host_hostpb_host_pb.UpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_host_hostpb_host_pb.UpdateResponse) => void): grpcWeb.ClientReadableStream<pkg_host_hostpb_host_pb.UpdateResponse>;

  update(
    request: pkg_host_hostpb_host_pb.UpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_host_hostpb_host_pb.UpdateResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.host.hostpb.HostService/Update',
        request,
        metadata || {},
        this.methodInfoUpdate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.host.hostpb.HostService/Update',
    request,
    metadata || {},
    this.methodInfoUpdate);
  }

}
