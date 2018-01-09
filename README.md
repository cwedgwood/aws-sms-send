aws-sms-send
===

This is a minimalist tool I made to send messages using the AWS
SNS/SMS API.

Usage
---

    ./aws-sms-send -verbose +15551231234 "your pizza is ready"

    ./aws-sms-send -transactional +15551231234 "ALERT: Thromdibulator malfunction, rebooting."
