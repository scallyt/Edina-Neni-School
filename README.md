# Edina-Neni-School

## Endpoints

| URL          | HTTP method | Auth | JSON Response     |
| ------------ | ----------- | ---- | ----------------- |
| /user/login  | POST        |      | user's token      |
| /tasks       | GET         |      | all task          |
| /tasks       | POST        | Y    | new task added    |
| /tasks       | PATCH       | Y    | task product      |
| /tasks       | DELETE      | Y    | id                |

tasks: ``` json
        
        tasks: {
            "Date": "ect..",
            "text": "ect..",

            "createdAt": "ect..",
            "updatesAt": "ect.."
        }

       ```

