// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// Copyright 2019 Cuemby, LLC
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//
'use strict';
var grpc = require('grpc');
var by_notificationemail_v1_by_email_notification_pb = require('../../../by/notificationemail/v1/by_email_notification_pb.js');

function serialize_by_notificationemail_v1_SendEmailRequest(arg) {
  if (!(arg instanceof by_notificationemail_v1_by_email_notification_pb.SendEmailRequest)) {
    throw new Error('Expected argument of type by.notificationemail.v1.SendEmailRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_by_notificationemail_v1_SendEmailRequest(buffer_arg) {
  return by_notificationemail_v1_by_email_notification_pb.SendEmailRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_by_notificationemail_v1_SendEmailResponse(arg) {
  if (!(arg instanceof by_notificationemail_v1_by_email_notification_pb.SendEmailResponse)) {
    throw new Error('Expected argument of type by.notificationemail.v1.SendEmailResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_by_notificationemail_v1_SendEmailResponse(buffer_arg) {
  return by_notificationemail_v1_by_email_notification_pb.SendEmailResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// NotificationEmailAPI ...
var NotificationEmailAPIServiceService = exports.NotificationEmailAPIServiceService = {
  // SendEmail ...
  sendEmail: {
    path: '/by.notificationemail.v1.NotificationEmailAPIService/SendEmail',
    requestStream: false,
    responseStream: false,
    requestType: by_notificationemail_v1_by_email_notification_pb.SendEmailRequest,
    responseType: by_notificationemail_v1_by_email_notification_pb.SendEmailResponse,
    requestSerialize: serialize_by_notificationemail_v1_SendEmailRequest,
    requestDeserialize: deserialize_by_notificationemail_v1_SendEmailRequest,
    responseSerialize: serialize_by_notificationemail_v1_SendEmailResponse,
    responseDeserialize: deserialize_by_notificationemail_v1_SendEmailResponse,
  },
};

exports.NotificationEmailAPIServiceClient = grpc.makeGenericClientConstructor(NotificationEmailAPIServiceService);
