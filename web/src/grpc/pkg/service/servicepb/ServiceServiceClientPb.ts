/**
 * @fileoverview gRPC-Web generated client stub for pkg.service.servicepb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as pkg_service_servicepb_service_pb from '../../../pkg/service/servicepb/service_pb';


export class ServiceServiceClient {
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
    pkg_service_servicepb_service_pb.GetAllResponse,
    (request: pkg_service_servicepb_service_pb.GetAllRequest) => {
      return request.serializeBinary();
    },
    pkg_service_servicepb_service_pb.GetAllResponse.deserializeBinary
  );

  getAll(
    request: pkg_service_servicepb_service_pb.GetAllRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_service_servicepb_service_pb.GetAllResponse>;

  getAll(
    request: pkg_service_servicepb_service_pb.GetAllRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_service_servicepb_service_pb.GetAllResponse) => void): grpcWeb.ClientReadableStream<pkg_service_servicepb_service_pb.GetAllResponse>;

  getAll(
    request: pkg_service_servicepb_service_pb.GetAllRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_service_servicepb_service_pb.GetAllResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.service.servicepb.ServiceService/GetAll',
        request,
        metadata || {},
        this.methodInfoGetAll,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.service.servicepb.ServiceService/GetAll',
    request,
    metadata || {},
    this.methodInfoGetAll);
  }

  methodInfoGetByID = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_service_servicepb_service_pb.GetByIDResponse,
    (request: pkg_service_servicepb_service_pb.GetByIDRequest) => {
      return request.serializeBinary();
    },
    pkg_service_servicepb_service_pb.GetByIDResponse.deserializeBinary
  );

  getByID(
    request: pkg_service_servicepb_service_pb.GetByIDRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_service_servicepb_service_pb.GetByIDResponse>;

  getByID(
    request: pkg_service_servicepb_service_pb.GetByIDRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_service_servicepb_service_pb.GetByIDResponse) => void): grpcWeb.ClientReadableStream<pkg_service_servicepb_service_pb.GetByIDResponse>;

  getByID(
    request: pkg_service_servicepb_service_pb.GetByIDRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_service_servicepb_service_pb.GetByIDResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.service.servicepb.ServiceService/GetByID',
        request,
        metadata || {},
        this.methodInfoGetByID,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.service.servicepb.ServiceService/GetByID',
    request,
    metadata || {},
    this.methodInfoGetByID);
  }

  methodInfoDelete = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_service_servicepb_service_pb.DeleteResponse,
    (request: pkg_service_servicepb_service_pb.DeleteRequest) => {
      return request.serializeBinary();
    },
    pkg_service_servicepb_service_pb.DeleteResponse.deserializeBinary
  );

  delete(
    request: pkg_service_servicepb_service_pb.DeleteRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_service_servicepb_service_pb.DeleteResponse>;

  delete(
    request: pkg_service_servicepb_service_pb.DeleteRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_service_servicepb_service_pb.DeleteResponse) => void): grpcWeb.ClientReadableStream<pkg_service_servicepb_service_pb.DeleteResponse>;

  delete(
    request: pkg_service_servicepb_service_pb.DeleteRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_service_servicepb_service_pb.DeleteResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.service.servicepb.ServiceService/Delete',
        request,
        metadata || {},
        this.methodInfoDelete,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.service.servicepb.ServiceService/Delete',
    request,
    metadata || {},
    this.methodInfoDelete);
  }

  methodInfoStore = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_service_servicepb_service_pb.StoreResponse,
    (request: pkg_service_servicepb_service_pb.StoreRequest) => {
      return request.serializeBinary();
    },
    pkg_service_servicepb_service_pb.StoreResponse.deserializeBinary
  );

  store(
    request: pkg_service_servicepb_service_pb.StoreRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_service_servicepb_service_pb.StoreResponse>;

  store(
    request: pkg_service_servicepb_service_pb.StoreRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_service_servicepb_service_pb.StoreResponse) => void): grpcWeb.ClientReadableStream<pkg_service_servicepb_service_pb.StoreResponse>;

  store(
    request: pkg_service_servicepb_service_pb.StoreRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_service_servicepb_service_pb.StoreResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.service.servicepb.ServiceService/Store',
        request,
        metadata || {},
        this.methodInfoStore,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.service.servicepb.ServiceService/Store',
    request,
    metadata || {},
    this.methodInfoStore);
  }

  methodInfoUpdate = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_service_servicepb_service_pb.UpdateResponse,
    (request: pkg_service_servicepb_service_pb.UpdateRequest) => {
      return request.serializeBinary();
    },
    pkg_service_servicepb_service_pb.UpdateResponse.deserializeBinary
  );

  update(
    request: pkg_service_servicepb_service_pb.UpdateRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_service_servicepb_service_pb.UpdateResponse>;

  update(
    request: pkg_service_servicepb_service_pb.UpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_service_servicepb_service_pb.UpdateResponse) => void): grpcWeb.ClientReadableStream<pkg_service_servicepb_service_pb.UpdateResponse>;

  update(
    request: pkg_service_servicepb_service_pb.UpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_service_servicepb_service_pb.UpdateResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.service.servicepb.ServiceService/Update',
        request,
        metadata || {},
        this.methodInfoUpdate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.service.servicepb.ServiceService/Update',
    request,
    metadata || {},
    this.methodInfoUpdate);
  }

  methodInfoTestService = new grpcWeb.AbstractClientBase.MethodInfo(
    pkg_service_servicepb_service_pb.TestServiceResponse,
    (request: pkg_service_servicepb_service_pb.TestServiceRequest) => {
      return request.serializeBinary();
    },
    pkg_service_servicepb_service_pb.TestServiceResponse.deserializeBinary
  );

  testService(
    request: pkg_service_servicepb_service_pb.TestServiceRequest,
    metadata: grpcWeb.Metadata | null): Promise<pkg_service_servicepb_service_pb.TestServiceResponse>;

  testService(
    request: pkg_service_servicepb_service_pb.TestServiceRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pkg_service_servicepb_service_pb.TestServiceResponse) => void): grpcWeb.ClientReadableStream<pkg_service_servicepb_service_pb.TestServiceResponse>;

  testService(
    request: pkg_service_servicepb_service_pb.TestServiceRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pkg_service_servicepb_service_pb.TestServiceResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pkg.service.servicepb.ServiceService/TestService',
        request,
        metadata || {},
        this.methodInfoTestService,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pkg.service.servicepb.ServiceService/TestService',
    request,
    metadata || {},
    this.methodInfoTestService);
  }

}

