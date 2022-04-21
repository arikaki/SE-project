# Project - KORA

## Sprint 4


We have accomplished in the project according to the issue/stories conveyed (closed stories in sprint_3) 
#### ISSUE 1 - Home page issue
#### ISSUE 2 - Question card Improvement
#### ISSUE 3 - Save the added question to the database
#### ISSUE 4 - Integraion of the files so far
#### ISSUE 5 - API for unaswered questions
#### ISSUE 6 - API to upvote,downvote question
#### ISSUE 7 - API to Report Question
#### ISSUE 8 - Test cases for API's
#### ISSUE 9 - API to upvote, downvote answers
#### ISSUE 10 - API to save answer to DB
#### ISSUE 11 - Review and fixes

# Backend Development server
Run the below commands to start the development server.

> `cd ../backend`

> `go run main.go `

![image](https://user-images.githubusercontent.com/38340989/156869950-cc8f8cd0-4a3a-49c9-b4bf-9dbcd394a263.png)

# Backend
1.API :: 
The below API's are used 
1.Added an API for unaswered questions
2.Added an API to upvote,downvote question
3.Added an API to Report Question
4.Added an API to save answer to DB
5.Added an API to upvote,downvote answers

2.Changes to DB for Report API
3.Added test cases
4.Review and fixes
## 1. API for unaswered questions
The `/getUnanswered` API is used to fetch unanswered questions.

## 2. API to upvote,downvote question
The `/upvotequestion` API is used to upvote questions.
![image](https://user-images.githubusercontent.com/38340989/164363075-7deb3a42-ac15-4b4c-8a95-6a079b9b7119.png)


The `/downvotequestion` API is used to downvote questions.
![image](https://user-images.githubusercontent.com/38340989/164363007-0f943aad-8044-4d44-9ea8-08d621fb5ed2.png)


## 3. API to Report Question
The `/report` API is used to report questions.
![image](https://user-images.githubusercontent.com/38340989/164362593-a950ae45-2148-46fb-bbbd-5ddabac980ec.png)

![image](https://user-images.githubusercontent.com/38340989/164362757-b9050f0a-f636-42fc-9ff9-a1d38a91f0e1.png)


## 3. API to save answer to db
The `/` API is used to report questions.



## 4. API to upvote,downvote answers
The `/upvoteanswer` API is used to upvote answers.
![image](https://user-images.githubusercontent.com/38340989/164362893-ca635ac0-7195-4e0e-9370-0634abdb96fd.png)


The `/downvoteanswer` API is used to downvote answers.
![image](https://user-images.githubusercontent.com/38340989/164362935-3273ba04-a0d0-4167-b56e-381a45e9e7b4.png)



## Added Test cases

run `go test`
successfully ran the test cases for backend.
![image](https://user-images.githubusercontent.com/38340989/164363286-f687e916-dd7f-4334-82ce-ee1ee4499547.png)

![image](https://user-images.githubusercontent.com/38340989/164363361-8db4dd18-f074-41cf-b198-fe45b323446b.png)

![image](https://user-images.githubusercontent.com/38340989/164363422-df202c89-1ae3-4286-b68a-9b52e9f5a87d.png)

![image](https://user-images.githubusercontent.com/38340989/164363455-14ccf78f-b019-4bd5-afb0-ffe751a30e6b.png)

![image](https://user-images.githubusercontent.com/38340989/164363491-974d7e8b-17fd-4f24-86a0-d7c2faeb647e.png)

![image](https://user-images.githubusercontent.com/38340989/164363528-79abc3c7-f0af-45c7-bc27-ef133d64b62e.png)

![image](https://user-images.githubusercontent.com/38340989/164363566-a3a7852a-39b0-4ed0-bcfa-ff0ca4b43534.png)

![image](https://user-images.githubusercontent.com/38340989/164363733-190b6553-2465-40a8-9188-40adecfe4aca.png)





# Frontend Web development

Run the below commands to start the development server.

> `cd ../frontend`
> `npm start `

Run the below commands to start the cypress testing.
> `npx cypress open`

### Option to choose topics dropdown

### Login page 

### Category page (selection atleast 5)

### Answer page 

### Profile page


## Cypress testing video




## Frontend
1. Added dropdown in the post question (Containing topics)
2. Created post,upvote, downvote, report and delete option in the question box
3. Created Answer page with answer box
4. Enhancement in Category page and routing
