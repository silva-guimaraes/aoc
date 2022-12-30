

(defun make-chamber ()
  (make-hash-table :test #'equalp))

(defparameter *chamber* (make-chamber))

(defparameter *gas* (uiop:read-file-string "teste.txt"))

(defparameter *gas-len* (length *gas*))

(defparameter *floor* 0)

(defparameter *gas-current* 0)

;; bloco mais baixo de todas as formas é o ultimo
(defparameter *shapes*
  (list '((0 . 0) (1 . 0) (2 . 0) (3 . 0))
        '((0 . 1) (1 . 1) (2 . 1) (1 . 2) (1 . 0))
        '((0 . 0) (1 . 0) (2 . 2) (2 . 1) (2 . 0))
        '((0 . 3) (0 . 2) (0 . 1) (0 . 0))
        '((0 . 1) (1 . 1) (0 . 0) (1 . 0))))

(defun x (shape)
  (car shape))

(defun y (shape)
  (cdr shape))

(defun current-rock (x)
  (nth (mod x 5) *shapes*))

(defun current-flow () ;; feio
  (prog1 (aref *gas* (mod *gas-current* *gas-len*))
    ;; (format t "~a~%" (aref *gas* (mod *gas-current* *gas-len*)))
    (incf *gas-current*)))

(defun highest (shape)
  (1+ (loop for x in shape
        maximize (y x))))

(defun transform-shape (shape x y)
  (loop with transform = shape
        for j in transform
	collect (cons (+ (x j) x)
		      (+ (y j) y))))

(defun move-down (shape)
  (transform-shape shape 0 -1))

(defun collided-p (shape)
  (loop for i in shape
	thereis (< (cdr i) *floor*)
        thereis (gethash i *chamber*)))

(defun place-shape (shape)
  (loop for x in shape
        do (setf (gethash x *chamber*) t)))

(defun chamber-fit-p (shape)
  (loop for i in shape
	never (< (x i) 0)
	never (> (x i) 6)))

(defun gas-direction (gas shape)
  (let ((moved (if (equalp gas #\<)
		      (transform-shape shape -1 0)
		      (transform-shape shape  1 0))))
    (if (and (chamber-fit-p moved)
	     (not (collided-p moved)))
	moved
	shape)))

(defun row-clogged-p (y)
  (loop for x from 0 to 6
	always (gethash (cons x y) *chamber*)))

(defun shape-clogged-p (shape)
  (loop for y in (sort (remove-duplicates (mapcar #'y shape)) #'>)
	if (row-clogged-p y)
	  do (return y)))

(defun trim-chamber (shape)
  (let ((clogged (shape-clogged-p shape)))
    (when clogged
      (defun remblock (key v)
	(declare (ignore v))
	(when (<= (y key) clogged) (remhash key *chamber*)))
      (maphash #'remblock *chamber*)
      (setf *floor* (1+ clogged)))))

(defun simulate-rock (highest rock)
  (loop with current = (transform-shape rock 2 (+ 3 highest))
	with down = nil
	do (setf current (gas-direction (current-flow) current)
		 down (move-down current))
        if (collided-p down)
          do (place-shape current)
	  and do (trim-chamber current)
	  and do (return (highest current))
	else do (setf current down)))

(defun main ()
  (setf *gas-current* 0)
  (setf *floor* 0)
  (setf *chamber* (make-chamber))
  (loop for x from 0 to (1- 1000000000000)
	maximize (simulate-rock highest (current-rock x)) into highest
	do (format t "~C~a/1000000000000" #\return x) do (finish-output)
	finally (print highest)))

