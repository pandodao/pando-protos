// THIS IS AN AUTOGENERATED FILE. DO NOT EDIT THIS FILE DIRECTLY.
// Source: explorer/v1/explorer.proto
/* eslint-disable */

import type { ByteSource, ClientConfiguration } from "twirpscript";
import {
  BinaryReader,
  BinaryWriter,
  JSONrequest,
  PBrequest,
} from "twirpscript";
// This is the minimum version supported by the current runtime.
// If this line fails typechecking, breaking changes have been introduced and this
// file needs to be regenerated by running `yarn twirpscript`.
export { MIN_SUPPORTED_VERSION_0_0_56 } from "twirpscript";

//========================================//
//    ExplorerService Protobuf Client     //
//========================================//

/**
 * curl -X POST -d '{"limit":5,"cursor":"0"}' -H "Content-Type: application/json" $HOST/twirp/explorer.v1.ExplorerService/ListOutputs
 */
export async function ListOutputs(
  listOutputsRequest: ListOutputsRequest,
  config?: ClientConfiguration
): Promise<ListOutputsResponse> {
  const response = await PBrequest(
    "/explorer.v1.ExplorerService/ListOutputs",
    ListOutputsRequest.encode(listOutputsRequest),
    config
  );
  return ListOutputsResponse.decode(response);
}

/**
 * curl -X POST -d '{"transaction_hash":"f31b695293651f065149166b83a9100456c2f03a941ca2156f19c830022c81f6","output_index":0}' -H "Content-Type: application/json" $HOST/twirp/explorer.v1.ExplorerService/GetTransaction
 */
export async function GetTransaction(
  getTransactionRequest: GetTransactionRequest,
  config?: ClientConfiguration
): Promise<GetTransactionResponse> {
  const response = await PBrequest(
    "/explorer.v1.ExplorerService/GetTransaction",
    GetTransactionRequest.encode(getTransactionRequest),
    config
  );
  return GetTransactionResponse.decode(response);
}

/**
 * curl -X POST -d '{"trace_id":"9a11cfec-0761-3403-be66-4168116dbcac"}' -H "Content-Type: application/json" $HOST/twirp/explorer.v1.ExplorerService/GetTransactionByTraceID
 */
export async function GetTransactionByTraceID(
  getTransactionByTraceIDRequest: GetTransactionByTraceIDRequest,
  config?: ClientConfiguration
): Promise<GetTransactionByTraceIDResponse> {
  const response = await PBrequest(
    "/explorer.v1.ExplorerService/GetTransactionByTraceID",
    GetTransactionByTraceIDRequest.encode(getTransactionByTraceIDRequest),
    config
  );
  return GetTransactionByTraceIDResponse.decode(response);
}

//========================================//
//      ExplorerService JSON Client       //
//========================================//

/**
 * curl -X POST -d '{"limit":5,"cursor":"0"}' -H "Content-Type: application/json" $HOST/twirp/explorer.v1.ExplorerService/ListOutputs
 */
export async function ListOutputsJSON(
  listOutputsRequest: ListOutputsRequest,
  config?: ClientConfiguration
): Promise<ListOutputsResponse> {
  const response = await JSONrequest(
    "/explorer.v1.ExplorerService/ListOutputs",
    ListOutputsRequestJSON.encode(listOutputsRequest),
    config
  );
  return ListOutputsResponseJSON.decode(response);
}

/**
 * curl -X POST -d '{"transaction_hash":"f31b695293651f065149166b83a9100456c2f03a941ca2156f19c830022c81f6","output_index":0}' -H "Content-Type: application/json" $HOST/twirp/explorer.v1.ExplorerService/GetTransaction
 */
export async function GetTransactionJSON(
  getTransactionRequest: GetTransactionRequest,
  config?: ClientConfiguration
): Promise<GetTransactionResponse> {
  const response = await JSONrequest(
    "/explorer.v1.ExplorerService/GetTransaction",
    GetTransactionRequestJSON.encode(getTransactionRequest),
    config
  );
  return GetTransactionResponseJSON.decode(response);
}

/**
 * curl -X POST -d '{"trace_id":"9a11cfec-0761-3403-be66-4168116dbcac"}' -H "Content-Type: application/json" $HOST/twirp/explorer.v1.ExplorerService/GetTransactionByTraceID
 */
export async function GetTransactionByTraceIDJSON(
  getTransactionByTraceIDRequest: GetTransactionByTraceIDRequest,
  config?: ClientConfiguration
): Promise<GetTransactionByTraceIDResponse> {
  const response = await JSONrequest(
    "/explorer.v1.ExplorerService/GetTransactionByTraceID",
    GetTransactionByTraceIDRequestJSON.encode(getTransactionByTraceIDRequest),
    config
  );
  return GetTransactionByTraceIDResponseJSON.decode(response);
}

