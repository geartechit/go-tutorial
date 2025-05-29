alter table employees
    drop column "department_id" cascade ;

-- alter table employees
--     drop constraint fk_employee_department foreign key (department_id) references departments(id);
