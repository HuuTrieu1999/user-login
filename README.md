# Login

## Overview
This document outlines the design and development of a highly scalable login and user creation system for a mobile application. The system is engineered to handle a significant user base, specifically targeting 100,000 concurrent users.

## System Design

<img src="./asset/login.drawio.png">

### User

| Field        | Type     | Description        | 
|:-------------|:---------|:-------------------| 
| _id          | ObjectID | ID of the user     | 
| userName     | string   | user name          |
| fullName     | string   | full name of user  |
| email        | string   | email of user      |
| phoneNumber  | string   | phone number       |
| hashPassword | string   | hashed password    |
| birthday     | string   | format: yyyy-mm-dd |
| createdDate  | string   | Time UTC +7        |

Index:
* email_idx (email): unique
* phone_idx (phoneNumber): unique
* username_idx (userName): unique

## API Spec

BaseResponse

| Field        | Type   | Description     | 
|:-------------|:-------|:----------------| 
| status       | bool   | success or fail | 
| errorCode    | string | error code      |
| errorMessage | string | error message   |
| data         | string | data json       |

**POST: /api/login**

Request

| Field    | Type   | Description          | 
|:---------|:-------|:---------------------| 
| account  | string | email/phone/username | 
| password | string | password             |

LoginDataResponse

| Field            | Type   | Description        | 
|:-----------------|:-------|:-------------------| 
| token            | string | jwt                | 
| user.UserID      | string |                    |
| user.UserName    | string |                    |
| user.FullName    | string |                    |
| user.Email       | string |                    |
| user.PhoneNumber | string |                    |
| user.Birthday    | string | format: yyyy-mm-dd |

**POST: /api/register**

Request

| Field       | Type    | Description        | 
|:------------|:--------|:-------------------| 
| userName    | string  |                    |
| fullName    | string  |                    |
| email       | string  |                    |
| phoneNumber | string  |                    |
| birthday    | string  | format: yyyy-mm-dd | 
| password    | string  | password           |

RegisterDataResponse

| Field            | Type   | Description        | 
|:-----------------|:-------|:-------------------|
| user.UserID      | string |                    |
| user.UserName    | string |                    |
| user.FullName    | string |                    |
| user.Email       | string |                    |
| user.PhoneNumber | string |                    |
| user.Birthday    | string | format: yyyy-mm-dd |


