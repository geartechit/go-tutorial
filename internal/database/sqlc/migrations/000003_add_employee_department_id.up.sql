alter table employees
    add column "department_id" varchar;

alter table employees
    add constraint fk_employee_department foreign key (department_id) references departments(id);
