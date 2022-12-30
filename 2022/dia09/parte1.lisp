


(defun main () 
(let ((lines (with-open-file (op "rope.txt" :direction :input)
	       (loop for x = (read-line op nil)
		     while x collect x))))
  (aref (car lines) 0)))
