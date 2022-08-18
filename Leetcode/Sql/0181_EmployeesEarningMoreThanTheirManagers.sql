select e1.name as Employee
from Employee as e1, Employee as e2
where e1.Salary > e2.Salary
and e1.managerId = e2.id;