let rec pow x y = if y = 0 then 1 else x * pow x (y - 1)
