# Project - KORA

## Sprint 3
 

We have accomplished in the project according to the issue/stories conveyed (closed stories in sprint_3) 
#### ISSUE 1 - Created an API to return user details
#### ISSUE 2 - Created an API for Deleting user
#### ISSUE 3 - Implemented Own account page
#### ISSUE 4 - Implemented Add question screen popup
#### ISSUE 5 - Changes for Questions collection in DB
#### ISSUE 6 - Created an API to return Questions based on user categories
#### ISSUE 7 - Added search api for fetching matching questions

# Backend Development server
Run the below commands to start the development server.

> `cd ../backend`

> `go run main.go `

### Back-end screenshot
![image](https://user-images.githubusercontent.com/38340989/156869950-cc8f8cd0-4a3a-49c9-b4bf-9dbcd394a263.png)

# Backend
1.API :: 
The below API's are used 
1.Added an API to return user details
2.Added an API for related top question
3.Added an API for deleting user
4.Added search api for fetching matching questions

2.Changes for Question Collection in DB
3.Added test cases

## 1. API to return User Details
The `/fetch-user` API is used to fetch user details.
![userown](https://user-images.githubusercontent.com/38340989/161362300-59888919-2224-4791-9ceb-fbde2db2dd55.jpeg)

Other user details by passing username.
![WhatsApp Image 2022-04-01 at 8 26 40 PM](https://user-images.githubusercontent.com/38340989/161362339-06b74bc7-0ef3-4fba-9cc1-b9ed1ee043a4.jpeg)

## 2. API for related top question
The `/topquestion` API is used to fetch Questions based on the category selected by user.
![image](https://user-images.githubusercontent.com/38340989/161362416-94580694-30fc-4774-acab-7187fecfca69.png)


## 3. API for deleting user
The `/deleteuser` API is used to delete user.


## 4. Search api for fetching matching questions
The `/search` API is used to search and find matching questions.
![search](https://user-images.githubusercontent.com/38340989/161362853-7c976786-34e8-40d6-a716-055898ebffac.jpeg)

## Added Test cases

run `go test`
successfully ran the test cases for backend

Test cases for Fetch user and search API
![WhatsApp Image 2022-04-01 at 8 28 28 PM](https://user-images.githubusercontent.com/38340989/161362883-29381eeb-4d84-4876-8da8-af6814e8426c.jpeg)

Test cases for Delete user, Related Top question and selected question
![image](https://user-images.githubusercontent.com/38340989/161363017-4d93c958-d579-45f2-806e-abf5a9d77b17.png)

![image](https://user-images.githubusercontent.com/38340989/161363044-948bf8d0-a494-428c-860a-f6fd21ab4e7b.png)



### Front-end screenshot
<img width="1434" alt="Screenshot 2022-04-01 at 9 14 08 PM" src="https://user-images.githubusercontent.com/40836174/161361030-639960d7-66e4-4a03-b64b-c0158bb04010.png">
<img width="1436" alt="Screenshot 2022-04-01 at 9 14 02 PM" src="https://user-images.githubusercontent.com/40836174/161361050-5b04d116-239e-43c7-9fea-9da07310e136.png">> 
<img width="1440" alt="Screenshot 2022-04-01 at 9 13 58 PM" src="https://user-images.githubusercontent.com/40836174/161361071-38071d1a-2d3b-45f1-a6fc-6fd69a0106f5.png">
<img width="1436" alt="Screenshot 2022-04-01 at 9 14 31 PM" src="https://user-images.githubusercontent.com/40836174/161361093-a181b450-2e53-4aea-8a1e-4d99a22ae743.png">
<img width="1437" alt="Screenshot 2022-04-01 at 9 14 37 PM" src="https://user-images.githubusercontent.com/40836174/161361106-3af6fae0-f65f-4041-8718-eab32b0158be.png">
<img width="1435" alt="Screenshot 2022-04-01 at 9 14 21 PM" src="https://user-images.githubusercontent.com/40836174/161361134-60426b81-dad4-4900-a3b0-e19a6c8e19ee.png">
<img width="1434" alt="Screenshot 2022-04-01 at 9 14 15 PM" src="https://user-images.githubusercontent.com/40836174/161361148-f341f64b-db73-49f0-9552-2b4343c8ae57.png">

