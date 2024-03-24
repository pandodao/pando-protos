// THIS IS AN AUTOGENERATED FILE. DO NOT EDIT THIS FILE DIRECTLY.
// Source: common/v1/common.proto
/* eslint-disable */

import type { ByteSource, PartialDeep } from "protoscript";
import * as protoscript from "protoscript";

//========================================//
//                 Types                  //
//========================================//

export interface Asset {
  assetId: string;
  chainId: string;
  symbol: string;
  name: string;
  logo: string;
  price: string;
}

export interface PaginationRequest {
  cursor: string;
  limit: bigint;
}

export interface PaginationResponse {
  nextCursor: string;
  hasNext: boolean;
}

//========================================//
//        Protobuf Encode / Decode        //
//========================================//

export const Asset = {
  /**
   * Serializes Asset to protobuf.
   */
  encode: function (msg: PartialDeep<Asset>): Uint8Array {
    return Asset._writeMessage(
      msg,
      new protoscript.BinaryWriter(),
    ).getResultBuffer();
  },

  /**
   * Deserializes Asset from protobuf.
   */
  decode: function (bytes: ByteSource): Asset {
    return Asset._readMessage(
      Asset.initialize(),
      new protoscript.BinaryReader(bytes),
    );
  },

  /**
   * Initializes Asset with all fields set to their default value.
   */
  initialize: function (msg?: Partial<Asset>): Asset {
    return {
      assetId: "",
      chainId: "",
      symbol: "",
      name: "",
      logo: "",
      price: "",
      ...msg,
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: PartialDeep<Asset>,
    writer: protoscript.BinaryWriter,
  ): protoscript.BinaryWriter {
    if (msg.assetId) {
      writer.writeString(1, msg.assetId);
    }
    if (msg.chainId) {
      writer.writeString(2, msg.chainId);
    }
    if (msg.symbol) {
      writer.writeString(3, msg.symbol);
    }
    if (msg.name) {
      writer.writeString(4, msg.name);
    }
    if (msg.logo) {
      writer.writeString(5, msg.logo);
    }
    if (msg.price) {
      writer.writeString(6, msg.price);
    }
    return writer;
  },

  /**
   * @private
   */
  _readMessage: function (msg: Asset, reader: protoscript.BinaryReader): Asset {
    while (reader.nextField()) {
      const field = reader.getFieldNumber();
      switch (field) {
        case 1: {
          msg.assetId = reader.readString();
          break;
        }
        case 2: {
          msg.chainId = reader.readString();
          break;
        }
        case 3: {
          msg.symbol = reader.readString();
          break;
        }
        case 4: {
          msg.name = reader.readString();
          break;
        }
        case 5: {
          msg.logo = reader.readString();
          break;
        }
        case 6: {
          msg.price = reader.readString();
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

export const PaginationRequest = {
  /**
   * Serializes PaginationRequest to protobuf.
   */
  encode: function (msg: PartialDeep<PaginationRequest>): Uint8Array {
    return PaginationRequest._writeMessage(
      msg,
      new protoscript.BinaryWriter(),
    ).getResultBuffer();
  },

  /**
   * Deserializes PaginationRequest from protobuf.
   */
  decode: function (bytes: ByteSource): PaginationRequest {
    return PaginationRequest._readMessage(
      PaginationRequest.initialize(),
      new protoscript.BinaryReader(bytes),
    );
  },

  /**
   * Initializes PaginationRequest with all fields set to their default value.
   */
  initialize: function (msg?: Partial<PaginationRequest>): PaginationRequest {
    return {
      cursor: "",
      limit: 0n,
      ...msg,
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: PartialDeep<PaginationRequest>,
    writer: protoscript.BinaryWriter,
  ): protoscript.BinaryWriter {
    if (msg.cursor) {
      writer.writeString(1, msg.cursor);
    }
    if (msg.limit) {
      writer.writeInt64String(2, msg.limit.toString() as any);
    }
    return writer;
  },

  /**
   * @private
   */
  _readMessage: function (
    msg: PaginationRequest,
    reader: protoscript.BinaryReader,
  ): PaginationRequest {
    while (reader.nextField()) {
      const field = reader.getFieldNumber();
      switch (field) {
        case 1: {
          msg.cursor = reader.readString();
          break;
        }
        case 2: {
          msg.limit = BigInt(reader.readInt64String());
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

export const PaginationResponse = {
  /**
   * Serializes PaginationResponse to protobuf.
   */
  encode: function (msg: PartialDeep<PaginationResponse>): Uint8Array {
    return PaginationResponse._writeMessage(
      msg,
      new protoscript.BinaryWriter(),
    ).getResultBuffer();
  },

  /**
   * Deserializes PaginationResponse from protobuf.
   */
  decode: function (bytes: ByteSource): PaginationResponse {
    return PaginationResponse._readMessage(
      PaginationResponse.initialize(),
      new protoscript.BinaryReader(bytes),
    );
  },

  /**
   * Initializes PaginationResponse with all fields set to their default value.
   */
  initialize: function (msg?: Partial<PaginationResponse>): PaginationResponse {
    return {
      nextCursor: "",
      hasNext: false,
      ...msg,
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: PartialDeep<PaginationResponse>,
    writer: protoscript.BinaryWriter,
  ): protoscript.BinaryWriter {
    if (msg.nextCursor) {
      writer.writeString(1, msg.nextCursor);
    }
    if (msg.hasNext) {
      writer.writeBool(2, msg.hasNext);
    }
    return writer;
  },

  /**
   * @private
   */
  _readMessage: function (
    msg: PaginationResponse,
    reader: protoscript.BinaryReader,
  ): PaginationResponse {
    while (reader.nextField()) {
      const field = reader.getFieldNumber();
      switch (field) {
        case 1: {
          msg.nextCursor = reader.readString();
          break;
        }
        case 2: {
          msg.hasNext = reader.readBool();
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

export const AssetJSON = {
  /**
   * Serializes Asset to JSON.
   */
  encode: function (msg: PartialDeep<Asset>): string {
    return JSON.stringify(AssetJSON._writeMessage(msg));
  },

  /**
   * Deserializes Asset from JSON.
   */
  decode: function (json: string): Asset {
    return AssetJSON._readMessage(AssetJSON.initialize(), JSON.parse(json));
  },

  /**
   * Initializes Asset with all fields set to their default value.
   */
  initialize: function (msg?: Partial<Asset>): Asset {
    return {
      assetId: "",
      chainId: "",
      symbol: "",
      name: "",
      logo: "",
      price: "",
      ...msg,
    };
  },

  /**
   * @private
   */
  _writeMessage: function (msg: PartialDeep<Asset>): Record<string, unknown> {
    const json: Record<string, unknown> = {};
    if (msg.assetId) {
      json["assetId"] = msg.assetId;
    }
    if (msg.chainId) {
      json["chainId"] = msg.chainId;
    }
    if (msg.symbol) {
      json["symbol"] = msg.symbol;
    }
    if (msg.name) {
      json["name"] = msg.name;
    }
    if (msg.logo) {
      json["logo"] = msg.logo;
    }
    if (msg.price) {
      json["price"] = msg.price;
    }
    return json;
  },

  /**
   * @private
   */
  _readMessage: function (msg: Asset, json: any): Asset {
    const _assetId_ = json["assetId"] ?? json["asset_id"];
    if (_assetId_) {
      msg.assetId = _assetId_;
    }
    const _chainId_ = json["chainId"] ?? json["chain_id"];
    if (_chainId_) {
      msg.chainId = _chainId_;
    }
    const _symbol_ = json["symbol"];
    if (_symbol_) {
      msg.symbol = _symbol_;
    }
    const _name_ = json["name"];
    if (_name_) {
      msg.name = _name_;
    }
    const _logo_ = json["logo"];
    if (_logo_) {
      msg.logo = _logo_;
    }
    const _price_ = json["price"];
    if (_price_) {
      msg.price = _price_;
    }
    return msg;
  },
};

export const PaginationRequestJSON = {
  /**
   * Serializes PaginationRequest to JSON.
   */
  encode: function (msg: PartialDeep<PaginationRequest>): string {
    return JSON.stringify(PaginationRequestJSON._writeMessage(msg));
  },

  /**
   * Deserializes PaginationRequest from JSON.
   */
  decode: function (json: string): PaginationRequest {
    return PaginationRequestJSON._readMessage(
      PaginationRequestJSON.initialize(),
      JSON.parse(json),
    );
  },

  /**
   * Initializes PaginationRequest with all fields set to their default value.
   */
  initialize: function (msg?: Partial<PaginationRequest>): PaginationRequest {
    return {
      cursor: "",
      limit: 0n,
      ...msg,
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: PartialDeep<PaginationRequest>,
  ): Record<string, unknown> {
    const json: Record<string, unknown> = {};
    if (msg.cursor) {
      json["cursor"] = msg.cursor;
    }
    if (msg.limit) {
      json["limit"] = String(msg.limit);
    }
    return json;
  },

  /**
   * @private
   */
  _readMessage: function (
    msg: PaginationRequest,
    json: any,
  ): PaginationRequest {
    const _cursor_ = json["cursor"];
    if (_cursor_) {
      msg.cursor = _cursor_;
    }
    const _limit_ = json["limit"];
    if (_limit_) {
      msg.limit = BigInt(_limit_);
    }
    return msg;
  },
};

export const PaginationResponseJSON = {
  /**
   * Serializes PaginationResponse to JSON.
   */
  encode: function (msg: PartialDeep<PaginationResponse>): string {
    return JSON.stringify(PaginationResponseJSON._writeMessage(msg));
  },

  /**
   * Deserializes PaginationResponse from JSON.
   */
  decode: function (json: string): PaginationResponse {
    return PaginationResponseJSON._readMessage(
      PaginationResponseJSON.initialize(),
      JSON.parse(json),
    );
  },

  /**
   * Initializes PaginationResponse with all fields set to their default value.
   */
  initialize: function (msg?: Partial<PaginationResponse>): PaginationResponse {
    return {
      nextCursor: "",
      hasNext: false,
      ...msg,
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: PartialDeep<PaginationResponse>,
  ): Record<string, unknown> {
    const json: Record<string, unknown> = {};
    if (msg.nextCursor) {
      json["nextCursor"] = msg.nextCursor;
    }
    if (msg.hasNext) {
      json["hasNext"] = msg.hasNext;
    }
    return json;
  },

  /**
   * @private
   */
  _readMessage: function (
    msg: PaginationResponse,
    json: any,
  ): PaginationResponse {
    const _nextCursor_ = json["nextCursor"] ?? json["next_cursor"];
    if (_nextCursor_) {
      msg.nextCursor = _nextCursor_;
    }
    const _hasNext_ = json["hasNext"] ?? json["has_next"];
    if (_hasNext_) {
      msg.hasNext = _hasNext_;
    }
    return msg;
  },
};
