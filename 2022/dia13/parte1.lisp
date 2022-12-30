

; em lisp por que lisp tambem tem o (eval)

(ql:quickload "cl-ppcre") ; meh

(defun sexp-translate (a)
  (let ((open-bracket (ppcre:create-scanner "\\[" ))
	(close-bracket (ppcre:create-scanner "\\]" ))
	(comma (ppcre:create-scanner "," )))
    (setf a (ppcre:regex-replace-all open-bracket a "(")
	  a (ppcre:regex-replace-all close-bracket a ")")
	  a (ppcre:regex-replace-all comma a " "))))

(defun line-empty-p (a) (equalp a ""))

(defvar inputs 
  (let ((ret nil))
    (setf ret (uiop:read-file-lines "packets.txt")
	  ret (remove-if #'line-empty-p ret)
	  ret (mapcar #'sexp-translate ret)
	  ret (mapcar #'read-from-string  ret))))

(print inputs)

