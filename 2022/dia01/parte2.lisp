

(print (let ((sums (loop with lines = (uiop:read-file-lines "calories.txt")
			 for x in lines
			 if (equalp x "")
			     collect sum into sums
			     and do (setf sum 0)
			 else sum (parse-integer x) into sum
			 finally (progn (push sum sums) ; incluir ultima soma antes do EOF
					(return sums)))))
	 (let ((sorted (sort sums #'>)))
	   (+ (first sorted) (second sorted) (third sorted)))))
