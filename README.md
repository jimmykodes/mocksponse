# mocksponse
An api response mocking server

## Example Recipe File

```yaml
routes:
  - path: /
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
  - path: /animals/{name:[a-zA-Z]}/
    responses:
      - code: 200
        data: |
          {
            "animal": "bear",
            "name": "{{ index .Vars "name" }}"
          }
  - path: /vehicles/
    responses:
      - code: 401
        file: ../permission_denied.json
```