//========================================//
//            ExplorerService             //
//========================================//

export interface ExplorerService<Context = unknown> {
  /**
   * curl -X POST -d '{"limit":5,"cursor":"0"}' -H "Content-Type: application/json" $HOST/twirp/explorer.v1.ExplorerService/ListOutputs
   */
  ListOutputs: (
    listOutputsRequest: ListOutputsRequest,
    context: Context
  ) => Promise<ListOutputsResponse> | ListOutputsResponse;
  /**
   * curl -X POST -d '{"transaction_hash":"f31b695293651f065149166b83a9100456c2f03a941ca2156f19c830022c81f6","output_index":0}' -H "Content-Type: application/json" $HOST/twirp/explorer.v1.ExplorerService/GetTransaction
   */
  GetTransaction: (
    getTransactionRequest: GetTransactionRequest,
    context: Context
  ) => Promise<GetTransactionResponse> | GetTransactionResponse;
  /**
   * curl -X POST -d '{"trace_id":"9a11cfec-0761-3403-be66-4168116dbcac"}' -H "Content-Type: application/json" $HOST/twirp/explorer.v1.ExplorerService/GetTransactionByTraceID
   */
  GetTransactionByTraceID: (
    getTransactionByTraceIDRequest: GetTransactionByTraceIDRequest,
    context: Context
  ) =>
    | Promise<GetTransactionByTraceIDResponse>
    | GetTransactionByTraceIDResponse;
}

export function createExplorerService<Context>(
  service: ExplorerService<Context>
) {
  return {
    name: "explorer.v1.ExplorerService",
    methods: {
      ListOutputs: {
        name: "ListOutputs",
        handler: service.ListOutputs,
        input: { protobuf: ListOutputsRequest, json: ListOutputsRequestJSON },
        output: {
          protobuf: ListOutputsResponse,
          json: ListOutputsResponseJSON,
        },
      },
      GetTransaction: {
        name: "GetTransaction",
        handler: service.GetTransaction,
        input: {
          protobuf: GetTransactionRequest,
          json: GetTransactionRequestJSON,
        },
        output: {
          protobuf: GetTransactionResponse,
          json: GetTransactionResponseJSON,
        },
      },
      GetTransactionByTraceID: {
        name: "GetTransactionByTraceID",
        handler: service.GetTransactionByTraceID,
        input: {
          protobuf: GetTransactionByTraceIDRequest,
          json: GetTransactionByTraceIDRequestJSON,
        },
        output: {
          protobuf: GetTransactionByTraceIDResponse,
          json: GetTransactionByTraceIDResponseJSON,
        },
      },
    },
  } as const;
}

//========================================//
//                 Types                  //
//========================================//

export interface Output {
  id: string;
  createdAt: number;
  assetId: string;
  amount: string;
  memo: string;
  transactionHash: string;
  outputIndex: number;
  headerAction: number;
  protocolId: number;
}

export interface Transfer {
  id: string;
  createdAt: number;
  assetId: string;
  amount: string;
  threshold: number;
  signedBy: string;
}

export interface ListOutputsRequest {
  cursor: string;
  limit: number;
}

export interface ListOutputsResponse {
  outputs: Output[];
  nextCursor: string;
}

export interface GetTransactionRequest {
  transactionHash: string;
  outputIndex: number;
}

export interface GetTransactionResponse {
  outputs: Output[];
  transfers: Transfer[];
}

export interface GetTransactionByTraceIDRequest {
  traceId: string;
}

export interface GetTransactionByTraceIDResponse {
  outputs: Output[];
  transfers: Transfer[];
}

//========================================//
//        Protobuf Encode / Decode        //
//========================================//

