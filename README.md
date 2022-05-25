# go-email-identify
## 1. introduction
The system of user login email authentication based on go lang

## 2.design ideas
user input username， password and eamil.
user click button to get code. the code is exist in the redis? if tue, try it agin after 1 minute.
for another, send code to user email, store code to redis, set the expired time for one minute.

Attention: the email_utils contains my private info, so i will upload this file at tomorrow.
