# mocksponse

An api response mocking server

## Example Recipe File

```yaml
routes:
  /:
    get:
      sequential: true
      responses:
        - code: 200
          data: |
            {
              "some": "data",
              "to": "return"
            }
        - code: 500
          data: "{\"error\": \"something went wrong\"}"
  /animals/:
    post:
      responses:
        - code: 204
          data: ""
    get:
      responses:
        - code: 200
          data: |
            [
              {
                "animal": "dog",
                "name": "spot"
              }
            ]
  /animals/{name:[a-zA-Z]}/:
    put:
      responses:
        - code: 200
          data: |
            {
              "animal": "bear",
              "name": "{{ index .Vars "name" }}"
            }
    get:
      responses:
        - code: 401
          file: ../permission_denied.json
default:
  responses:
    - code: 404
      data: |
        {"error": "route not found"}
```