export const Output = {
  /**
   * Serializes Output to protobuf.
   */
  encode: function (msg: Partial<Output>): Uint8Array {
    return Output._writeMessage(msg, new BinaryWriter()).getResultBuffer();
  },

  /**
   * Deserializes Output from protobuf.
   */
  decode: function (bytes: ByteSource): Output {
    return Output._readMessage(Output.initialize(), new BinaryReader(bytes));
  },

  /**
   * Initializes Output with all fields set to their default value.
   */
  initialize: function (): Output {
    return {
      id: "",
      createdAt: 0,
      assetId: "",
      amount: "",
      memo: "",
      transactionHash: "",
      outputIndex: 0,
      headerAction: 0,
      protocolId: 0,
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: Partial<Output>,
    writer: BinaryWriter
  ): BinaryWriter {
    if (msg.id) {
      writer.writeString(1, msg.id);
    }
    if (msg.createdAt) {
      writer.writeInt32(2, msg.createdAt);
    }
    if (msg.assetId) {
      writer.writeString(3, msg.assetId);
    }
    if (msg.amount) {
      writer.writeString(4, msg.amount);
    }
    if (msg.memo) {
      writer.writeString(5, msg.memo);
    }
    if (msg.transactionHash) {
      writer.writeString(6, msg.transactionHash);
    }
    if (msg.outputIndex) {
      writer.writeInt32(7, msg.outputIndex);
    }
    if (msg.headerAction) {
      writer.writeInt32(8, msg.headerAction);
    }
    if (msg.protocolId) {
      writer.writeInt32(9, msg.protocolId);
    }
    return writer;
  },

  /**
   * @private
   */
  _readMessage: function (msg: Output, reader: BinaryReader): Output {
    while (reader.nextField()) {
      const field = reader.getFieldNumber();
      switch (field) {
        case 1: {
          msg.id = reader.readString();
          break;
        }
        case 2: {
          msg.createdAt = reader.readInt32();
          break;
        }
        case 3: {
          msg.assetId = reader.readString();
          break;
        }
        case 4: {
          msg.amount = reader.readString();
          break;
        }
        case 5: {
          msg.memo = reader.readString();
          break;
        }
        case 6: {
          msg.transactionHash = reader.readString();
          break;
        }
        case 7: {
          msg.outputIndex = reader.readInt32();
          break;
        }
        case 8: {
          msg.headerAction = reader.readInt32();
          break;
        }
        case 9: {
          msg.protocolId = reader.readInt32();
          break;
        }
        default: {
          reader.skipField();
          break;
        }
      }
    }
    return msg;
  },
};

export const Transfer = {
  /**
   * Serializes Transfer to protobuf.
   */
  encode: function (msg: Partial<Transfer>): Uint8Array {
    return Transfer._writeMessage(msg, new BinaryWriter()).getResultBuffer();
  },

  /**
   * Deserializes Transfer from protobuf.
   */
  decode: function (bytes: ByteSource): Transfer {
    return Transfer._readMessage(
      Transfer.initialize(),
      new BinaryReader(bytes)
    );
  },

  /**
   * Initializes Transfer with all fields set to their default value.
   */
  initialize: function (): Transfer {
    return {
      id: "",
      createdAt: 0,
      assetId: "",
      amount: "",
      threshold: 0,
      signedBy: "",
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: Partial<Transfer>,
    writer: BinaryWriter
  ): BinaryWriter {
    if (msg.id) {
      writer.writeString(1, msg.id);
    }
    if (msg.createdAt) {
      writer.writeInt32(2, msg.createdAt);
    }
    if (msg.assetId) {
      writer.writeString(3, msg.assetId);
    }
    if (msg.amount) {
      writer.writeString(4, msg.amount);
    }
    if (msg.threshold) {
      writer.writeInt32(5, msg.threshold);
    }
    if (msg.signedBy) {
      writer.writeString(6, msg.signedBy);
    }
    return writer;
  },

  /**
   * @private
   */
  _readMessage: function (msg: Transfer, reader: BinaryReader): Transfer {
    while (reader.nextField()) {
      const field = reader.getFieldNumber();
      switch (field) {
        case 1: {
          msg.id = reader.readString();
          break;
        }
        case 2: {
          msg.createdAt = reader.readInt32();
          break;
        }
        case 3: {
          msg.assetId = reader.readString();
          break;
        }
        case 4: {
          msg.amount = reader.readString();
          break;
        }
        case 5: {
          msg.threshold = reader.readInt32();
          break;
        }
        case 6: {
          msg.signedBy = reader.readString();
          break;
        }
        default: {
          reader.skipField();
          break;
        }
      }
    }
    return msg;
  },
};

export const ListOutputsRequest = {
  /**
   * Serializes ListOutputsRequest to protobuf.
   */
  encode: function (msg: Partial<ListOutputsRequest>): Uint8Array {
    return ListOutputsRequest._writeMessage(
      msg,
      new BinaryWriter()
    ).getResultBuffer();
  },

  /**
   * Deserializes ListOutputsRequest from protobuf.
   */
  decode: function (bytes: ByteSource): ListOutputsRequest {
    return ListOutputsRequest._readMessage(
      ListOutputsRequest.initialize(),
      new BinaryReader(bytes)
    );
  },

  /**
   * Initializes ListOutputsRequest with all fields set to their default value.
   */
  initialize: function (): ListOutputsRequest {
    return {
      cursor: "",
      limit: 0,
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: Partial<ListOutputsRequest>,
    writer: BinaryWriter
  ): BinaryWriter {
    if (msg.cursor) {
      writer.writeString(1, msg.cursor);
    }
    if (msg.limit) {
      writer.writeInt32(2, msg.limit);
    }
    return writer;
  },

  /**
   * @private
   */
  _readMessage: function (
    msg: ListOutputsRequest,
    reader: BinaryReader
  ): ListOutputsRequest {
    while (reader.nextField()) {
      const field = reader.getFieldNumber();
      switch (field) {
        case 1: {
          msg.cursor = reader.readString();
          break;
        }
        case 2: {
          msg.limit = reader.readInt32();
          break;
        }
        default: {
          reader.skipField();
          break;
        }
      }
    }
    return msg;
  },
};

export const ListOutputsResponse = {
  /**
   * Serializes ListOutputsResponse to protobuf.
   */
  encode: function (msg: Partial<ListOutputsResponse>): Uint8Array {
    return ListOutputsResponse._writeMessage(
      msg,
      new BinaryWriter()
    ).getResultBuffer();
  },

  /**
   * Deserializes ListOutputsResponse from protobuf.
   */
  decode: function (bytes: ByteSource): ListOutputsResponse {
    return ListOutputsResponse._readMessage(
      ListOutputsResponse.initialize(),
      new BinaryReader(bytes)
    );
  },

  /**
   * Initializes ListOutputsResponse with all fields set to their default value.
   */
  initialize: function (): ListOutputsResponse {
    return {
      outputs: [],
      nextCursor: "",
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: Partial<ListOutputsResponse>,
    writer: BinaryWriter
  ): BinaryWriter {
    if (msg.outputs?.length) {
      writer.writeRepeatedMessage(1, msg.outputs as any, Output._writeMessage);
    }
    if (msg.nextCursor) {
      writer.writeString(2, msg.nextCursor);
    }
    return writer;
  },

  /**
   * @private
   */
  _readMessage: function (
    msg: ListOutputsResponse,
    reader: BinaryReader
  ): ListOutputsResponse {
    while (reader.nextField()) {
      const field = reader.getFieldNumber();
      switch (field) {
        case 1: {
          const m = Output.initialize();
          reader.readMessage(m, Output._readMessage);
          msg.outputs.push(m);
          break;
        }
        case 2: {
          msg.nextCursor = reader.readString();
          break;
        }
        default: {
          reader.skipField();
          break;
        }
      }
    }
    return msg;
  },
};

export const GetTransactionRequest = {
  /**
   * Serializes GetTransactionRequest to protobuf.
   */
  encode: function (msg: Partial<GetTransactionRequest>): Uint8Array {
    return GetTransactionRequest._writeMessage(
      msg,
      new BinaryWriter()
    ).getResultBuffer();
  },

  /**
   * Deserializes GetTransactionRequest from protobuf.
   */
  decode: function (bytes: ByteSource): GetTransactionRequest {
    return GetTransactionRequest._readMessage(
      GetTransactionRequest.initialize(),
      new BinaryReader(bytes)
    );
  },

  /**
   * Initializes GetTransactionRequest with all fields set to their default value.
   */
  initialize: function (): GetTransactionRequest {
    return {
      transactionHash: "",
      outputIndex: 0,
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: Partial<GetTransactionRequest>,
    writer: BinaryWriter
  ): BinaryWriter {
    if (msg.transactionHash) {
      writer.writeString(1, msg.transactionHash);
    }
    if (msg.outputIndex) {
      writer.writeInt32(2, msg.outputIndex);
    }
    return writer;
  },

  /**
   * @private
   */
  _readMessage: function (
    msg: GetTransactionRequest,
    reader: BinaryReader
  ): GetTransactionRequest {
    while (reader.nextField()) {
      const field = reader.getFieldNumber();
      switch (field) {
        case 1: {
          msg.transactionHash = reader.readString();
          break;
        }
        case 2: {
          msg.outputIndex = reader.readInt32();
          break;
        }
        default: {
          reader.skipField();
          break;
        }
      }
    }
    return msg;
  },
};

export const GetTransactionResponse = {
  /**
   * Serializes GetTransactionResponse to protobuf.
   */
  encode: function (msg: Partial<GetTransactionResponse>): Uint8Array {
    return GetTransactionResponse._writeMessage(
      msg,
      new BinaryWriter()
    ).getResultBuffer();
  },

  /**
   * Deserializes GetTransactionResponse from protobuf.
   */
  decode: function (bytes: ByteSource): GetTransactionResponse {
    return GetTransactionResponse._readMessage(
      GetTransactionResponse.initialize(),
      new BinaryReader(bytes)
    );
  },

  /**
   * Initializes GetTransactionResponse with all fields set to their default value.
   */
  initialize: function (): GetTransactionResponse {
    return {
      outputs: [],
      transfers: [],
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: Partial<GetTransactionResponse>,
    writer: BinaryWriter
  ): BinaryWriter {
    if (msg.outputs?.length) {
      writer.writeRepeatedMessage(1, msg.outputs as any, Output._writeMessage);
    }
    if (msg.transfers?.length) {
      writer.writeRepeatedMessage(
        2,
        msg.transfers as any,
        Transfer._writeMessage
      );
    }
    return writer;
  },

  /**
   * @private
   */
  _readMessage: function (
    msg: GetTransactionResponse,
    reader: BinaryReader
  ): GetTransactionResponse {
    while (reader.nextField()) {
      const field = reader.getFieldNumber();
      switch (field) {
        case 1: {
          const m = Output.initialize();
          reader.readMessage(m, Output._readMessage);
          msg.outputs.push(m);
          break;
        }
        case 2: {
          const m = Transfer.initialize();
          reader.readMessage(m, Transfer._readMessage);
          msg.transfers.push(m);
          break;
        }
        default: {
          reader.skipField();
          break;
        }
      }
    }
    return msg;
  },
};

export const GetTransactionByTraceIDRequest = {
  /**
   * Serializes GetTransactionByTraceIDRequest to protobuf.
   */
  encode: function (msg: Partial<GetTransactionByTraceIDRequest>): Uint8Array {
    return GetTransactionByTraceIDRequest._writeMessage(
      msg,
      new BinaryWriter()
    ).getResultBuffer();
  },

  /**
   * Deserializes GetTransactionByTraceIDRequest from protobuf.
   */
  decode: function (bytes: ByteSource): GetTransactionByTraceIDRequest {
    return GetTransactionByTraceIDRequest._readMessage(
      GetTransactionByTraceIDRequest.initialize(),
      new BinaryReader(bytes)
    );
  },

  /**
   * Initializes GetTransactionByTraceIDRequest with all fields set to their default value.
   */
  initialize: function (): GetTransactionByTraceIDRequest {
    return {
      traceId: "",
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: Partial<GetTransactionByTraceIDRequest>,
    writer: BinaryWriter
  ): BinaryWriter {
    if (msg.traceId) {
      writer.writeString(1, msg.traceId);
    }
    return writer;
  },

  /**
   * @private
   */
  _readMessage: function (
    msg: GetTransactionByTraceIDRequest,
    reader: BinaryReader
  ): GetTransactionByTraceIDRequest {
    while (reader.nextField()) {
      const field = reader.getFieldNumber();
      switch (field) {
        case 1: {
          msg.traceId = reader.readString();
          break;
        }
        default: {
          reader.skipField();
          break;
        }
      }
    }
    return msg;
  },
};

export const GetTransactionByTraceIDResponse = {
  /**
   * Serializes GetTransactionByTraceIDResponse to protobuf.
   */
  encode: function (msg: Partial<GetTransactionByTraceIDResponse>): Uint8Array {
    return GetTransactionByTraceIDResponse._writeMessage(
      msg,
      new BinaryWriter()
    ).getResultBuffer();
  },

  /**
   * Deserializes GetTransactionByTraceIDResponse from protobuf.
   */
  decode: function (bytes: ByteSource): GetTransactionByTraceIDResponse {
    return GetTransactionByTraceIDResponse._readMessage(
      GetTransactionByTraceIDResponse.initialize(),
      new BinaryReader(bytes)
    );
  },

  /**
   * Initializes GetTransactionByTraceIDResponse with all fields set to their default value.
   */
  initialize: function (): GetTransactionByTraceIDResponse {
    return {
      outputs: [],
      transfers: [],
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: Partial<GetTransactionByTraceIDResponse>,
    writer: BinaryWriter
  ): BinaryWriter {
    if (msg.outputs?.length) {
      writer.writeRepeatedMessage(1, msg.outputs as any, Output._writeMessage);
    }
    if (msg.transfers?.length) {
      writer.writeRepeatedMessage(
        2,
        msg.transfers as any,
        Transfer._writeMessage
      );
    }
    return writer;
  },

  /**
   * @private
   */
  _readMessage: function (
    msg: GetTransactionByTraceIDResponse,
    reader: BinaryReader
  ): GetTransactionByTraceIDResponse {
    while (reader.nextField()) {
      const field = reader.getFieldNumber();
      switch (field) {
        case 1: {
          const m = Output.initialize();
          reader.readMessage(m, Output._readMessage);
          msg.outputs.push(m);
          break;
        }
        case 2: {
          const m = Transfer.initialize();
          reader.readMessage(m, Transfer._readMessage);
          msg.transfers.push(m);
          break;
        }
        default: {
          reader.skipField();
          break;
        }
      }
    }
    return msg;
  },
};

//========================================//
//          JSON Encode / Decode          //
//========================================//

export const OutputJSON = {
  /**
   * Serializes Output to JSON.
   */
  encode: function (msg: Partial<Output>): string {
    return JSON.stringify(OutputJSON._writeMessage(msg));
  },

  /**
   * Deserializes Output from JSON.
   */
  decode: function (json: string): Output {
    return OutputJSON._readMessage(OutputJSON.initialize(), JSON.parse(json));
  },

  /**
   * Initializes Output with all fields set to their default value.
   */
  initialize: function (): Output {
    return {
      id: "",
      createdAt: 0,
      assetId: "",
      amount: "",
      memo: "",
      transactionHash: "",
      outputIndex: 0,
      headerAction: 0,
      protocolId: 0,
    };
  },

  /**
   * @private
   */
  _writeMessage: function (msg: Partial<Output>): Record<string, unknown> {
    const json: Record<string, unknown> = {};
    if (msg.id) {
      json.id = msg.id;
    }
    if (msg.createdAt) {
      json.createdAt = msg.createdAt;
    }
    if (msg.assetId) {
      json.assetId = msg.assetId;
    }
    if (msg.amount) {
      json.amount = msg.amount;
    }
    if (msg.memo) {
      json.memo = msg.memo;
    }
    if (msg.transactionHash) {
      json.transactionHash = msg.transactionHash;
    }
    if (msg.outputIndex) {
      json.outputIndex = msg.outputIndex;
    }
    if (msg.headerAction) {
      json.headerAction = msg.headerAction;
    }
    if (msg.protocolId) {
      json.protocolId = msg.protocolId;
    }
    return json;
  },

  /**
   * @private
   */
  _readMessage: function (msg: Output, json: any): Output {
    const _id = json.id;
    if (_id) {
      msg.id = _id;
    }
    const _createdAt = json.createdAt ?? json.created_at;
    if (_createdAt) {
      msg.createdAt = _createdAt;
    }
    const _assetId = json.assetId ?? json.asset_id;
    if (_assetId) {
      msg.assetId = _assetId;
    }
    const _amount = json.amount;
    if (_amount) {
      msg.amount = _amount;
    }
    const _memo = json.memo;
    if (_memo) {
      msg.memo = _memo;
    }
    const _transactionHash = json.transactionHash ?? json.transaction_hash;
    if (_transactionHash) {
      msg.transactionHash = _transactionHash;
    }
    const _outputIndex = json.outputIndex ?? json.output_index;
    if (_outputIndex) {
      msg.outputIndex = _outputIndex;
    }
    const _headerAction = json.headerAction ?? json.header_action;
    if (_headerAction) {
      msg.headerAction = _headerAction;
    }
    const _protocolId = json.protocolId ?? json.protocol_id;
    if (_protocolId) {
      msg.protocolId = _protocolId;
    }
    return msg;
  },
};

export const TransferJSON = {
  /**
   * Serializes Transfer to JSON.
   */
  encode: function (msg: Partial<Transfer>): string {
    return JSON.stringify(TransferJSON._writeMessage(msg));
  },

  /**
   * Deserializes Transfer from JSON.
   */
  decode: function (json: string): Transfer {
    return TransferJSON._readMessage(
      TransferJSON.initialize(),
      JSON.parse(json)
    );
  },

  /**
   * Initializes Transfer with all fields set to their default value.
   */
  initialize: function (): Transfer {
    return {
      id: "",
      createdAt: 0,
      assetId: "",
      amount: "",
      threshold: 0,
      signedBy: "",
    };
  },

  /**
   * @private
   */
  _writeMessage: function (msg: Partial<Transfer>): Record<string, unknown> {
    const json: Record<string, unknown> = {};
    if (msg.id) {
      json.id = msg.id;
    }
    if (msg.createdAt) {
      json.createdAt = msg.createdAt;
    }
    if (msg.assetId) {
      json.assetId = msg.assetId;
    }
    if (msg.amount) {
      json.amount = msg.amount;
    }
    if (msg.threshold) {
      json.threshold = msg.threshold;
    }
    if (msg.signedBy) {
      json.signedBy = msg.signedBy;
    }
    return json;
  },

  /**
   * @private
   */
  _readMessage: function (msg: Transfer, json: any): Transfer {
    const _id = json.id;
    if (_id) {
      msg.id = _id;
    }
    const _createdAt = json.createdAt ?? json.created_at;
    if (_createdAt) {
      msg.createdAt = _createdAt;
    }
    const _assetId = json.assetId ?? json.asset_id;
    if (_assetId) {
      msg.assetId = _assetId;
    }
    const _amount = json.amount;
    if (_amount) {
      msg.amount = _amount;
    }
    const _threshold = json.threshold;
    if (_threshold) {
      msg.threshold = _threshold;
    }
    const _signedBy = json.signedBy ?? json.signed_by;
    if (_signedBy) {
      msg.signedBy = _signedBy;
    }
    return msg;
  },
};

export const ListOutputsRequestJSON = {
  /**
   * Serializes ListOutputsRequest to JSON.
   */
  encode: function (msg: Partial<ListOutputsRequest>): string {
    return JSON.stringify(ListOutputsRequestJSON._writeMessage(msg));
  },

  /**
   * Deserializes ListOutputsRequest from JSON.
   */
  decode: function (json: string): ListOutputsRequest {
    return ListOutputsRequestJSON._readMessage(
      ListOutputsRequestJSON.initialize(),
      JSON.parse(json)
    );
  },

  /**
   * Initializes ListOutputsRequest with all fields set to their default value.
   */
  initialize: function (): ListOutputsRequest {
    return {
      cursor: "",
      limit: 0,
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: Partial<ListOutputsRequest>
  ): Record<string, unknown> {
    const json: Record<string, unknown> = {};
    if (msg.cursor) {
      json.cursor = msg.cursor;
    }
    if (msg.limit) {
      json.limit = msg.limit;
    }
    return json;
  },

  /**
   * @private
   */
  _readMessage: function (
    msg: ListOutputsRequest,
    json: any
  ): ListOutputsRequest {
    const _cursor = json.cursor;
    if (_cursor) {
      msg.cursor = _cursor;
    }
    const _limit = json.limit;
    if (_limit) {
      msg.limit = _limit;
    }
    return msg;
  },
};

export const ListOutputsResponseJSON = {
  /**
   * Serializes ListOutputsResponse to JSON.
   */
  encode: function (msg: Partial<ListOutputsResponse>): string {
    return JSON.stringify(ListOutputsResponseJSON._writeMessage(msg));
  },

  /**
   * Deserializes ListOutputsResponse from JSON.
   */
  decode: function (json: string): ListOutputsResponse {
    return ListOutputsResponseJSON._readMessage(
      ListOutputsResponseJSON.initialize(),
      JSON.parse(json)
    );
  },

  /**
   * Initializes ListOutputsResponse with all fields set to their default value.
   */
  initialize: function (): ListOutputsResponse {
    return {
      outputs: [],
      nextCursor: "",
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: Partial<ListOutputsResponse>
  ): Record<string, unknown> {
    const json: Record<string, unknown> = {};
    if (msg.outputs?.length) {
      json.outputs = msg.outputs.map(OutputJSON._writeMessage);
    }
    if (msg.nextCursor) {
      json.nextCursor = msg.nextCursor;
    }
    return json;
  },

  /**
   * @private
   */
  _readMessage: function (
    msg: ListOutputsResponse,
    json: any
  ): ListOutputsResponse {
    const _outputs = json.outputs;
    if (_outputs) {
      for (const item of _outputs) {
        const m = Output.initialize();
        OutputJSON._readMessage(m, item);
        msg.outputs.push(m);
      }
    }
    const _nextCursor = json.nextCursor ?? json.next_cursor;
    if (_nextCursor) {
      msg.nextCursor = _nextCursor;
    }
    return msg;
  },
};

export const GetTransactionRequestJSON = {
  /**
   * Serializes GetTransactionRequest to JSON.
   */
  encode: function (msg: Partial<GetTransactionRequest>): string {
    return JSON.stringify(GetTransactionRequestJSON._writeMessage(msg));
  },

  /**
   * Deserializes GetTransactionRequest from JSON.
   */
  decode: function (json: string): GetTransactionRequest {
    return GetTransactionRequestJSON._readMessage(
      GetTransactionRequestJSON.initialize(),
      JSON.parse(json)
    );
  },

  /**
   * Initializes GetTransactionRequest with all fields set to their default value.
   */
  initialize: function (): GetTransactionRequest {
    return {
      transactionHash: "",
      outputIndex: 0,
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: Partial<GetTransactionRequest>
  ): Record<string, unknown> {
    const json: Record<string, unknown> = {};
    if (msg.transactionHash) {
      json.transactionHash = msg.transactionHash;
    }
    if (msg.outputIndex) {
      json.outputIndex = msg.outputIndex;
    }
    return json;
  },

  /**
   * @private
   */
  _readMessage: function (
    msg: GetTransactionRequest,
    json: any
  ): GetTransactionRequest {
    const _transactionHash = json.transactionHash ?? json.transaction_hash;
    if (_transactionHash) {
      msg.transactionHash = _transactionHash;
    }
    const _outputIndex = json.outputIndex ?? json.output_index;
    if (_outputIndex) {
      msg.outputIndex = _outputIndex;
    }
    return msg;
  },
};

export const GetTransactionResponseJSON = {
  /**
   * Serializes GetTransactionResponse to JSON.
   */
  encode: function (msg: Partial<GetTransactionResponse>): string {
    return JSON.stringify(GetTransactionResponseJSON._writeMessage(msg));
  },

  /**
   * Deserializes GetTransactionResponse from JSON.
   */
  decode: function (json: string): GetTransactionResponse {
    return GetTransactionResponseJSON._readMessage(
      GetTransactionResponseJSON.initialize(),
      JSON.parse(json)
    );
  },

  /**
   * Initializes GetTransactionResponse with all fields set to their default value.
   */
  initialize: function (): GetTransactionResponse {
    return {
      outputs: [],
      transfers: [],
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: Partial<GetTransactionResponse>
  ): Record<string, unknown> {
    const json: Record<string, unknown> = {};
    if (msg.outputs?.length) {
      json.outputs = msg.outputs.map(OutputJSON._writeMessage);
    }
    if (msg.transfers?.length) {
      json.transfers = msg.transfers.map(TransferJSON._writeMessage);
    }
    return json;
  },

  /**
   * @private
   */
  _readMessage: function (
    msg: GetTransactionResponse,
    json: any
  ): GetTransactionResponse {
    const _outputs = json.outputs;
    if (_outputs) {
      for (const item of _outputs) {
        const m = Output.initialize();
        OutputJSON._readMessage(m, item);
        msg.outputs.push(m);
      }
    }
    const _transfers = json.transfers;
    if (_transfers) {
      for (const item of _transfers) {
        const m = Transfer.initialize();
        TransferJSON._readMessage(m, item);
        msg.transfers.push(m);
      }
    }
    return msg;
  },
};

export const GetTransactionByTraceIDRequestJSON = {
  /**
   * Serializes GetTransactionByTraceIDRequest to JSON.
   */
  encode: function (msg: Partial<GetTransactionByTraceIDRequest>): string {
    return JSON.stringify(
      GetTransactionByTraceIDRequestJSON._writeMessage(msg)
    );
  },

  /**
   * Deserializes GetTransactionByTraceIDRequest from JSON.
   */
  decode: function (json: string): GetTransactionByTraceIDRequest {
    return GetTransactionByTraceIDRequestJSON._readMessage(
      GetTransactionByTraceIDRequestJSON.initialize(),
      JSON.parse(json)
    );
  },

  /**
   * Initializes GetTransactionByTraceIDRequest with all fields set to their default value.
   */
  initialize: function (): GetTransactionByTraceIDRequest {
    return {
      traceId: "",
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: Partial<GetTransactionByTraceIDRequest>
  ): Record<string, unknown> {
    const json: Record<string, unknown> = {};
    if (msg.traceId) {
      json.traceId = msg.traceId;
    }
    return json;
  },

  /**
   * @private
   */
  _readMessage: function (
    msg: GetTransactionByTraceIDRequest,
    json: any
  ): GetTransactionByTraceIDRequest {
    const _traceId = json.traceId ?? json.trace_id;
    if (_traceId) {
      msg.traceId = _traceId;
    }
    return msg;
  },
};

export const GetTransactionByTraceIDResponseJSON = {
  /**
   * Serializes GetTransactionByTraceIDResponse to JSON.
   */
  encode: function (msg: Partial<GetTransactionByTraceIDResponse>): string {
    return JSON.stringify(
      GetTransactionByTraceIDResponseJSON._writeMessage(msg)
    );
  },

  /**
   * Deserializes GetTransactionByTraceIDResponse from JSON.
   */
  decode: function (json: string): GetTransactionByTraceIDResponse {
    return GetTransactionByTraceIDResponseJSON._readMessage(
      GetTransactionByTraceIDResponseJSON.initialize(),
      JSON.parse(json)
    );
  },

  /**
   * Initializes GetTransactionByTraceIDResponse with all fields set to their default value.
   */
  initialize: function (): GetTransactionByTraceIDResponse {
    return {
      outputs: [],
      transfers: [],
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: Partial<GetTransactionByTraceIDResponse>
  ): Record<string, unknown> {
    const json: Record<string, unknown> = {};
    if (msg.outputs?.length) {
      json.outputs = msg.outputs.map(OutputJSON._writeMessage);
    }
    if (msg.transfers?.length) {
      json.transfers = msg.transfers.map(TransferJSON._writeMessage);
    }
    return json;
  },

  /**
   * @private
   */
  _readMessage: function (
    msg: GetTransactionByTraceIDResponse,
    json: any
  ): GetTransactionByTraceIDResponse {
    const _outputs = json.outputs;
    if (_outputs) {
      for (const item of _outputs) {
        const m = Output.initialize();
        OutputJSON._readMessage(m, item);
        msg.outputs.push(m);
      }
    }
    const _transfers = json.transfers;
    if (_transfers) {
      for (const item of _transfers) {
        const m = Transfer.initialize();
        TransferJSON._readMessage(m, item);
        msg.transfers.push(m);
      }
    }
    return msg;
  },
};