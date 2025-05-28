```
HTTP Request (DTO)  --->  handler
    handler parses DTO, maps to domain.Employee
        v
    service.Create(ctx, *Employee)
        v
    repository.Create(ctx, CreateEmployeeParams)
```