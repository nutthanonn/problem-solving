select employee_id,
case 
    when mod(employee_id, 2)=0 then 0
    when name like 'm%' then 0
    else salary
    end as bonus
from employees