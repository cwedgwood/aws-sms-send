aws-sms-send
===

This is a minimalist tool I made to send messages using the AWS
SNS/SMS API.

Usage
---

    ./aws-sms-send -verbose +15551231234 "your pizza is ready"

    ./aws-sms-send -transactional +15551231234 "ALERT: Thromdibulator malfunction, rebooting."

How to get
---

If you have Go:

	go get github.com/cwedgwood/aws-sms-send


Or else you might choose to clone & build:

    git clone https://github.com/cwedgwood/aws-sms-send.git
	cd aws-sms-send
	make
	#or
	sudo make install
