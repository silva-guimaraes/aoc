
(print
 (let ((hands (loop for line in (uiop:read-file-lines "strategy_guide.txt")
		    collect (cons (+ 1 (- (char-int (aref line 0))
					  (char-int #\A) ))
				  (+ 1 (- (char-int (aref line 2))
					  (char-int #\X)))))))
   (loop for x in hands
	 for opponent = (car x)
	 for you = (cdr x)
	 for rotate = (+ opponent 1)
	 if (= opponent you)
	   sum (+ 3 you)
	 else if (= you (if (= rotate 4) 1 rotate))
		sum (+ 6 you)
	 else sum you)))
