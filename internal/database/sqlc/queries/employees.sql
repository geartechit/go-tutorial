-- name: GetAllEmployee :many
select * from employees;

-- name: GetEmployeeById :one
select * from employees
where id = @id;

-- name: CreateEmployee :one
insert into employees (
    name,
    dob,
    department,
    job_title,
    address,
    joined_at
)
values (
    @name,
    @dob,
    @department,
    @job_title,
    @address,
    @joined_at
)
returning *;

-- name: UpdateEmployee :one
update employees
set
--     name = coalesce(sqlc.narg('name'), name),
--     dob = coalesce(sqlc.narg('dob'), dob),
--     department = coalesce(sqlc.narg('department'), department),
--     job_title = coalesce(sqlc.narg('job_title'), job_title),
--     address = coalesce(sqlc.narg('address'), address),
--     joined_at = coalesce(sqlc.narg('joined_at'), joined_at)
    name = coalesce(@name, name),
    dob = coalesce(@dob, dob),
    department = coalesce(@department, department),
    job_title = coalesce(@job_title, job_title),
    address = coalesce(@address, address),
    joined_at = coalesce(@joined_at, joined_at)
where id = @id
returning *;

-- name: DeleteEmployee :one
delete from employees
where id = @id
returning id;