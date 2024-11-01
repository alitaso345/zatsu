import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb'

export class User extends jspb.Message {
  getId(): number
  setId(value: number): User

  getName(): string
  setName(value: string): User

  getTwitterhashtag(): string
  setTwitterhashtag(value: string): User

  getTwitchchannel(): string
  setTwitchchannel(value: string): User

  serializeBinary(): Uint8Array
  toObject(includeInstance?: boolean): User.AsObject
  static toObject(includeInstance: boolean, msg: User): User.AsObject
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void
  static deserializeBinary(bytes: Uint8Array): User
  static deserializeBinaryFromReader(
    message: User,
    reader: jspb.BinaryReader
  ): User
}

export namespace User {
  export type AsObject = {
    id: number
    name: string
    twitterhashtag: string
    twitchchannel: string
  }
}

export class GetUserRequest extends jspb.Message {
  getName(): string
  setName(value: string): GetUserRequest

  serializeBinary(): Uint8Array
  toObject(includeInstance?: boolean): GetUserRequest.AsObject
  static toObject(
    includeInstance: boolean,
    msg: GetUserRequest
  ): GetUserRequest.AsObject
  static serializeBinaryToWriter(
    message: GetUserRequest,
    writer: jspb.BinaryWriter
  ): void
  static deserializeBinary(bytes: Uint8Array): GetUserRequest
  static deserializeBinaryFromReader(
    message: GetUserRequest,
    reader: jspb.BinaryReader
  ): GetUserRequest
}

export namespace GetUserRequest {
  export type AsObject = {
    name: string
  }
}

export class NewUserRequest extends jspb.Message {
  getName(): string
  setName(value: string): NewUserRequest

  serializeBinary(): Uint8Array
  toObject(includeInstance?: boolean): NewUserRequest.AsObject
  static toObject(
    includeInstance: boolean,
    msg: NewUserRequest
  ): NewUserRequest.AsObject
  static serializeBinaryToWriter(
    message: NewUserRequest,
    writer: jspb.BinaryWriter
  ): void
  static deserializeBinary(bytes: Uint8Array): NewUserRequest
  static deserializeBinaryFromReader(
    message: NewUserRequest,
    reader: jspb.BinaryReader
  ): NewUserRequest
}

export namespace NewUserRequest {
  export type AsObject = {
    name: string
  }
}

export class UpdateUserRequest extends jspb.Message {
  getUser(): User | undefined
  setUser(value?: User): UpdateUserRequest
  hasUser(): boolean
  clearUser(): UpdateUserRequest

  serializeBinary(): Uint8Array
  toObject(includeInstance?: boolean): UpdateUserRequest.AsObject
  static toObject(
    includeInstance: boolean,
    msg: UpdateUserRequest
  ): UpdateUserRequest.AsObject
  static serializeBinaryToWriter(
    message: UpdateUserRequest,
    writer: jspb.BinaryWriter
  ): void
  static deserializeBinary(bytes: Uint8Array): UpdateUserRequest
  static deserializeBinaryFromReader(
    message: UpdateUserRequest,
    reader: jspb.BinaryReader
  ): UpdateUserRequest
}

export namespace UpdateUserRequest {
  export type AsObject = {
    user?: User.AsObject
  }
}

export class UserResponse extends jspb.Message {
  getUser(): User | undefined
  setUser(value?: User): UserResponse
  hasUser(): boolean
  clearUser(): UserResponse

  serializeBinary(): Uint8Array
  toObject(includeInstance?: boolean): UserResponse.AsObject
  static toObject(
    includeInstance: boolean,
    msg: UserResponse
  ): UserResponse.AsObject
  static serializeBinaryToWriter(
    message: UserResponse,
    writer: jspb.BinaryWriter
  ): void
  static deserializeBinary(bytes: Uint8Array): UserResponse
  static deserializeBinaryFromReader(
    message: UserResponse,
    reader: jspb.BinaryReader
  ): UserResponse
}

export namespace UserResponse {
  export type AsObject = {
    user?: User.AsObject
  }
}

export class UsersResponse extends jspb.Message {
  getUsersList(): Array<User>
  setUsersList(value: Array<User>): UsersResponse
  clearUsersList(): UsersResponse
  addUsers(value?: User, index?: number): User

  serializeBinary(): Uint8Array
  toObject(includeInstance?: boolean): UsersResponse.AsObject
  static toObject(
    includeInstance: boolean,
    msg: UsersResponse
  ): UsersResponse.AsObject
  static serializeBinaryToWriter(
    message: UsersResponse,
    writer: jspb.BinaryWriter
  ): void
  static deserializeBinary(bytes: Uint8Array): UsersResponse
  static deserializeBinaryFromReader(
    message: UsersResponse,
    reader: jspb.BinaryReader
  ): UsersResponse
}

export namespace UsersResponse {
  export type AsObject = {
    usersList: Array<User.AsObject>
  }
}
