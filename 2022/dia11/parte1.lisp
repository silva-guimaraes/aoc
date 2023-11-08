


(defclass monkey ()
  ((items :initarg :items :accessor items)
   (operation :initarg :operation :accessor operation)
   (test :initarg :test :accessor test)
   (inspection-count :initform 0 :accessor inspection-count)))

(defun monkey-make (items operation divisible true false)
  (make-instance 'monkey :items items 
	      :operation operation
	      :test (lambda (x) 
			(if (= 0 (mod x divisible)) true false))))

;; exemplo
;; (defvar monkeys (list
;;                   (monkey-make '(79 98) #'(lambda (old) (* old 19)) 23 2 3)
;;                   (monkey-make '(54 65 75 74) #'(lambda (old) (+ old 6)) 19 2 0)
;;                   (monkey-make '(79 60 97) #'(lambda (old) (* old old)) 13 1 3)
;;                   (monkey-make '(74) #'(lambda (old) (+ old 3)) 17 0 1) ))

(defvar monkeys (list  
		  (monkey-make '(80) #'(lambda (old) (* old 5)) 2 4 3)
		  (monkey-make '(75 83 74) #'(lambda (old) (+ old 7)) 7 5 6)
		  (monkey-make '(86 67 61 96 52 63 73) #'(lambda (old) (+ old 5)) 3 7 0)
		  (monkey-make '(85 83 55 85 57 70 85 52) #'(lambda (old) (+ old 8)) 17 1 5)
		  (monkey-make '(67 75 91 72 89) #'(lambda (old) (+ old 4)) 11 3 1)
		  (monkey-make '(66 64 68 92 68 77) #'(lambda (old) (* old 2)) 19 6 2)
		  (monkey-make '(97 94 79 88) #'(lambda (old) (* old old)) 5 2 7)
		  (monkey-make '(77 85) #'(lambda (old) (+ old 6)) 13 4 0)))

; isso não funciona dentro do loop. o macaco não recebe o item. não sei porque
(defun append-item (monkey item) 
  (with-slots (items) monkey 
    (setf items (append items (list item)))))


(defvar times 20)

(loop for x from 1 to times
      do (format t "~C~a/~a" #\return x times)
      do (finish-output)
      do (loop with rounds = monkeys
	       for monkey in rounds
	       do (with-slots (items operation test inspection-count) monkey
		    (loop for item = (pop items)
			  while item
			  for worry = (funcall operation item)
			  for relax = (floor (/ worry 3)) 
			  for next-monkey = (nth (funcall test relax) monkeys)
			  ; do (print worry)
			  do (append-item next-monkey relax)
			  do (incf inspection-count)))))


(let ((sorted (sort (mapcar #'inspection-count monkeys) #'>)))
  (print (* (first sorted) (second sorted))))
