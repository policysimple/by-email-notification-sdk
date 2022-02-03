<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: by/notificationemail/v1/by_email_notification.proto

namespace By\Notificationemail\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * SendEmailRequest ...
 *
 * Generated from protobuf message <code>by.notificationemail.v1.SendEmailRequest</code>
 */
class SendEmailRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .by.notificationemail.v1.To to = 1 [json_name = "to"];</code>
     */
    private $to;
    /**
     * Generated from protobuf field <code>repeated .by.notificationemail.v1.Cc cc = 2 [json_name = "cc"];</code>
     */
    private $cc;
    /**
     * Generated from protobuf field <code>repeated .by.notificationemail.v1.Bbc bbc = 3 [json_name = "bbc"];</code>
     */
    private $bbc;
    /**
     * Generated from protobuf field <code>repeated .by.notificationemail.v1.ReplyTo reply_to = 4 [json_name = "replyTo"];</code>
     */
    private $reply_to;
    /**
     * Generated from protobuf field <code>string subject = 5 [json_name = "subject"];</code>
     */
    protected $subject = '';
    /**
     * Generated from protobuf field <code>string tag = 6 [json_name = "tag"];</code>
     */
    protected $tag = '';
    /**
     * Generated from protobuf field <code>string text = 7 [json_name = "text"];</code>
     */
    protected $text = '';
    /**
     * Generated from protobuf field <code>string html = 8 [json_name = "html"];</code>
     */
    protected $html = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \By\Notificationemail\V1\To[]|\Google\Protobuf\Internal\RepeatedField $to
     *     @type \By\Notificationemail\V1\Cc[]|\Google\Protobuf\Internal\RepeatedField $cc
     *     @type \By\Notificationemail\V1\Bbc[]|\Google\Protobuf\Internal\RepeatedField $bbc
     *     @type \By\Notificationemail\V1\ReplyTo[]|\Google\Protobuf\Internal\RepeatedField $reply_to
     *     @type string $subject
     *     @type string $tag
     *     @type string $text
     *     @type string $html
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\By\Notificationemail\V1\ByEmailNotification::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .by.notificationemail.v1.To to = 1 [json_name = "to"];</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getTo()
    {
        return $this->to;
    }

    /**
     * Generated from protobuf field <code>repeated .by.notificationemail.v1.To to = 1 [json_name = "to"];</code>
     * @param \By\Notificationemail\V1\To[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setTo($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \By\Notificationemail\V1\To::class);
        $this->to = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .by.notificationemail.v1.Cc cc = 2 [json_name = "cc"];</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getCc()
    {
        return $this->cc;
    }

    /**
     * Generated from protobuf field <code>repeated .by.notificationemail.v1.Cc cc = 2 [json_name = "cc"];</code>
     * @param \By\Notificationemail\V1\Cc[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setCc($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \By\Notificationemail\V1\Cc::class);
        $this->cc = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .by.notificationemail.v1.Bbc bbc = 3 [json_name = "bbc"];</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getBbc()
    {
        return $this->bbc;
    }

    /**
     * Generated from protobuf field <code>repeated .by.notificationemail.v1.Bbc bbc = 3 [json_name = "bbc"];</code>
     * @param \By\Notificationemail\V1\Bbc[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setBbc($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \By\Notificationemail\V1\Bbc::class);
        $this->bbc = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .by.notificationemail.v1.ReplyTo reply_to = 4 [json_name = "replyTo"];</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getReplyTo()
    {
        return $this->reply_to;
    }

    /**
     * Generated from protobuf field <code>repeated .by.notificationemail.v1.ReplyTo reply_to = 4 [json_name = "replyTo"];</code>
     * @param \By\Notificationemail\V1\ReplyTo[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setReplyTo($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \By\Notificationemail\V1\ReplyTo::class);
        $this->reply_to = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string subject = 5 [json_name = "subject"];</code>
     * @return string
     */
    public function getSubject()
    {
        return $this->subject;
    }

    /**
     * Generated from protobuf field <code>string subject = 5 [json_name = "subject"];</code>
     * @param string $var
     * @return $this
     */
    public function setSubject($var)
    {
        GPBUtil::checkString($var, True);
        $this->subject = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string tag = 6 [json_name = "tag"];</code>
     * @return string
     */
    public function getTag()
    {
        return $this->tag;
    }

    /**
     * Generated from protobuf field <code>string tag = 6 [json_name = "tag"];</code>
     * @param string $var
     * @return $this
     */
    public function setTag($var)
    {
        GPBUtil::checkString($var, True);
        $this->tag = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string text = 7 [json_name = "text"];</code>
     * @return string
     */
    public function getText()
    {
        return $this->text;
    }

    /**
     * Generated from protobuf field <code>string text = 7 [json_name = "text"];</code>
     * @param string $var
     * @return $this
     */
    public function setText($var)
    {
        GPBUtil::checkString($var, True);
        $this->text = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string html = 8 [json_name = "html"];</code>
     * @return string
     */
    public function getHtml()
    {
        return $this->html;
    }

    /**
     * Generated from protobuf field <code>string html = 8 [json_name = "html"];</code>
     * @param string $var
     * @return $this
     */
    public function setHtml($var)
    {
        GPBUtil::checkString($var, True);
        $this->html = $var;

        return $this;
    }

}

