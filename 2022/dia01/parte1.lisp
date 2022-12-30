

(print (loop with lines = (uiop:read-file-lines "calories.txt")
	     for x in lines
	     if (equalp x "")
		 do (setf sum 0)
	     else sum (parse-integer x) into sum
	     maximize sum into max
	     finally (return max)))
