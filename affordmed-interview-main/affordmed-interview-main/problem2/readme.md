Develop a HTTP microservice called prefix-management-service that exposes a GET REST API "/prefixes". This API receives a query parameter. This query param is called "keywords".

http://localhost:port/prefixes?keywords=bonfire,bool

Now, consider we have list of words already present in the server

[bonfire, cardio, case, character, bonsai]

```
Your task is to:

1. find if the keywords are found in the server or not
2. if any keyword is found, you need to return the smallest unique prefix that identifies the keyword
```

Response of above API call would be like this

```json
[
    {
        "keyword": "bonfire",
        "status": "found",
        "prefix": "bonf"
    },
    {
        "keyword": "bool",
        "status": "not_found",
        "prefix": "not_applicable"
    }
]
```

Please note you can hardcode 20 words in the server for testing.

Test Case 1:

http://localhost:port/prefixes?keywords=bonfire,bonsai

Expected Output:

```json
[
    {
        "keyword": "bonfire",
        "status": "found",
        "prefix": "bonf"
    },
    {
        "keyword": "bonsai",
        "status": "found",
        "prefix": "bons"
    }
]
```

### Please attach a screenshot of Insomnia or Postman making an API call with response body
