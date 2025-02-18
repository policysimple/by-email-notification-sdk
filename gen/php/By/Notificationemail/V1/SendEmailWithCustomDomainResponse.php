<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: by/notificationemail/v1/by_email_notification.proto

namespace By\Notificationemail\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * SendEmailWithCustomDomainResponse...
 *
 * Generated from protobuf message <code>by.notificationemail.v1.SendEmailWithCustomDomainResponse</code>
 */
class SendEmailWithCustomDomainResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string status = 1 [json_name = "status"];</code>
     */
    protected $status = '';
    /**
     * Generated from protobuf field <code>string message = 2 [json_name = "message"];</code>
     */
    protected $message = '';
    /**
     * Generated from protobuf field <code>.by.notificationemail.v1.SendEmailInfoResponse info = 3 [json_name = "info"];</code>
     */
    protected $info = null;
    /**
     * Generated from protobuf field <code>string error = 4 [json_name = "error"];</code>
     */
    protected $error = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $status
     *     @type string $message
     *     @type \By\Notificationemail\V1\SendEmailInfoResponse $info
     *     @type string $error
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\By\Notificationemail\V1\ByEmailNotification::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string status = 1 [json_name = "status"];</code>
     * @return string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * Generated from protobuf field <code>string status = 1 [json_name = "status"];</code>
     * @param string $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkString($var, True);
        $this->status = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string message = 2 [json_name = "message"];</code>
     * @return string
     */
    public function getMessage()
    {
        return $this->message;
    }

    /**
     * Generated from protobuf field <code>string message = 2 [json_name = "message"];</code>
     * @param string $var
     * @return $this
     */
    public function setMessage($var)
    {
        GPBUtil::checkString($var, True);
        $this->message = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.by.notificationemail.v1.SendEmailInfoResponse info = 3 [json_name = "info"];</code>
     * @return \By\Notificationemail\V1\SendEmailInfoResponse|null
     */
    public function getInfo()
    {
        return $this->info;
    }

    public function hasInfo()
    {
        return isset($this->info);
    }

    public function clearInfo()
    {
        unset($this->info);
    }

    /**
     * Generated from protobuf field <code>.by.notificationemail.v1.SendEmailInfoResponse info = 3 [json_name = "info"];</code>
     * @param \By\Notificationemail\V1\SendEmailInfoResponse $var
     * @return $this
     */
    public function setInfo($var)
    {
        GPBUtil::checkMessage($var, \By\Notificationemail\V1\SendEmailInfoResponse::class);
        $this->info = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string error = 4 [json_name = "error"];</code>
     * @return string
     */
    public function getError()
    {
        return $this->error;
    }

    /**
     * Generated from protobuf field <code>string error = 4 [json_name = "error"];</code>
     * @param string $var
     * @return $this
     */
    public function setError($var)
    {
        GPBUtil::checkString($var, True);
        $this->error = $var;

        return $this;
    }

}

